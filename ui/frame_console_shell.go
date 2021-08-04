package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/network/sshc"
	"github.com/nomos/go-tools/cmds"
	"github.com/nomos/go-lokas/promise"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/keys"
	"go.uber.org/zap/zapcore"
	"io"
	"regexp"
	"strings"
	"time"
)

type ConsoleShell struct {
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

var _ cmds.IConsole = (*ConsoleShell)(nil)

func NewConsoleShell(owner vcl.IComponent,option... FrameOption) (root *ConsoleShell) {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *ConsoleShell) setup(){
	this.SetHeight(200)
	this.SetWidth(500)
	this.BottomPanel = vcl.NewPanel(this)
	this.BottomPanel.SetParent(this)
	this.BottomPanel.SetAlign(types.AlBottom)
	this.BottomPanel.SetHeight(50)
	this.CmdEdit = vcl.NewEdit(this.BottomPanel)
	this.CmdEdit.SetParent(this.BottomPanel)
	this.CmdEdit.SetAlign(types.AlClient)
	this.CmdEdit.BorderSpacing().SetBottom(12)
	this.CmdEdit.BorderSpacing().SetLeft(3)
	this.CmdEdit.BorderSpacing().SetRight(10)
	this.CmdEdit.BorderSpacing().SetTop(12)
	this.CmdEdit.SetHeight(24)
	this.CmdEdit.SetWidth(487)
	this.SendButton = vcl.NewButton(this.BottomPanel)
	this.SendButton.SetParent(this.BottomPanel)
	this.SendButton.SetAlign(types.AlRight)
	this.SendButton.BorderSpacing().SetAround(11)
	this.SendButton.SetHeight(20)
	this.SendButton.SetLeft(611)
	this.SendButton.SetWidth(75)
	this.SendButton.SetCaption("Send")
	this.ShellSelect = vcl.NewComboBox(this.BottomPanel)
	this.ShellSelect.SetParent(this.BottomPanel)
	this.ShellSelect.SetAlign(types.AlLeft)
	this.ShellSelect.BorderSpacing().SetAround(12)
	this.ShellSelect.SetHeight(20)
	this.ShellSelect.SetItemHeight(26)
	this.ShellSelect.SetLeft(11)
	this.ShellSelect.SetStyle(types.CsDropDownList)
	this.ShellSelect.SetWidth(100)
	this.Panel1 = vcl.NewPanel(this)
	this.Panel1.SetParent(this)
	this.Panel1.SetAlign(types.AlClient)
	this.Console = vcl.NewMemo(this.Panel1)
	this.Console.SetParent(this.Panel1)
	this.Console.SetAlign(types.AlClient)
	this.Console.SetColor(0x001D180F)
	this.Console.Font().SetColor(0x00c6b7a9)
	this.Console.Font().SetHeight(-13)
	this.Console.SetScrollBars(types.SsAutoBoth)
	this.Console.SetTop(1)
}

func (this *ConsoleShell) OnCreate(){
	this.setup()
	this.start()
	if log.IsConsole(){
		this.ComposeLogger = log.NewComposeLogger(true,log.ConsoleConfig(""),1)
	} else {
		this.ComposeLogger = log.NewComposeLogger(true,log.DefaultConfig(""),1)
	}
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
	loadCommands(this)
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
	this.ssh = sshc.NewSshClient("","","",log.IsConsole())
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

func (this *ConsoleShell) RegisterSender(s string,sender cmds.ICommandSender){
	this.AddShellType(s)
	this.senders[s] = sender
}

func (this *ConsoleShell) AddShellType(s string){
	count:=this.ShellSelect.Items().Count()
	for i :=int32(0);i<count;i++ {
		str:=this.ShellSelect.Items().S(i)
		if str== s {
			return
		}
	}
	this.ShellSelect.Items().Add(s)
	this.ShellSelect.SetText(s)
}

func (this *ConsoleShell) RegisterCommonCmd(command cmds.ICommand){
	this.commonCommands[command.Name()] = command
}

func (this *ConsoleShell) RegisterCommonCmdFunc(name string,tips string,f func(value *cmds.ParamsValue,console cmds.IConsole)*promise.Promise) {
	command:=cmds.NewCommand(name,tips,f,this)
	this.commonCommands[command.Name()] = command
}

func (this *ConsoleShell) RegisterCmd(typ string,command cmds.ICommand) {
	this.AddShellType(typ)
	if _,ok:=this.commands[typ];!ok {
		this.commands[typ] = make(map[string]cmds.ICommand)
	}
	this.commands[typ][command.Name()] = command
}

func (this *ConsoleShell) RegisterCmdFunc (typ string,name string,tips string,f func (value *cmds.ParamsValue,console cmds.IConsole)*promise.Promise) {
	this.AddShellType(typ)
	command:=cmds.NewCommand(name,tips,f,this)
	if _,ok:=this.commands[typ];!ok {
		this.commands[typ] = make(map[string]cmds.ICommand)
	}
	this.commands[typ][name] = command
}

func loadCommands(shell *ConsoleShell){
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


func (this *ConsoleShell) start(){
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

func (this *ConsoleShell) stop(){
	if this.done!=nil {
		this.done<- struct{}{}
	}
}

func (this *ConsoleShell) registerWrappedCmd(s string,cmd *cmds.WrappedCmd){
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

func (this *ConsoleShell) resetCache() {
	this.cacheReset = true
	this.cachedIndex = len(this.cachedText) -1
}


func (this *ConsoleShell) addCachedText(text string) {

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

func (this *ConsoleShell) sendCmd(text string){
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

func (this *ConsoleShell) OnDestroy(){
	this.stop()
}

func (this *ConsoleShell) Connect(user,pass,addr string)*promise.Promise{
	return this.ssh.Disconnect().Then(func(data interface{}) interface{} {
		this.ssh.SetAddr(user,pass,addr)
		return this.ssh.Connect()
	})
}

func (this *ConsoleShell) WriteString(str string){
	str = strings.TrimRight(str," ")
	str = strings.TrimRight(str,"\n")
	str = strings.TrimRight(str," ")
	if str!="" {
		this.msgChan<-str
	}
}

func (this *ConsoleShell) Write(p []byte)(int,error){
	str:=string(p)
	str = strings.TrimRight(str," ")
	str = strings.TrimRight(str,"\n")
	str = strings.TrimRight(str," ")
	if str!="" {
		this.msgChan<-str
	}
	return 0,nil
}

func (this *ConsoleShell)WriteConsole(e zapcore.Entry, p []byte) error {
	return nil
}
func (this *ConsoleShell)WriteJson(e zapcore.Entry, p []byte) error {
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
func (this *ConsoleShell)WriteObject(e zapcore.Entry, o map[string]interface{}) error {
	return nil
}

func (this *ConsoleShell) Disconnect()*promise.Promise {
	return this.ssh.Disconnect()
}

func (this *ConsoleShell) Clear(){
	vcl.ThreadSync(func() {
		this.Console.Clear()
	})
}

func (this *ConsoleShell) SendCmd(s string){
	this.WriteString(">"+s)
	go this.ssh.RunShellCmd(s,false).Await()
}

func (this *ConsoleShell) OnSelect(){

}

func (this *ConsoleShell) OnDeselect(){

}

func (this *ConsoleShell) ExecShellCmd(s string,isExpect bool)*promise.Promise{
	this.WriteString(">"+s)
	return this.ssh.NewShellSession().Run(s,isExpect)
}

func (this *ConsoleShell) ExecWrappedCmd(cmd *cmds.WrappedCmd,args... string)*promise.Promise{
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

func (this *ConsoleShell) ExecSshCmd(s string)*promise.Promise {
	this.WriteString(this.ssh.GetConnStr()+">"+s)
	return this.ssh.RunCmd(s)
}

func (this *ConsoleShell) OnCmdEditChange(sender vcl.IObject) {

}

func (this *ConsoleShell) RegisterCommands(commands []cmds.ICommand){
	for _,c:=range commands {
		this.RegisterCmd("shell",c)
	}
}


func (this *ConsoleShell) OnSendButtonClick(sender vcl.IObject) {

}

func (this *ConsoleShell) GetOutputs(data interface{})[]string{
	return data.([]string)
}

func (this *ConsoleShell) GetLastOutput(data interface{})string{
	if data == nil {
		return ""
	}
	ret:=data.([]string)
	if len(ret) == 0 {
		return ""
	}
	return ret[len(ret)-1]
}

func (this *ConsoleShell) OnEnter(){

}

func (this *ConsoleShell) OnExit(){
	this.Console.Clear()
}