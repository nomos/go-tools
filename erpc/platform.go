package erpc

import (
	"encoding/json"
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/rpc"
	"github.com/super-l/machine-code/machine"
	"runtime"
)

func init(){
	rpc.RegisterAdminFunc(PLATFORM, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		return []byte(runtime.GOOS),nil
	})
	rpc.RegisterAdminFunc(MACHINE_DATA, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		data:=machine.GetMachineData()
		return json.Marshal(data)
	})
}
