package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/lox"
)

func init(){
	registerFunc(OPEN_DEV_TOOL, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		Instance().SetDevTool(true)
		return nil,nil
	})
	registerFunc(CLOSE_DEV_TOOL, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		Instance().SetDevTool(false)
		return nil,nil
	})
}