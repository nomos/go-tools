package imkey

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/nomos/go-lokas/util/keys"
	hook "github.com/robotn/gohook"
)

func (this *App) start()error{
	addEvent()
	addMouse()
	return nil
}

func addEvent() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		robotgo.EventEnd()
	})

	fmt.Println("--- Please press w---")
	robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("w")
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func addMouse() {
	fmt.Println("--- Please press left mouse button to see it's position and the right mouse button to exit ---")
	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			fmt.Printf("mouse left @ %v - %v\n", e.X, e.Y)
		} else if e.Button == hook.MouseMap["right"] {
			robotgo.EventEnd()
		}
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func (this *App) stop()error{

	return nil
}

func (this *App) sendKeyboardEvent(event *keys.KeyEvent){

}

func (this *App) sendMouseEvent(event *keys.MouseEvent){

}

func (this *App) onKeyboardEvent(event *keys.KeyEvent){

}

func (this *App) onMouseEvent(event *keys.MouseEvent){

}