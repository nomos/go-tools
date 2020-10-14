// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-log/log"
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
}

func (this *TConsoleShell) OnCreate(){
	this.cachedText = make([]string,0)
	this.ssh = ssh.NewSshClient("root","9ayl02bf","192.168.110.197:22")
	//go this.ssh.Connect().Await()
	this.ssh.SetConsoleWriter(this)
	this.CmdEdit.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		switch *key {
		case keys.VkReturn:
			text := this.CmdEdit.Text()
			this.SendCmd(text)
		case keys.VkUp:
			log.Warnf(this.cachedIndex,this.cachedText)
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
	this.Console.Clear()
	this.Panel1.SetOnResize(func(sender vcl.IObject) {
	})
}

func (this *TConsoleShell) OnDestroy(){

}

func (this *TConsoleShell) Connect(user,pass,addr string)*promise.Promise{
	return this.ssh.Disconnect().Then(func(data interface{}) interface{} {
		this.ssh.SetAddr(user,pass,addr)
		return this.ssh.Connect()
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
	this.Console.Clear()
}

func (this *TConsoleShell) SendCmd(s string){
	this.CmdEdit.SetText("")
	if strings.TrimSpace(s) == "clear" {
		this.Console.Clear()
		return
	}
	if len(this.cachedText) == 0 ||this.cachedText[len(this.cachedText)-1] != s {
		this.cachedText = append(this.cachedText, s)
		this.cachedIndex = len(this.cachedText)-1
	}
	if this.ssh.IsConnect() {
		go this.ssh.RunCmd(s).Await()
	} else {
		go this.ssh.RunShellCmd(s).Await()
	}
}

func (this *TConsoleShell) ExecShellCmd(s string)*promise.Promise{
	return this.ssh.NewShellSession().Run(s)
}

func (this *TConsoleShell) ExecCmd(s string)*promise.Promise {
	return this.ssh.RunCmd(s)
}