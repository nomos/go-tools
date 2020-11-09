// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"bytes"
	"fmt"
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/network/sshc"
	"github.com/nomos/go-promise"
	"github.com/nomos/go-tools/cmds"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/keys"
	"io"
	"strings"
)

//::private::
type TConsoleShellFields struct {
	ConfigAble
	*log.ComposeLogger
	ssh            *sshc.SshClient
	addr           string
	user           string
	pwd            string
	cachedText     []string
	cachedIndex    int
	writeChan      chan string
	sender         ICommandSender
	senders        map[string]ICommandSender
	commonCommands map[string]ICommand
	commands map[string]map[string]ICommand
	buffer *bytes.Buffer
	stdIn io.WriteCloser
	shell bool
	cacheReset bool
}

func (this *TConsoleShell) RegisterSender(s string,sender ICommandSender){
	this.AddShellType(s)
	this.senders[s] = sender
}

func (this *TConsoleShell) AddShellType(s string){
	count:=this.ShellSelect.Items().Count()
	for i :=int32(0);i<count;i++ {
		str:=this.ShellSelect.Items().S(i)
		if str== s {
			return
		}
	}
	this.ShellSelect.Items().Add(s)
}

func (this *TConsoleShell) RegisterCommonCmd(command ICommand){
	this.commonCommands[command.Name()] = command
}

func (this *TConsoleShell) RegisterCommonCmdFunc(name string,tips string,f func(value *ParamsValue,console *TConsoleShell)*promise.Promise) {
	command:=NewCommand(name,tips,f)
	this.commonCommands[command.Name()] = command
}

func (this *TConsoleShell) RegisterCmd(typ string,command ICommand) {
	this.AddShellType(typ)
	if _,ok:=this.commands[typ];!ok {
		this.commands[typ] = make(map[string]ICommand)
	}
	this.commands[typ][command.Name()] = command
}

func (this *TConsoleShell) RegisterCmdFunc (typ string,name string,tips string,f func (value *ParamsValue,console *TConsoleShell)*promise.Promise) {
	this.AddShellType(typ)
	command:=NewCommand(name,tips,f)
	if _,ok:=this.commands[typ];!ok {
		this.commands[typ] = make(map[string]ICommand)
	}
	this.commands[typ][name] = command
}

func (this *TConsoleShell) OnCreate(){
	this.ComposeLogger = log.NewComposeLogger(true,log.DefaultConfig(),1)
	this.ComposeLogger.SetConsoleWriter(this)
	this.senders = make(map[string]ICommandSender)
	this.commonCommands = make(map[string]ICommand)
	this.commands = make(map[string]map[string]ICommand)
	this.cachedText = this.conf.GetStringSlice("cached_text")
	if this.cachedText == nil {
		this.cachedText = make([]string,0)
	}
	this.cachedIndex = len(this.cachedText)-1
	this.cachedIndex = this.conf.GetInt("cached_index")
	loadCommands(this)
	this.RegisterSender("shell",this)
	this.RegisterCommonCmdFunc("clear","clear console", func(value *ParamsValue,console *TConsoleShell) *promise.Promise {
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

func (this *TConsoleShell) registerWrappedCmd(s string,cmd *cmds.WrappedCmd){
	this.RegisterCommonCmdFunc(s,cmd.Tips, func(value *ParamsValue, console *TConsoleShell) *promise.Promise {
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

func (this *TConsoleShell) resetCache() {
	this.cacheReset = true
	this.cachedIndex = len(this.cachedText) -1
}


func (this *TConsoleShell) addCachedText(text string) {

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

func (this *TConsoleShell) sendCmd(text string){
	if this.shell {
		int,err:=this.stdIn.Write([]byte(text+"\n"))
		if err != nil {
			log.Error(err.Error())
		}
		log.Warnf(int)

		return
	}
	name,params:=SplitCommand(text)
	s:=this.ShellSelect.Text()
	if this.commands[s]!=nil {
		if cmd,ok:=this.commands[s][name];ok {
			para:=&ParamsValue{
				cmd: name,
				value:  params,
				offset: 0,
			}
			go cmd.Exec(para,this).Await()
			return
		}
	}
	if cmd,ok:=this.commonCommands[name];ok {
		para:=&ParamsValue{
			cmd: name,
			value:  params,
			offset: 0,
		}
		go cmd.Exec(para,this).Await()
		return
	}
	sender:=this.senders[s]
	if sender != nil {
		sender.SendCmd(text)
	}
}

func (this *TConsoleShell) OnDestroy(){

}

func (this *TConsoleShell) Connect(user,pass,addr string)*promise.Promise{
	return this.ssh.Disconnect().Then(func(data interface{}) interface{} {
		this.ssh.SetAddr(user,pass,addr)
		return this.ssh.Connect()
	})
}

func (this *TConsoleShell) WriteString(str string){
	vcl.ThreadSync(func() {
		str = strings.TrimRight(str," ")
		str = strings.TrimRight(str,"\n")
		str = strings.TrimRight(str," ")
		if str!="" {
			this.Console.Lines().Add(str)
		}
	})
}

func (this *TConsoleShell) Write(p []byte)(int,error){
	vcl.ThreadSync(func() {
		str:=string(p)
		str = strings.TrimRight(str," ")
		str = strings.TrimRight(str,"\n")
		str = strings.TrimRight(str," ")
		if str!="" {
			this.Console.Lines().Add(str)
		}
	})
	return 0,nil
}

func (this *TConsoleShell) Disconnect()*promise.Promise {
	return this.ssh.Disconnect()
}

func (this *TConsoleShell) Clear(){
	vcl.ThreadSync(func() {
		this.Console.Clear()
	})
}

func (this *TConsoleShell) SendCmd(s string){
	this.WriteString(">"+s)
	go this.ssh.RunShellCmd(s,false).Await()
}

func (this *TConsoleShell) OnSelect(){

}

func (this *TConsoleShell) OnDeselect(){

}

func (this *TConsoleShell) ExecShellCmd(s string,isExpect bool)*promise.Promise{
	this.WriteString(">"+s)
	return this.ssh.NewShellSession().Run(s,isExpect)
}

func (this *TConsoleShell) ExecWrappedCmd(cmd *cmds.WrappedCmd,args... string)*promise.Promise{
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

func (this *TConsoleShell) ExecSshCmd(s string)*promise.Promise {
	this.WriteString(this.ssh.GetConnStr()+">"+s)
	return this.ssh.RunCmd(s)
}

func (this *TConsoleShell) OnCmdEditChange(sender vcl.IObject) {

}


func (this *TConsoleShell) OnSendButtonClick(sender vcl.IObject) {

}

func (this *TConsoleShell) GetOutputs(data interface{})[]string{
	return data.([]string)
}

func (this *TConsoleShell) GetLastOutput(data interface{})string{
	ret:=data.([]string)
	if len(ret) == 0 {
		return ""
	}
	return ret[len(ret)-1]
}
