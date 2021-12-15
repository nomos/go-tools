package imkey

import (
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

func (this *App) PressKey(key keys.KEY){
	this.SendKeyEvent(key,keys.KEY_EVENT_TYPE_DOWN)
}

func (this *App) ReleaseKey(key keys.KEY){
	this.SendKeyEvent(key,keys.KEY_EVENT_TYPE_UP)
}

func (this *App) SendKeyEvent(key keys.KEY,event_type keys.KEY_EVENT_TYPE){
	this.sendKeyboardEvent(key,event_type)
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
	if this.keyEventHandler!=nil {
		this.keyEventHandler(e)
	}
}
