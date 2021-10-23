package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
)

func init(){
	RegisterAdminFunc(OPEN_DEV_TOOL, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		Instance().SetDevTool(true)
		return nil,nil
	})
	RegisterAdminFunc(CLOSE_DEV_TOOL, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		Instance().SetDevTool(false)
		return nil,nil
	})
}
