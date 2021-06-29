package vcltool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/network/sshc"
	"github.com/nomos/go-tools/cmds"
	"github.com/nomos/promise"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/keys"
	"go.uber.org/zap/zapcore"
	"io"
	"regexp"
	"strings"
	"time"
)

type TConsoleFrame struct {
	*vcl.TFrame
	BottomPanel  *vcl.TPanel
	CmdEdit      *vcl.TEdit
	SendButton   *vcl.TButton
	ShellSelect  *vcl.TComboBox
	Panel1       *vcl.TPanel
	Console      *vcl.TMemo

	ConfigAble
	*log.ComposeLogger
	ssh            *sshc.SshClient
	addr           string
	user           string
	pwd            string
	cachedText     []string
	cachedIndex    int
	writeChan      chan string
	sender         cmds.ICommandSender
	senders        map[string]cmds.ICommandSender
	commonCommands map[string]cmds.ICommand
	commands map[string]map[string]cmds.ICommand
	buffer *bytes.Buffer
	stdIn io.WriteCloser
	shell bool
	cacheReset bool
	msgChan chan string
	done chan struct{}
}

func NewConsoleFrame(owner vcl.IComponent) (root *TConsoleFrame)  {
	vcl.CreateResFrame(owner, &root)
	return
}

func (this *TConsoleFrame) setup(){
	this.SetAlign(types.AlClient)
	this.BottomPanel = vcl.NewPanel(this)
	this.BottomPanel.SetAlign(types.AlBottom)
	this.BottomPanel.SetParent(this)
	this.BottomPanel.SetHeight(50)
	this.Panel1 = vcl.NewPanel(this)
	this.Panel1.SetParent(this)
	this.Panel1.SetAlign(types.AlClient)
	this.Console = vcl.NewMemo(this.Panel1)
	this.Console.SetParent(this.Panel1)
	this.Console.SetAlign(types.AlClient)
	this.Console.SetText("")
	this.Console.SetColor(0x001d180f)
	this.Console.SetScrollBars(types.SsAutoBoth)
}

func (this *TConsoleFrame) OnCreate(){
	this.start()
	this.ComposeLogger = log.NewComposeLogger(true,log.DefaultConfig(""),1)
	this.ComposeLogger.SetConsoleWriter(this)
	this.senders = make(map[string]cmds.ICommandSender)
	this.commonCommands = make(map[string]cmds.ICommand)
	this.commands = make(map[string]map[string]cmds.ICommand)
	this.cachedText = this.conf.GetStringSlice("cached_text")
	if this.cachedText == nil {
		this.cachedText = make([]string,0)
	}
	this.cachedIndex = len(this.cachedText)-1
	this.cachedIndex = this.conf.GetInt("cached_index")
	loadCmds(this)
	this.RegisterSender("shell",this)
	this.RegisterCommonCmdFunc("clear","clear console", func(value *cmds.ParamsValue,console cmds.IConsole) *promise.Promise {
		return promise.Async(func(resolve func(interface{}), reject func(interface{})) {
			console.Clear()
			resolve(nil)
		})
	})
	for k,v:=range cmds.GetAllCmds() {
		this.registerWrappedCmd(k,v)
	}
	this.ssh = sshc.NewSshClient("","","")
	this.ssh.SetStringWriter(this)
	this.CmdEdit.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		switch *key {
		case keys.VkReturn:
			text := this.CmdEdit.Text()
			this.sendCmd(text)
			this.CmdEdit.SetText("")
			this.addCachedText(text)
		case keys.VkUp:
			if this.cachedIndex >= 0 && len(this.cachedText) > this.cachedIndex {
				text := this.cachedText[this.cachedIndex]
				this.CmdEdit.SetText(text)
				this.cachedIndex--
			} else {
				this.cachedIndex = 0
			}
		case keys.VkDown:
			if this.cachedIndex >= 0 && len(this.cachedText) > this.cachedIndex {
				text := this.cachedText[this.cachedIndex]
				this.CmdEdit.SetText(text)
				this.cachedIndex++
			} else {
				this.cachedIndex = len(this.cachedText) - 1
			}
		default:
			this.resetCache()
		}
	})
	this.SendButton.SetOnClick(func(sender vcl.IObject) {
		text := this.CmdEdit.Text()
		this.sendCmd(text)
		this.CmdEdit.SetText("")
		this.addCachedText(text)
	})
	this.Console.Clear()
}

func (this *TConsoleFrame) RegisterSender(s string,sender cmds.ICommandSender){
	this.AddShellType(s)
	this.senders[s] = sender
}

func (this *TConsoleFrame) AddShellType(s string){
	count:=this.ShellSelect.Items().Count()
	for i :=int32(0);i<count;i++ {
		str:=this.ShellSelect.Items().S(i)
		if str== s {
			return
		}
	}
	this.ShellSelect.Items().Add(s)
}

func (this *TConsoleFrame) RegisterCommonCmd(command cmds.ICommand){
	this.commonCommands[command.Name()] = command
}

func (this *TConsoleFrame) RegisterCommonCmdFunc(name string,tips string,f func(value *cmds.ParamsValue,console cmds.IConsole)*promise.Promise) {
	command:=cmds.NewCommand(name,tips,f,this)
	this.commonCommands[command.Name()] = command
}

func (this *TConsoleFrame) RegisterCmd(typ string,command cmds.ICommand) {
	this.AddShellType(typ)
	if _,ok:=this.commands[typ];!ok {
		this.commands[typ] = make(map[string]cmds.ICommand)
	}
	this.commands[typ][command.Name()] = command
}

func (this *TConsoleFrame) RegisterCmdFunc (typ string,name string,tips string,f func (value *cmds.ParamsValue,console cmds.IConsole)*promise.Promise) {
	this.AddShellType(typ)
	command:=cmds.NewCommand(name,tips,f,this)
	if _,ok:=this.commands[typ];!ok {
		this.commands[typ] = make(map[string]cmds.ICommand)
	}
	this.commands[typ][name] = command
}

func loadCmds(shell *TConsoleFrame){
	shell.RegisterCmdFunc("shell",".makeSsh [name]",".makeSsh", func(value *cmds.ParamsValue, console cmds.IConsole) *promise.Promise {
		name:=value.String()
		cmdstr1:=`
		#!/usr/bin/expect;
		ssh-keygen -t rsa -C `+name+`;
		
		`
		return shell.ssh.RunShellCmd(cmdstr1,false).Then(func(data interface{}) interface{} {
			return shell.ssh.RunShellCmd("ssh-add ~/.ssh/id_rsa",false)
		})
	})
}

func (this *TConsoleFrame) start(){
	this.msgChan = make(chan string,1000)
	this.done = make(chan struct{})
	go func() {
		for {
			select{
			case str:=<-this.msgChan:
				vcl.ThreadSync(func() {
					this.Console.Lines().Add(str)
				})
			case <-this.done:
				return
			}
			time.Sleep(1)
		}
	}()
}

func (this *TConsoleFrame) stop(){
	if this.done!=nil {
		this.done<- struct{}{}
	}
}

func (this *TConsoleFrame) registerWrappedCmd(s string,cmd *cmds.WrappedCmd){
	this.RegisterCommonCmdFunc(s,cmd.Tips, func(value *cmds.ParamsValue, console cmds.IConsole) *promise.Promise {
		params:=[]string{}
		defer func() {
			if r := recover(); r != nil {
				err:=fmt.Sprintf("%v",r)
				this.Error(err)
			}
		}()
		for i:=0;i<cmd.ParamsNum;i++{
			params = append(params, value.String())
		}
		return this.ExecWrappedCmd(cmd,params...)
	})
}

func (this *TConsoleFrame) resetCache() {
	this.cacheReset = true
	this.cachedIndex = len(this.cachedText) -1
}


func (this *TConsoleFrame) addCachedText(text string) {

	if strings.TrimSpace(text)=="" {
		return
	}
	texts:=make([]string,0)
	for _,cmds:=range this.cachedText {
		if cmds != text {
			texts = append(texts, cmds)
		}
	}
	texts = append(texts, text)
	this.cachedText = texts
	this.resetCache()
	this.conf.Set("cached_text",this.cachedText)
	this.conf.Set("cached_index",this.cachedIndex)
}

func (this *TConsoleFrame) sendCmd(text string){
	if this.shell {
		int,err:=this.stdIn.Write([]byte(text+"\n"))
		if err != nil {
			log.Error(err.Error())
		}
		log.Warnf(int)

		return
	}
	name,params:=cmds.SplitCommandParams(text)
	s:=this.ShellSelect.Text()
	if this.commands[s]!=nil {
		if cmd,ok:=this.commands[s][name];ok {
			para:=cmds.NewParamsValue(name,params...)
			go cmd.ConsoleExec(para,this).Await()
			return
		}
	}
	if cmd,ok:=this.commonCommands[name];ok {
		para:=cmds.NewParamsValue(name,params...)
		go cmd.ConsoleExec(para,this).Await()
		return
	}
	sender:=this.senders[s]
	if sender != nil {
		sender.SendCmd(text)
	}
}

func (this *TConsoleFrame) OnDestroy(){
	this.stop()
}

func (this *TConsoleFrame) Connect(user,pass,addr string)*promise.Promise{
	return this.ssh.Disconnect().Then(func(data interface{}) interface{} {
		this.ssh.SetAddr(user,pass,addr)
		return this.ssh.Connect()
	})
}

func (this *TConsoleFrame) WriteString(str string){
	str = strings.TrimRight(str," ")
	str = strings.TrimRight(str,"\n")
	str = strings.TrimRight(str," ")
	if str!="" {
		this.msgChan<-str
	}
}

func (this *TConsoleFrame) Write(p []byte)(int,error){
	str:=string(p)
	str = strings.TrimRight(str," ")
	str = strings.TrimRight(str,"\n")
	str = strings.TrimRight(str," ")
	if str!="" {
		this.msgChan<-str
	}
	return 0,nil
}

func (this *TConsoleFrame)WriteConsole(e zapcore.Entry, p []byte) error {
	return nil
}
func (this *TConsoleFrame)WriteJson(e zapcore.Entry, p []byte) error {
	var j map[string]interface{}
	json.Unmarshal(p,&j)
	level:=j["level"].(string)
	time:=j["time"].(string)
	caller:=j["caller"].(string)
	msg:=j["msg"].(string)
	o:=make(map[string]interface{})
	for k,v:=range j {
		if k!="level"&&k!="time"&&k!="caller"&&k!="msg" {
			o[k] = v
		}
	}
	level = regexp.MustCompile(`[[][A-Z]+[]][A-z]*`).FindString(level)
	jstr,_:=json.Marshal(o)
	str := time+" "+level+"   "+caller+" "+msg+" "+string(jstr)
	this.Write([]byte(str))
	return nil
}
func (this *TConsoleFrame)WriteObject(e zapcore.Entry, o map[string]interface{}) error {
	return nil
}

func (this *TConsoleFrame) Disconnect()*promise.Promise {
	return this.ssh.Disconnect()
}

func (this *TConsoleFrame) Clear(){
	vcl.ThreadSync(func() {
		this.Console.Clear()
	})
}

func (this *TConsoleFrame) SendCmd(s string){
	this.WriteString(">"+s)
	go this.ssh.RunShellCmd(s,false).Await()
}

func (this *TConsoleFrame) OnSelect(){

}

func (this *TConsoleFrame) OnDeselect(){

}

func (this *TConsoleFrame) ExecShellCmd(s string,isExpect bool)*promise.Promise{
	this.WriteString(">"+s)
	return this.ssh.NewShellSession().Run(s,isExpect)
}

func (this *TConsoleFrame) ExecWrappedCmd(cmd *cmds.WrappedCmd,args... string)*promise.Promise{
	s:=cmd.FillParams(args...)
	this.WriteString(">"+s)
	return promise.Async(func(resolve func(interface{}), reject func(interface{})) {
		go func() {
			outputs,err:=this.ssh.NewShellSession().Run(s,cmd.CmdType==cmds.Cmd_Expect).Await()
			if err != nil {
				log.Error(err.Error())
				reject(err)
				return
			}
			if cmd.CmdHandler!=nil {
				res:=cmd.CmdHandler(outputs.([]string))
				if res.Success {
					this.Infof(s+" success ",res.Results)
				} else {
					this.Errorf(s+" failed ",res.Results)
				}
				resolve(res)
				return
			}
			resolve(&cmds.CmdResult{
				Outputs: outputs.([]string),
				Success: false,
				Results: make(map[string]interface{}),
			})
			this.Warnf(fmt.Sprintf(s+" failed"))
		}()

	})
}

func (this *TConsoleFrame) ExecSshCmd(s string)*promise.Promise {
	this.WriteString(this.ssh.GetConnStr()+">"+s)
	return this.ssh.RunCmd(s)
}

func (this *TConsoleFrame) OnCmdEditChange(sender vcl.IObject) {

}


func (this *TConsoleFrame) OnSendButtonClick(sender vcl.IObject) {

}

func (this *TConsoleFrame) GetOutputs(data interface{})[]string{
	return data.([]string)
}

func (this *TConsoleFrame) GetLastOutput(data interface{})string{
	if data == nil {
		return ""
	}
	ret:=data.([]string)
	if len(ret) == 0 {
		return ""
	}
	return ret[len(ret)-1]
}