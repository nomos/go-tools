package uicmds

import "github.com/nomos/go-tools/cmds"

var Commands = make([]cmds.ICommand,0)

func RegisterCommand(c cmds.ICommand){
	Commands = append(Commands, c)
}
