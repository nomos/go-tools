// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-lokas/network/ssh"
	"github.com/nomos/go-promise"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/keys"
	"strings"
)

//::private::
type TConsoleShellFields struct {
	ConfigAble
	ssh *ssh.SshClient
	addr string
	user string
	pwd string
	cachedText      []string
	cachedIndex     int
	writeChan chan string
	sender ICommandSender
	senders map[string]ICommandSender
	commands map[string]ICommand
}

func (this *TConsoleShell) RegisterSender(s string,sender ICommandSender){
	this.ShellSelect.Items().Add(s)
	this.senders[s] = sender
}

func (this *TConsoleShell) RegisterCmd(command ICommand){
	this.commands[command.Name()] = command
}

func (this *TConsoleShell) RegisterCmdFunc(name string,tips string,f func(value *ParamsValue,console *TConsoleShell)*promise.Promise) {
	command:=NewCommand(name,tips,f)
	this.commands[command.Name()] = command
}

func (this *TConsoleShell) OnCreate(){
	this.senders = make(map[string]ICommandSender)
	this.commands = make(map[string]ICommand)
	this.RegisterSender("shell",this)
	this.RegisterCmdFunc("clear","clear console", func(value *ParamsValue,console *TConsoleShell) *promise.Promise {
		return promise.Async(func(resolve func(interface{}), reject func(interface{})) {
			console.Clear()
			resolve(nil)
		})
	})
	this.cachedText = make([]string,0)
	this.ssh = ssh.NewSshClient("","","")
	this.ssh.SetConsoleWriter(this)
	this.CmdEdit.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		switch *key {
		case keys.VkReturn:
			text := this.CmdEdit.Text()
			this.sendCmd(text)
			this.CmdEdit.SetText("")
			if len(this.cachedText) == 0 ||this.cachedText[len(this.cachedText)-1] != text {
				this.cachedText = append(this.cachedText, text)
				this.cachedIndex = len(this.cachedText)-1
			}
		case keys.VkUp:
			if this.cachedIndex >= 0 && len(this.cachedText) > this.cachedIndex {
				text := this.cachedText[len(this.cachedText)-this.cachedIndex-1]
				this.CmdEdit.SetText(text)
				this.cachedIndex++
			} else {
				this.cachedIndex = len(this.cachedText) - 1
			}
		case keys.VkDown:
			if this.cachedIndex >= 0 && len(this.cachedText) > this.cachedIndex {
				text := this.cachedText[len(this.cachedText)-this.cachedIndex-1]
				this.CmdEdit.SetText(text)
				this.cachedIndex--
			} else {
				this.cachedIndex = 0
			}
		}
	})
	this.SendButton.SetOnClick(func(sender vcl.IObject) {
		text := this.CmdEdit.Text()
		this.sendCmd(text)
		this.CmdEdit.SetText("")
		if len(this.cachedText) == 0 ||this.cachedText[len(this.cachedText)-1] != text {
			this.cachedText = append(this.cachedText, text)
			this.cachedIndex = len(this.cachedText)-1
		}
	})
	this.Console.Clear()
}

func (this *TConsoleShell) sendCmd(text string){
	s:=this.ShellSelect.Text()
	sender:=this.senders[s]
	sender.SendCmd(text)
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
	name,params:=SplitCommand(s)
	if cmd,ok:=this.commands[name];ok {
		para:=&ParamsValue{
			cmd: name,
			value:  params,
			offset: 0,
		}
		go cmd.Exec(para,this).Await()
		return
	}
	go this.ssh.RunShellCmd(s).Await()
}

func (this *TConsoleShell) OnSelect(){

}

func (this *TConsoleShell) OnDeselect(){

}

func (this *TConsoleShell) ExecShellCmd(s string)*promise.Promise{
	return this.ssh.NewShellSession().Run(s)
}

func (this *TConsoleShell) ExecSshCmd(s string)*promise.Promise {
	return this.ssh.RunCmd(s)
}

func (this *TConsoleShell) OnCmdEditChange(sender vcl.IObject) {

}


func (this *TConsoleShell) OnSendButtonClick(sender vcl.IObject) {

}

