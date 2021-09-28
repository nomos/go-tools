package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/network"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"sync"
)


//一个和electron typescript通信来实现访问golang原生功能的框架
func WithPort(port string)Option{
	return func(app *App) {
		app.Port = port
	}
}

type Option func(app *App)

var instance *App
var once sync.Once

func Start()*App{
	once.Do(func() {
		if instance ==nil {
			instance = NewApp(WithPort("13333"))
			instance.Start()
		}
	})
	return instance
}

func Close(){
	if instance!=nil {
		instance.Stop()
	}
}

func Call(command *lox.AdminCommand)([]byte,error){
	return Start().CallAdminCommand(command)
}

type App struct {
	*lox.Gate
	Port string
	*Session
	sid util.ID
	mutex sync.Mutex
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
	ret.SessionCreatorFunc = ret.SessionCreator
	ret.Gate.LoadCustom("0.0.0.0",ret.Port,protocol.BINARY,lox.Websocket)
	return ret
}

func (this *App) genId()util.ID{
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.sid++
	return this.sid
}

func (this *App) SessionCreator(conn lokas.IConn) lokas.ISession {
	if this.Session!=nil {
		this.Session.Conn.Close()
		this.Session = nil
	}
	this.Session = NewSession(conn, this.genId(), this)
	go this.Session.Start()
	return this.Session
}

func (this *App) Start() error{
	return this.Gate.Start()
}

func (this *App) Stop() error{
	return this.Gate.Stop()
}
