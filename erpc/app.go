package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/network"
	"github.com/nomos/go-lokas/protocol"
)

//一个和electron typescript通信来实现访问golang原生功能的框架
func WithPort(port string)Option{
	return func(app *App) {
		app.Port = port
	}
}

type Option func(app *App)

type App struct {
	*lox.Gate
	Port string
}

func NewApp(opts ...Option) *App {
	ret := &App{
		Gate: &lox.Gate{
			Actor:           lox.NewActor(),
			ISessionManager: network.NewDefaultSessionManager(true),
		},
	}
	for _, o := range opts {
		o(ret)
	}
	ret.Gate.LoadCustom("0.0.0.0",ret.Port,protocol.BINARY,lox.Websocket)
	return ret
}

func (this *App) SessionCreator(conn lokas.IConn) lokas.ISession {
	sess := lox.NewPassiveSession(conn, this.GetProcess().GenId(), this)
	sess.AuthFunc = func(data []byte) error {
		return nil
	}
	sess.Protocol = this.Protocol
	this.ISessionManager.AddSession(sess.GetId(), sess)
	this.GetProcess().AddActor(sess)
	this.GetProcess().StartActor(sess)
	return sess
}

func (this *App) Start() error{
	return this.Gate.Start()
}

func (this *App) Stop() error{
	return this.Gate.Stop()
}
