package erpc

import (
	"context"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/nomos/clipboard"
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/network"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"sync"
)


//一个和electron typescript通信来实现访问golang原生功能的框架

func (this *App) watchClipBoard(){
	ch := clipboard.Watch(context.TODO(), clipboard.FmtImage)
	ch1 := clipboard.Watch(context.TODO(), clipboard.FmtBMP)
	go func() {
		for {
			select {
			case data:=<-ch:
				log.Infof("recv",data)
				println(`"text data" is no longer available from clipboard.`)
			case data:=<-ch1:
				log.Infof("recv",data)
				println(`"text data" is no longer available from clipboard.`)
			}
		}
	}()
}

var defaultWinOpt = &astilectron.WindowOptions{
	Center: astikit.BoolPtr(true),
	Height: astikit.IntPtr(800),
	Width:  astikit.IntPtr(1580),
	Title: astikit.StrPtr(""),
	WebPreferences: &astilectron.WebPreferences{
		WebSecurity: astikit.BoolPtr(false),
	},
}

func WithElectron(name string,defaultUrl string)Option{
	return func(a *App) {
		pwd,_:=util.ExecPath()
		a.electronApp, _ = astilectron.New(log.NewAstilecTronLogger(false), astilectron.Options{
			AppName: name,
			BaseDirectoryPath: pwd+"/astiletron/",
		})
		a.url = defaultUrl
	}
}

func WithConfig(name string)Option{
	return func(app *App) {
		app.config = lox.NewAppConfig(name)
	}
}

func WithElectronOption(opt *astilectron.WindowOptions)Option{
	return func(a *App) {
		a.gameWindowOpt = opt
	}
}

func WithPort(port string)Option{
	return func(app *App) {
		app.Port = port
	}
}

type Option func(app *App)

var instance *App
var once sync.Once

func Instance(opts ...Option)*App{
	var err error
	once.Do(func() {
		if instance ==nil {
			instance = NewApp(opts...)
			err=instance.Start()

		}
	})
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return instance
}

func Close()error{
	if instance!=nil {
		return instance.Stop()
	}
	return nil
}

type App struct {
	*lox.Gate
	Port string
	*Session

	electronApp *astilectron.Astilectron
	gameWindow *astilectron.Window
	gameWindowOpt *astilectron.WindowOptions
	config *lox.AppConfig

	devTool bool
	url string
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
	ret.start()
	return ret
}

func (this *App) start()error{
	if this.electronApp!=nil {
		this.electronApp.HandleSignals()
		this.electronApp.Start()
		err:=this.createGameWindow(this.url)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		if this.devTool {
			this.gameWindow.OpenDevTools()
		}
	}

	return nil
}

func (this *App) createGameWindow(url string)error{
	var err error
	winOpt:=defaultWinOpt
	if this.gameWindowOpt!=nil {
		winOpt = this.gameWindowOpt
	}
	this.gameWindow, err = this.electronApp.NewWindow(url, winOpt)
	if err!= nil  {
		return log.Error(err.Error())
	}
	err = this.gameWindow.Create()
	if err!= nil {
		return log.Error(err.Error())
	}
	return nil
}

func (this *App) closeGameWindow()error{
	err:=this.gameWindow.Close()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.gameWindow = nil
	return nil
}

func (this *App) SetUrl(url string)error{
	if this.url != url {
		this.url = url
		if this.electronApp!=nil {
			err:=this.closeGameWindow()
			if err != nil {
				log.Error(err.Error())
				return err
			}
			err=this.createGameWindow(url)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
	}
	return nil
}

func (this *App) SetDevTool(dev bool){
	this.devTool = dev
	if this.gameWindow==nil {
		return
	}
	if this.devTool {
		this.gameWindow.OpenDevTools()
	} else {
		this.gameWindow.CloseDevTools()
	}
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

func (this *App) Wait() {
	util.WaitForTerminate()
}

func (this *App) Stop() error{
	return this.Gate.Stop()
}
