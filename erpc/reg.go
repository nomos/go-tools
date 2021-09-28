package erpc

import "github.com/nomos/go-lokas/lox"

type rpcFunc func(cmd *lox.AdminCommand)([]byte,error)

var rpcHandlers = map[string]rpcFunc{}

func registerFunc(name string,f rpcFunc) {
	rpcHandlers[name] = f
}