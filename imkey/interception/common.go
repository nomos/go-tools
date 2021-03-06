package interception

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/keys"
	"golang.org/x/sys/windows"
	"unsafe"
)

const (
	interceptionDllName         = "interception.dll"
	MOUSE               uintptr = 11
	KEYBOARD            uintptr = 1
)

func LoadInterceptionDll() (*InterceptionDLL, error) {
	temp := windows.LazyDLL{
		Name:   interceptionDllName,
		System: false,
	}
	interception := &windows.DLL{
		Name:   temp.Name,
		Handle: windows.Handle(temp.Handle()),
	}
	interception_is_keyboard, err := interception.FindProc("interception_is_keyboard")
	if err != nil {
		return nil, err
	}
	interception_is_mouse, err := interception.FindProc("interception_is_mouse")
	if err != nil {
		return nil, err
	}
	interception_create_context, err := interception.FindProc("interception_create_context")
	if err != nil {
		return nil, err
	}
	interception_set_filter, err := interception.FindProc("interception_set_filter")
	if err != nil {
		return nil, err
	}
	interception_receive, err := interception.FindProc("interception_receive")
	if err != nil {
		return nil, err
	}
	interception_wait, err := interception.FindProc("interception_wait")
	if err != nil {
		return nil, err
	}
	interception_send, err := interception.FindProc("interception_send")
	if err != nil {
		return nil, err
	}
	interception_destroy_context, err := interception.FindProc("interception_destroy_context")
	if err != nil {
		return nil, err
	}
	ret := &InterceptionDLL{
		interception:                 interception,
		interception_is_keyboard:     interception_is_keyboard,
		interception_is_mouse:        interception_is_mouse,
		interception_create_context:  interception_create_context,
		interception_set_filter:      interception_set_filter,
		interception_receive:         interception_receive,
		interception_wait:            interception_wait,
		interception_send:            interception_send,
		interception_destroy_context: interception_destroy_context,
	}
	err = ret.init()
	if err != nil {
		log.Error(err.Error())
	}
	return ret, nil

}

type InterceptionDLL struct {
	interception                 *windows.DLL
	context                      uintptr
	interception_is_keyboard     *windows.Proc
	interception_is_mouse        *windows.Proc
	interception_create_context  *windows.Proc
	interception_set_filter      *windows.Proc
	interception_receive         *windows.Proc
	interception_wait            *windows.Proc
	interception_send            *windows.Proc
	interception_destroy_context *windows.Proc
}

func (this *InterceptionDLL) init() error {
	var err error
	this.context, _, err = this.interception_create_context.Call()
	if err != nil {
		//log.Error(err.Error())
	}
	return nil
}

func (this *InterceptionDLL) Release() error {
	this.interception_destroy_context.Call(this.context)
	return this.interception.Release()
}

func (this *InterceptionDLL) SendMouseMoveRelative(x, y int32) error {
	stroke := &InterceptionMouseStroke{
		Code:        0,
		State:       uint16(INTERCEPTION_MOUSE_MOVE_RELATIVE),
		Rolling:     0,
		X:           x,
		Y:           y,
		Information: 0,
	}
	this.interception_send.Call(this.context, MOUSE, uintptr(unsafe.Pointer(stroke)), 1)
	return nil
}
func (this *InterceptionDLL) SetFilter() {
	this.interception_set_filter.Call(this.context, this.interception_is_keyboard.Addr())
	this.interception_set_filter.Call(this.context, this.interception_is_mouse.Addr())
}

func (this *InterceptionDLL) SendMouseMoveTo(x, y int32) error {
	stroke := &InterceptionMouseStroke{
		Code:        0,
		State:       uint16(INTERCEPTION_MOUSE_MOVE_ABSOLUTE),
		Rolling:     0,
		X:           0,
		Y:           0,
		Information: 0,
	}
	this.interception_send.Call(this.context, MOUSE, uintptr(unsafe.Pointer(stroke)), 1)
	return nil
}

func (this *InterceptionDLL) SendMouseButtonPress(button keys.MOUSE_BUTTON) error {
	stroke := &InterceptionMouseStroke{
		Code:        0,
		State:       uint16(INTERCEPTION_MOUSE_CUSTOM),
		Rolling:     0,
		X:           0,
		Y:           0,
		Information: 0,
	}
	switch button {
	case keys.MOUSE_BUTTON_LEFT:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_1_DOWN)
	case keys.MOUSE_BUTTON_RIGHT:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_2_DOWN)
	case keys.MOUSE_BUTTON_MIDDLE:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_3_DOWN)
	case keys.MOUSE_BUTTON_EXTRA_1:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_4_DOWN)
	case keys.MOUSE_BUTTON_EXTRA_2:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_5_DOWN)
	}
	this.interception_send.Call(this.context, MOUSE, uintptr(unsafe.Pointer(stroke)), 1)
	return nil
}

func (this *InterceptionDLL) SendMouseButtonRelease(button keys.MOUSE_BUTTON) error {
	stroke := &InterceptionMouseStroke{
		Code:        0,
		State:       uint16(INTERCEPTION_MOUSE_CUSTOM),
		Rolling:     0,
		X:           0,
		Y:           0,
		Information: 0,
	}
	switch button {
	case keys.MOUSE_BUTTON_LEFT:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_1_UP)
	case keys.MOUSE_BUTTON_RIGHT:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_2_UP)
	case keys.MOUSE_BUTTON_MIDDLE:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_3_UP)
	case keys.MOUSE_BUTTON_EXTRA_1:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_4_UP)
	case keys.MOUSE_BUTTON_EXTRA_2:
		stroke.Code = uint16(INTERCEPTION_MOUSE_BUTTON_5_UP)
	}
	this.interception_send.Call(this.context, MOUSE, uintptr(unsafe.Pointer(stroke)), 1)
	return nil
}

func (this *InterceptionDLL) SendKeyStoke(code keys.KEY, event_type keys.KEY_EVENT_TYPE) error {
	stroke := &InterceptionKeyStroke{
		code:        uint16(code.ScanCode()),
		information: 0,
	}
	switch event_type {
	case keys.KEY_EVENT_TYPE_DOWN:
		stroke.state = uint16(INTERCEPTION_KEY_DOWN)
	case keys.KEY_EVENT_TYPE_UP:
		stroke.state = uint16(INTERCEPTION_KEY_UP)
	default:
		log.Warnf("undefined event", event_type)
		return nil
	}
	this.interception_send.Call(this.context, KEYBOARD, uintptr(unsafe.Pointer(stroke)), 1)
	return nil
}

func (this *InterceptionDLL) Receive(device uintptr, stroke *InterceptionKeyStroke) uintptr {
	ret, _, _ := this.interception_receive.Call(this.context, device, uintptr(unsafe.Pointer(stroke)), 1)
	return ret
}

func (this *InterceptionDLL) GetContext() uintptr {
	return this.context
}

func (this *InterceptionDLL) Wait() uintptr {
	ret, _, err := this.interception_wait.Call(this.context)
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("ret", ret)
	return ret
}
