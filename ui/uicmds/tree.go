package uicmds

import (
	"github.com/nomos/go-tools/cmds"
	"github.com/nomos/promise"
)

func init() {
	RegisterCommand(TreeAdd)
}


var TreeAdd = cmds.NewCommandNoConsole("tadd","tadd", func(value *cmds.ParamsValue, console cmds.IConsole) *promise.Promise {
	console.Warnf("Tree Add")
	return promise.Resolve(nil)
})
