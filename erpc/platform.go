package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/rpc"
	"runtime"
)

func init(){
	rpc.RegisterAdminFunc(PLATFORM, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		return []byte(runtime.GOOS),nil
	})
}
