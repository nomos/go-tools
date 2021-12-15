package imkey

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox/flog"
	"github.com/nomos/go-lokas/util/keys"
	"github.com/nomos/go-tools/imkey/user32util"
)

type app struct {
	user32 *user32util.User32DLL
	keyboardListener *user32util.LowLevelKeyboardEventListener
	mouseListener *user32util.LowLevelMouseEventListener
}

func (this *App) init(){

}

func (this *App) start()error{
	var err error
	this.user32, err = user32util.LoadUser32DLL()
	if err != nil {
		log.Error("load dll error",flog.Error(err))
		// Error handling.
	}
	this.keyboardListener, err = user32util.NewLowLevelKeyboardListener(func(event user32util.LowLevelKeyboardEvent) {
		a:=event.KeyboardButtonAction()
		if a==user32util.WMKeyDown||a==user32util.WHSystemKeyDown {

			log.Infof("KeyDown event:%+v",event.Struct.VkCode,keys.KEY(int32(event.Struct.VkCode)).ToString())
		} else if a==user32util.WMKeyUp||a==user32util.WMSystemKeyUp {

			log.Infof("KeyUp event:%+v",event.Struct.VkCode)
		}
	}, this.user32)
	if err != nil {
		log.Fatalf("failed to create listener - %s", err.Error())
	}
	this.mouseListener,err = user32util.NewLowLevelMouseListener(func(event user32util.LowLevelMouseEvent) {
		switch event.MouseButtonAction() {
		case user32util.WMLButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_DOWN,
				Button:    keys.MOUSE_BUTTON_LEFT,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32util.WMLButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_UP,
				Button:    keys.MOUSE_BUTTON_LEFT,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32util.WMMouseMove:

		case user32util.WMMouseWheel:
			if event.Struct.MouseData==4287102976 {
				log.Infof("scroll",event.Struct.MouseData)
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
		case user32util.WMMouseHWheel:
		case user32util.WMRButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_DOWN,
				Button:    keys.MOUSE_BUTTON_RIGHT,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32util.WMRButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_UP,
				Button:    keys.MOUSE_BUTTON_RIGHT,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32util.WMXButtonDown:
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
		case user32util.WMXButtonUp:
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
		case user32util.WMMButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_DOWN,
				Button:    keys.MOUSE_BUTTON_MIDDLE,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32util.WMMButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:     keys.MOUSE_EVENT_TYPE_UP,
				Button:    keys.MOUSE_BUTTON_MIDDLE,
				X:         0,
				Y:         0,
				Num:       0,
			})
		case user32util.WMNCXButtonDown:
			log.Infof("mouse event WMNCXButtonDown:%+v%+v",event.Struct.Point,event.Struct.MouseData)
		case user32util.WMNCXButtonUp:
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