package vcltool

import "github.com/nomos/go-promise"

func loadCommands(shell *TConsoleShell){
	shell.RegisterCmdFunc("shell",".makeSsh [name]",".makeSsh", func(value *ParamsValue, console *TConsoleShell) *promise.Promise {
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