package imkey

import (
	"errors"
	"github.com/lxn/win"
	"github.com/nomos/go-lokas/ecs"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox/flog"
	"github.com/nomos/go-lokas/util/keys"
	"github.com/nomos/go-tools/imkey/interception"
	"github.com/nomos/go-tools/imkey/kernel32"
	"github.com/nomos/go-tools/imkey/user32"
	"syscall"
	"time"
)

type app struct {
	k32              *kernel32.Kernel32DLL
	user32           *user32.User32DLL
	interception     *interception.InterceptionDLL
	keyboardListener *user32.LowLevelKeyboardEventListener
	mouseListener    *user32.LowLevelMouseEventListener
}

func (this *App) init() {

}
func (this *App) start() error {
	//return this.startUseInterception()
	return this.startUseUser32()
}

func (this *App) startUseInterception()error {
	var err error
	this.k32, err = kernel32.LoadKernel32Dll()
	if err != nil {
		log.Error("load kernel32.dll error", flog.Error(err))
	}
	process := this.k32.GetCurrentProcess()
	this.k32.SetPriorityClass(process, kernel32.HIGH_PRIORITY_CLASS)
	this.interception, err = interception.LoadInterceptionDll()
	if err != nil {
		log.Error("load interception.dll error", flog.Error(err))
	}
	return nil
}

func (this *App) startUseUser32() error {
	var err error
	this.k32, err = kernel32.LoadKernel32Dll()
	if err != nil {
		log.Error("load kernel32.dll error", flog.Error(err))
	}
	process := this.k32.GetCurrentProcess()
	this.k32.SetPriorityClass(process, kernel32.HIGH_PRIORITY_CLASS)
	this.interception, err = interception.LoadInterceptionDll()
	if err != nil {
		log.Error("load interception.dll error", flog.Error(err))
	}
	this.user32, err = user32.LoadUser32DLL()
	if err != nil {
		log.Error("load user32.dll error", flog.Error(err))
	}
	this.keyboardListener, err = user32.NewLowLevelKeyboardListener(func(event user32.LowLevelKeyboardEvent) {
		a := event.KeyboardButtonAction()
		if a == user32.WMKeyDown || a == user32.WHSystemKeyDown {
			this.emitKeyEvent(&keys.KeyEvent{
				Component: ecs.Component{},
				Code:      keys.KEY(int32(event.Struct.VkCode)),
				Event:     keys.KEY_EVENT_TYPE_DOWN,
			})
		} else if a == user32.WMKeyUp || a == user32.WMSystemKeyUp {
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
	this.mouseListener, err = user32.NewLowLevelMouseListener(func(event user32.LowLevelMouseEvent) {
		switch event.MouseButtonAction() {
		case user32.WMLButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:  keys.MOUSE_EVENT_TYPE_DOWN,
				Button: keys.MOUSE_BUTTON_LEFT,
				X:      0,
				Y:      0,
				Num:    0,
			})
		case user32.WMLButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:  keys.MOUSE_EVENT_TYPE_UP,
				Button: keys.MOUSE_BUTTON_LEFT,
				X:      0,
				Y:      0,
				Num:    0,
			})
		case user32.WMMouseMove:

			this.emitMouseEvent(&keys.MouseEvent{
				Event:  keys.MOUSE_EVENT_TYPE_MOVE,
				Button: keys.MOUSE_BUTTON_LEFT,
				X:      event.Struct.Point.X,
				Y:      event.Struct.Point.Y,
				Num:    0,
			})
		case user32.WMMouseWheel:
			if event.Struct.MouseData == 4287102976 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:  keys.MOUSE_EVENT_TYPE_SCROLL_UP,
					Button: 0,
					X:      0,
					Y:      0,
					Num:    0,
				})
			} else {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:  keys.MOUSE_EVENT_TYPE_SCROLL_DOWN,
					Button: 0,
					X:      0,
					Y:      0,
					Num:    0,
				})
			}
		case user32.WMMouseHWheel:
		case user32.WMRButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:  keys.MOUSE_EVENT_TYPE_DOWN,
				Button: keys.MOUSE_BUTTON_RIGHT,
				X:      0,
				Y:      0,
				Num:    0,
			})
		case user32.WMRButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:  keys.MOUSE_EVENT_TYPE_UP,
				Button: keys.MOUSE_BUTTON_RIGHT,
				X:      0,
				Y:      0,
				Num:    0,
			})
		case user32.WMXButtonDown:
			if event.Struct.MouseData == 65536 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:  keys.MOUSE_EVENT_TYPE_DOWN,
					Button: keys.MOUSE_BUTTON_EXTRA_1,
					X:      0,
					Y:      0,
					Num:    0,
				})
			} else if event.Struct.MouseData == 131072 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:  keys.MOUSE_EVENT_TYPE_DOWN,
					Button: keys.MOUSE_BUTTON_EXTRA_2,
					X:      0,
					Y:      0,
					Num:    0,
				})
			}
		case user32.WMXButtonUp:
			if event.Struct.MouseData == 65536 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:  keys.MOUSE_EVENT_TYPE_UP,
					Button: keys.MOUSE_BUTTON_EXTRA_1,
					X:      0,
					Y:      0,
					Num:    0,
				})
			} else if event.Struct.MouseData == 131072 {
				this.emitMouseEvent(&keys.MouseEvent{
					Event:  keys.MOUSE_EVENT_TYPE_UP,
					Button: keys.MOUSE_BUTTON_EXTRA_2,
					X:      0,
					Y:      0,
					Num:    0,
				})
			}
		case user32.WMMButtonDown:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:  keys.MOUSE_EVENT_TYPE_DOWN,
				Button: keys.MOUSE_BUTTON_MIDDLE,
				X:      0,
				Y:      0,
				Num:    0,
			})
		case user32.WMMButtonUp:
			this.emitMouseEvent(&keys.MouseEvent{
				Event:  keys.MOUSE_EVENT_TYPE_UP,
				Button: keys.MOUSE_BUTTON_MIDDLE,
				X:      0,
				Y:      0,
				Num:    0,
			})
		case user32.WMNCXButtonDown:
			log.Infof("mouse event WMNCXButtonDown:%+v%+v", event.Struct.Point, event.Struct.MouseData)
		case user32.WMNCXButtonUp:
			log.Infof("mouse event WMNCXButtonUp:%+v%+v", event.Struct.Point, event.Struct.MouseData)
		default:
			log.Warnf("unknown case", event.MouseButtonAction())
		}
	}, this.user32)

	return nil
}

func (this *App) getWindow(str string) win.HWND {
	hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr(str))
	return hwnd
}

func (this *App) getDesktopRect() (int32,int32,int32,int32){
	left,top,right,bottom,_:= this.getWindowRectHwnd(win.GetDesktopWindow())
	return left,top,right,bottom
}

func (this *App) isForegroundWindow(str string)bool{
	hwnd:=this.getWindow(str)
	activeHwnd:=win.GetForegroundWindow()
	return hwnd!=0&&activeHwnd==hwnd
}

func (this *App) getWindowRectHwnd(hwnd win.HWND) (int32,int32,int32,int32,error) {

	var rect win.RECT
	ret:=win.GetWindowRect(hwnd,&rect)
	if !ret {
		return 0,0,0,0,errors.New("windows not found")
	}
	return rect.Left,rect.Top,rect.Right,rect.Bottom,nil
}

func (this *App) getWindowRect(str string) (int32,int32,int32,int32,error) {
	hwnd:=this.getWindow(str)
	if hwnd==0 {
		return 0,0,0,0,errors.New("windows not found")
	}
	return this.getWindowRectHwnd(hwnd)
}

func (this *App) hasWindow(str string) bool {
	return this.getWindow(str) != 0
}

func (this *App) stop() error {
	this.keyboardListener.Release()
	this.mouseListener.Release()
	this.k32.Release()
	this.interception.Release()
	return nil
}

func (this *App) sendKeyboardEvent(key keys.KEY, event_type keys.KEY_EVENT_TYPE) {
	this.interception.SendKeyStoke(key, event_type)
}

func (this *App) sendMouseEvent(event *keys.MouseEvent) {

	if event.Event == keys.MOUSE_EVENT_TYPE_MOVE {
		this.interception.SendMouseMoveTo(event.X,event.Y)
	} else if event.Event == keys.MOUSE_EVENT_TYPE_MOVE_RELATIVE {
		this.interception.SendMouseMoveRelative(event.X,event.Y)
	} else if event.Event == keys.MOUSE_EVENT_TYPE_UP {

		this.interception.SendMouseButtonRelease(event.Button)
	} else if event.Event == keys.MOUSE_EVENT_TYPE_DOWN {
		this.interception.SendMouseButtonPress(event.Button)
	} else if event.Event == keys.MOUSE_EVENT_TYPE_PRESS {
		this.interception.SendMouseButtonPress(event.Button)
		time.Sleep(time.Millisecond*10)
		this.interception.SendMouseButtonRelease(event.Button)
	}
}
