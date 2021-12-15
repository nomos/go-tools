package imkey

import (
	"github.com/nomos/go-lokas/ecs"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox/flog"
	"github.com/nomos/go-lokas/util/keys"
	"github.com/nomos/go-tools/imkey/user32"
)

type app struct {
	user32 *user32.User32DLL
	keyboardListener *user32.LowLevelKeyboardEventListener
	mouseListener *user32.LowLevelMouseEventListener
}

func (this *App) init(){

}

func (this *App) start()error{
	var err error
	this.user32, err = user32.LoadUser32DLL()
	if err != nil {
		log.Error("load dll error",flog.Error(err))
		// Error handling.
	}
	this.keyboardListener, err = user32.NewLowLevelKeyboardListener(func(event user32.LowLevelKeyboardEvent) {
		a:=event.KeyboardButtonAction()
		if a== user32.WMKeyDown||a== user32.WHSystemKeyDown {
			this.emitKeyEvent(&keys.KeyEvent{
				Component: ecs.Component{},
				Code:      keys.KEY(int32(event.Struct.VkCode)),
				Event:     keys.KEY_EVENT_TYPE_DOWN,
			})
		} else if a== user32.WMKeyUp||a== user32.WMSystemKeyUp {
			this.emitKeyEvent(&keys.KeyEvent{
				Component: ecs.Component{},
				Code:      keys.KEY(int32(event.Struct.VkCode)),
				Event:     keys.KEY_EVENT_TYPE_UP,
			})
		}
	}, this.user32)
	if err != nil {
		log.Fatalf("failed to create listener - %s", err.Error())
	}
	this.mouseListener,err = user32.NewLowLevelMouseListener(func(event user32.LowLevelMouseEvent) {
		switch event.MouseButtonAction() {
		case user32.WMLButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_DOWN,
				Button:    keys.MOUSE_BUTTON_LEFT,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32.WMLButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_UP,
				Button:    keys.MOUSE_BUTTON_LEFT,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32.WMMouseMove:

		case user32.WMMouseWheel:
			if event.Struct.MouseData==4287102976 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:     keys.MOUSE_EVENT_TYPE_SCROLL_UP,
					Button:    0,
					X:         0,
					Y:         0,
					Num:       0,
				})
			} else {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:     keys.MOUSE_EVENT_TYPE_SCROLL_DOWN,
					Button:    0,
					X:         0,
					Y:         0,
					Num:       0,
				})
			}
		case user32.WMMouseHWheel:
		case user32.WMRButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_DOWN,
				Button:    keys.MOUSE_BUTTON_RIGHT,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32.WMRButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_UP,
				Button:    keys.MOUSE_BUTTON_RIGHT,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32.WMXButtonDown:
			if event.Struct.MouseData==65536 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:     keys.MOUSE_EVENT_TYPE_DOWN,
					Button:    keys.MOUSE_BUTTON_EXTRA_1,
					X:         0,
					Y:         0,
					Num:       0,
				})
			} else if event.Struct.MouseData == 131072 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:     keys.MOUSE_EVENT_TYPE_DOWN,
					Button:    keys.MOUSE_BUTTON_EXTRA_2,
					X:         0,
					Y:         0,
					Num:       0,
				})
			}
		case user32.WMXButtonUp:
			if event.Struct.MouseData==65536 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:     keys.MOUSE_EVENT_TYPE_UP,
					Button:    keys.MOUSE_BUTTON_EXTRA_1,
					X:         0,
					Y:         0,
					Num:       0,
				})
			} else if event.Struct.MouseData == 131072 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:     keys.MOUSE_EVENT_TYPE_UP,
					Button:    keys.MOUSE_BUTTON_EXTRA_2,
					X:         0,
					Y:         0,
					Num:       0,
				})
			}
		case user32.WMMButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_DOWN,
				Button:    keys.MOUSE_BUTTON_MIDDLE,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32.WMMButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_UP,
				Button:    keys.MOUSE_BUTTON_MIDDLE,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32.WMNCXButtonDown:
			log.Infof("mouse event WMNCXButtonDown:%+v%+v",event.Struct.Point,event.Struct.MouseData)
		case user32.WMNCXButtonUp:
			log.Infof("mouse event WMNCXButtonUp:%+v%+v",event.Struct.Point,event.Struct.MouseData)
		default:
			log.Warnf("unknown case",event.MouseButtonAction())
		}
	},this.user32)

	return nil
}

func (this *App) stop()error{
	this.keyboardListener.Release()
	this.mouseListener.Release()
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