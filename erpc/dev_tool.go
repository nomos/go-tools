package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/rpc"
)

func init(){
	rpc.RegisterAdminFunc(OPEN_DEV_TOOL, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		Instance().SetDevTool(true)
		return nil,nil
	})
	rpc.RegisterAdminFunc(CLOSE_DEV_TOOL, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		Instance().SetDevTool(false)
		return nil,nil
	})
}
