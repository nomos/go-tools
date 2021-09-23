package erpc

import (
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/network"
)

//一个和electron typescript通信来实现访问golang原生功能的框架

type Option func(app *App)

type App struct {
	*lox.Gate
	Port int
}

func NewApp(opts ...Option) *App {
	ret := &App{
		Gate: &lox.Gate{
			Actor:           lox.NewActor(),
			ISessionManager: network.NewDefaultSessionManager(true),
		},
	}
	ret.SetType("Gate")
	for _, o := range opts {
		o(ret)
	}
	return ret
}

func (this *App) Start() {

}
