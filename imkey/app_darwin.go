package imkey

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/nomos/go-lokas/util/keys"
	hook "github.com/robotn/gohook"
)

type app struct {
}

func (this *App) start() error {
	//addEvent()
	//addMouse()
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

func (this *App) init() error {
	return nil
}

func (this *App) stop() error {

	return nil
}

func (this *App) hasWindow(str string) bool {
	return false
}

func (this *App) sendKeyboardEvent(key keys.KEY, event_type keys.KEY_EVENT_TYPE) {

}

func (this *App) sendMouseEvent(event *keys.MouseEvent) {

}

func (this *App) setActiveWindow(str string) {

}

func (this *App) isActiveWindow(str string)bool{
	return true
}

func (this *App) getWindowRect(str string) (int32,int32,int32,int32,error) {
	return 0,0,0,0,nil
}


func (this *App) getDesktopRect() (int32,int32,int32,int32){
	return 0,0,0,0
}

func (this *App) isForegroundWindow(str string)bool{
	return true
}