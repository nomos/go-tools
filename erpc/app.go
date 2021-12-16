package erpc

import (
	"bytes"
	"context"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/nomos/clipboard"
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/lox/errs"
	"github.com/nomos/go-lokas/network"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-lokas/util/events"
	"golang.org/x/image/tiff"
	"image"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

//一个和electron typescript通信来实现访问golang原生功能的框架

func (this *App) watchImgClipBoard() {
	ch := clipboard.Watch(context.TODO(), clipboard.FmtImage)
	ch1 := clipboard.Watch(context.TODO(), clipboard.FmtBMP)
	go func() {
		for {
			select {
			case data := <-ch1:
				log.Infof("clip bmp data")
				this.clipImgData = data
			case data := <-ch:
				log.Infof("clip png data")
				this.clipImgData = data
			}
		}
	}()
}

func (this *App) GetClipBoardImg()(image.Image,error){
	var img image.Image
	var err error
	if runtime.GOOS=="darwin" {
		img,err=tiff.Decode(bytes.NewBuffer(this.clipImgData))
		if err != nil {
			log.Error(err.Error())
			return nil,err
		}
		return img,nil
	}
	img, _, err = image.Decode(bytes.NewBuffer(this.clipImgData))
	if err != nil {
		log.Error(err.Error())
		return nil,err
	}
	return img,nil
}

var defaultWinOpt = &astilectron.WindowOptions{
	Center: astikit.BoolPtr(true),
	Height: astikit.IntPtr(800),
	Width:  astikit.IntPtr(1580),
	MinWidth: astikit.IntPtr(960),
	MinHeight: astikit.IntPtr(480),
	Title:  astikit.StrPtr(""),
	WebPreferences: &astilectron.WebPreferences{
		WebSecurity: astikit.BoolPtr(false),
	},
}

func WithHandler(handler func(msg *protocol.BinaryMessage)(protocol.ISerializable,error))Option {
	return func(app *App) {
		app.handler = func(msg *protocol.BinaryMessage, session *Session) {
			ret,err:=handler(msg)
			if err!=nil {
				if e,ok:=err.(*protocol.ErrMsg);ok {
					session.SendMessage(0,msg.TransId,e)
				} else {
					session.SendMessage(0,msg.TransId,protocol.NewError(errs.ERR_INTERNAL_SERVER))
				}
				return
			}
			session.SendMessage(0,msg.TransId,ret)
		}
	}
}

func WithElectron(name string, defaultUrl string) Option {
	return func(a *App) {
		pwd:=filepath.Dir(os.Args[0])
		a.electronApp, _ = astilectron.New(log.NewAstilecTronLogger(false), astilectron.Options{
			AppName:           name,
			BaseDirectoryPath: pwd + "/astiletron/",
		})
		a.url = defaultUrl
	}
}

func WithConfig(name string) Option {
	return func(app *App) {
		app.config = lox.NewAppConfig(name)
	}
}

func WithAppConfig(conf lokas.IConfig) Option {
	return func(app *App) {
		app.config = conf
	}
}

func WithElectronOption(opt *astilectron.WindowOptions) Option {
	return func(a *App) {
		a.gameWindowOpt = opt
	}
}

func WithDevTool() Option {
	return func(app *App) {
		app.devTool = true
	}
}

func WithPort(port string) Option {
	return func(app *App) {
		app.Port = port
	}
}

type Option func(app *App)

var instance *App
var once sync.Once

func Instance(opts ...Option) *App {
	var err error
	once.Do(func() {
		if instance == nil {
			instance = NewApp(opts...)
		}
	})
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return instance
}

func Close() error {
	if instance != nil {
		return instance.Stop()
	}
	return nil
}

type App struct {
	events.EventEmmiter
	*lox.Gate
	Port string
	*Session
	handler func(msg *protocol.BinaryMessage,session *Session)
	electronApp   *astilectron.Astilectron
	gameWindow    *astilectron.Window
	gameWindowOpt *astilectron.WindowOptions
	config        lokas.IConfig

	devTool bool
	url     string
	sid     util.ID
	mutex   sync.Mutex
	done    chan struct{}
	clipImgData []byte
}

func NewApp(opts ...Option) *App {
	ret := &App{
		EventEmmiter:events.New(),
		Gate: &lox.Gate{
			Actor:           lox.NewActor(),
			ISessionManager: network.NewDefaultSessionManager(true),
		},
	}
	for _, o := range opts {
		o(ret)
	}
	ret.SessionCreatorFunc = ret.SessionCreator
	return ret
}

func (this *App) start() error {
	this.watchImgClipBoard()
	if this.electronApp != nil {
		this.electronApp.HandleSignals()
		this.electronApp.Start()
		err := this.createGameWindow(this.url)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		if this.devTool {
			this.gameWindow.OpenDevTools()
		}
		this.electronApp.Wait()
		this.Gate.Stop()
		this.electronApp.Close()
	}
	return nil
}

func (this *App) stop() error {
	if this.gameWindow != nil {
		err := this.gameWindow.Close()
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}
	this.electronApp.Close()
	return nil
}

func reformWidth(s int)int{
	if runtime.GOOS == "windows" {
		return (s-1)/2*2+1
	}
	return (s-1)/2*2+1
}

func reformHeight(s int)int{
	if runtime.GOOS == "windows" {
		return (s-1)/2*2+1
	}
	return (s)/2*2
}

func (this *App) createGameWindow(url string) error {
	var err error
	winOpt := defaultWinOpt
	if this.gameWindowOpt != nil {
		winOpt = this.gameWindowOpt
	}
	width:=1280
	height:=720
	if this.config.GetInt("width")!=0 {
		width=this.config.GetInt("width")
	}
	if this.config.GetInt("height")!=0 {
		height=this.config.GetInt("height")
	}
	*(winOpt.Width) = reformWidth(width)
	*(winOpt.Height) = reformHeight(height)
	this.gameWindow, err = this.electronApp.NewWindow(url, winOpt)
	if err != nil {
		return log.Error(err.Error())
	}
	err = this.gameWindow.Create()
	if err != nil {
		return log.Error(err.Error())
	}
	this.gameWindow.On(astilectron.EventNameWindowEventResize, func(e astilectron.Event) (deleteListener bool) {
		log.Infof(this.config,e)
		this.config.Set("width",reformWidth(*(e.WindowOptions.Width)))
		this.config.Set("height",reformHeight(*(e.WindowOptions.Height)))
		return false
	})
	return nil
}

func (this *App) ResetSize(){
	width:=1280
	height:=720
	if this.config.GetInt("width")!=0 {
		width=this.config.GetInt("width")
	}
	if this.config.GetInt("height")!=0 {
		height=this.config.GetInt("height")
	}
	this.gameWindow.Resize(reformWidth(width),reformHeight(height))
}

func (this *App) closeGameWindow() error {
	err := this.gameWindow.Close()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.gameWindow = nil
	return nil
}

func (this *App) SetUrl(url string) error {
	if this.url != url {
		this.url = url
		if this.electronApp != nil {
			err := this.closeGameWindow()
			if err != nil {
				log.Error(err.Error())
				return err
			}
			err = this.createGameWindow(url)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
	}
	return nil
}

func (this *App) SetDevTool(dev bool) {
	this.devTool = dev
	if this.gameWindow == nil {
		return
	}
	if this.devTool {
		this.gameWindow.OpenDevTools()
	} else {
		this.gameWindow.CloseDevTools()
	}
}

func (this *App) genId() util.ID {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.sid++
	return this.sid
}

func (this *App) SessionCreator(conn lokas.IConn) lokas.ISession {
	if this.Session != nil {
		this.Session.Conn.Close()
		this.Session = nil
	}
	this.Session = NewSession(conn, this.genId(), this, WithSessionHandler(this.handler))
	go this.Session.Start()
	return this.Session
}

func (this *App) Start() {
	this.Gate.LoadCustom("0.0.0.0", this.Port, protocol.BINARY, lox.Websocket)
	this.Gate.Start()
	this.Emit("start")
	this.start()
	this.Emit("stop")
	this.config.Save()
	log.Warn("save complete")
	os.Exit(1)
}

func (this *App) mainLoop() {
	signalChan := make(chan os.Signal, 1)
	this.done = make(chan struct{})
LOOP:
	for {
		select {
		case <-this.done:
			break LOOP
		case <-signalChan:
			break LOOP
		}
	}
	close(this.done)
	this.done = nil
	this.Stop()
	log.Warnf("stop")
}

func (this *App) Resize(width,height int){
	this.gameWindow.Resize(width,height)
}

func (this *App) Clear() {

}

func (this *App) Stop() error {
	if this.done!=nil {
		this.done <- struct{}{}
	}
	log.Error("Stop1")
	err := this.stop()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Error("Stop2")
	err = this.Gate.Stop()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Error("Stop3")
	return nil
}
