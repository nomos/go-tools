package imkey

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/keys"
	"sync"
)

var instance *App
var once sync.Once

func Instance() *App {
	once.Do(func() {
		if instance == nil {
			instance = &App{
				app:&app{},
			}
			instance.Init()
		}
	})
	return instance
}

type App struct {
	*app
	keyEventHandler KeyEventHandler
	mouseEventHandler MouseEventHandler
}

type KeyEventHandler func(event *keys.KeyEvent)
type MouseEventHandler func(event *keys.MouseEvent)

func (this *App) Init(){
	this.init()
}

func (this *App) Start()error{
	return this.start()
}

func (this *App) Stop()error{
	return this.stop()
}

func (this *App) SendKeyboardEvent(event *keys.KeyEvent){
	this.sendKeyboardEvent(event)
}

func (this *App) SendMouseEvent(event *keys.MouseEvent){
	this.sendMouseEvent(event)
}

func (this *App) SetOnKeyboardEvent(f KeyEventHandler){
	this.keyEventHandler = f
}

func (this *App) SetOnMouseEvent(f MouseEventHandler){
	this.mouseEventHandler = f
}

func (this *App) emitMouseEvent(e *keys.MouseEvent){
	if this.mouseEventHandler!=nil {
		this.mouseEventHandler(e)
	}
}

func (this *App) emitKeyEvent(e *keys.KeyEvent){
	log.Warnf(e.Event.ToString(),e.Code.ToString())
	if this.keyEventHandler!=nil {
		this.keyEventHandler(e)
	}
}



