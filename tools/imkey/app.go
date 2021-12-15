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
			instance = &App{}
		}
	})
	return instance
}

type App struct {

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

func (this *App) OnKeyboardEvent(event *keys.KeyEvent){
	this.onKeyboardEvent(event)
}

func (this *App) OnMouseEvent(event *keys.MouseEvent){
	this.onMouseEvent(event)
}


