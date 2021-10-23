package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
)

type rpcFunc func(cmd *lox.AdminCommand,params *cmds.ParamsValue,logger log.ILogger)([]byte,error)

var rpcHandlers = map[string]rpcFunc{}

func RegisterAdminFunc(name string,f rpcFunc) {
	rpcHandlers[name] = f
}