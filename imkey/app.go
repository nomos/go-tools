package imkey

import (
	"github.com/kbinani/screenshot"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-lokas/util/colors"
	"github.com/nomos/go-lokas/util/events"
	"github.com/nomos/go-lokas/util/keys"
	"image"
	"sync"
	"syscall"
)

var instance *App
var once sync.Once

func Instance() *App {
	once.Do(func() {
		if instance == nil {
			instance = &App{
				EventEmmiter: events.New(),
				tasks:        map[string]ITask{},
				keyStatus:    map[keys.KEY]bool{},
				mouseStatus: map[keys.MOUSE_BUTTON]bool{},
				app:          &app{},
			}
			instance.Init()
		}
	})
	return instance
}

type App struct {
	events.EventEmmiter
	*app
	keyEventHandler   KeyEventHandler
	mouseEventHandler MouseEventHandler
	tasks             map[string]ITask
	enabled           bool
	keyStatus         map[keys.KEY]bool
	mouseStatus       map[keys.MOUSE_BUTTON]bool
	taskMutex         sync.Mutex
	resetMutex        sync.Mutex
	mouseX            int32
	mouseY            int32
}

type KeyEventHandler func(event *keys.KeyEvent)
type MouseEventHandler func(event *keys.MouseEvent)

type Task struct {
	*App
	Name     string
	taskOn   bool
	taskFunc TaskFunc
}

func NewTask(name string, app *App, taskFunc TaskFunc) *Task {
	ret := &Task{
		App:      app,
		Name:     name,
		taskOn:   false,
		taskFunc: taskFunc,
	}
	return ret
}

func (this *Task) Sleep(duration int64) bool {
	var tv syscall.Timeval
	_ = syscall.Gettimeofday(&tv)
	util.SleepUtil(duration, func() bool {
		return !this.taskOn
	})
	startTick := int64(tv.Sec)*int64(1000000) + int64(tv.Usec) + int64(duration)*1000
	endTick := int64(0)
	for endTick < startTick {
		if !this.taskOn {
			return true
		}
		_ = syscall.Gettimeofday(&tv)
		endTick = int64(tv.Sec)*int64(1000000) + int64(tv.Usec)
	}
	return false
}

func (this *Task) Run() {
	this.taskOn = false
	this.taskFunc(this)
	this.taskOn = true
	this.App.RemoveTask(this.Name)
}

func (this *Task) TaskOff() {
	this.taskOn = false
}

type TaskFunc func(task *Task)

func (this *App) Init() {
	this.keyStatus = map[keys.KEY]bool{}
	this.init()
}

func (this *App) Start() error {
	err := this.start()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.enabled = true
	return nil
}

func (this *App) ResetAllKeys() {
	ks := []keys.KEY{}
	this.resetMutex.Lock()
	for k, v := range this.keyStatus {
		if v {
			ks = append(ks, k)
		}
	}
	this.resetMutex.Unlock()
	for _, v := range ks {
		this.sendKeyEvent(v, keys.KEY_EVENT_TYPE_UP)
	}
	for _, b := range keys.ALL_MOUSE_BUTTON {
		this.ReleaseMouseButton(keys.MOUSE_BUTTON(b.Enum()))
	}
}

func (this *App) Stop() error {
	this.enabled = false
	this.StopAllTask()
	this.ResetAllKeys()
	return this.stop()
}

func (this *App) PressKey(key keys.KEY) {
	if !this.enabled {
		return
	}
	go this.sendKeyEvent(key, keys.KEY_EVENT_TYPE_DOWN)
}

func (this *App) ReleaseKey(key keys.KEY) {
	if !this.enabled {
		return
	}
	go this.sendKeyEvent(key, keys.KEY_EVENT_TYPE_UP)
}

func (this *App) ClickKey(key keys.KEY) {
	if !this.enabled {
		return
	}
	go func() {
		go this.sendKeyEvent(key, keys.KEY_EVENT_TYPE_DOWN)
		util.Sleep(50)
		this.sendKeyEvent(key, keys.KEY_EVENT_TYPE_UP)
	}()
}

func (this *App) IsKeyPressed(key keys.KEY) bool {
	this.resetMutex.Lock()
	defer this.resetMutex.Unlock()
	return this.keyStatus[key]
}

func (this *App) IsMousePressed(btn keys.MOUSE_BUTTON) bool {
	this.resetMutex.Lock()
	defer this.resetMutex.Unlock()
	return this.mouseStatus[btn]
}

func (this *App) sendKeyEvent(key keys.KEY, event_type keys.KEY_EVENT_TYPE) {
	this.sendKeyboardEvent(key, event_type)
}

func (this *App) RunTask(name string, taskFunc TaskFunc) {
	if !this.enabled {
		return
	}
	this.taskMutex.Lock()
	t := this.tasks[name]
	if t != nil {
		this.taskMutex.Unlock()
		return
	}
	t = NewTask(name, this, taskFunc)
	this.tasks[name] = t
	go t.Run()
}

func (this *App) AddTask(name string, task ITask) {
	if !this.enabled {
		return
	}
	this.taskMutex.Lock()
	t := this.tasks[name]
	if t != nil {
		this.taskMutex.Unlock()
		return
	}
	this.tasks[name] = task
	this.taskMutex.Unlock()
	go task.Run()
}

func (this *App) ScreenShot() (*image.RGBA, error) {
	return screenshot.CaptureDisplay(0)
}

func (this *App) ScreenCapture(x, y, w, h int32) (*image.RGBA, error) {
	return screenshot.Capture(int(x), int(y), int(w), int(h))
}

func (this *App) ScreenPixel(x, y int32) (*colors.Color, error) {
	img, err := screenshot.Capture(int(x), int(y), 1, 1)
	var ret colors.Color
	if err != nil {
		return nil, err
	}
	ret = colors.NewColorRGBA(img.Pix[0], img.Pix[1], img.Pix[2], img.Pix[3])
	return &ret, nil
}

func (this *App) StopTask(name string) {
	this.taskMutex.Lock()
	defer this.taskMutex.Unlock()
	t := this.tasks[name]
	if t != nil {
		t.TaskOff()
	}
}

func (this *App) StopAllTask() {
	this.taskMutex.Lock()
	for _, v := range this.tasks {
		v.TaskOff()
	}
	this.taskMutex.Unlock()
	for {
		if len(this.tasks) == 0 {
			break
		}
	}
}

func (this *App) MoveMouseTo(x, y int32) {
	if !this.enabled {
		return
	}
	deltaX := x - this.mouseX
	deltaY := y - this.mouseY
	this.MoveMouseRelative(deltaX, deltaY)
}

func (this *App) PressMouseButton(button keys.MOUSE_BUTTON) {
	if !this.enabled {
		return
	}
	e := keys.NewMouseEvent()
	e.Event = keys.MOUSE_EVENT_TYPE_DOWN
	e.Button = button
	this.sendMouseEvent(e)
}

func (this *App) ReleaseMouseButton(button keys.MOUSE_BUTTON) {
	if !this.enabled {
		return
	}
	e := keys.NewMouseEvent()
	e.Event = keys.MOUSE_EVENT_TYPE_UP
	e.Button = button
	this.sendMouseEvent(e)
}

func (this *App) ClickMouseButton(button keys.MOUSE_BUTTON) {
	if !this.enabled {
		return
	}
	e := keys.NewMouseEvent()
	e.Event = keys.MOUSE_EVENT_TYPE_PRESS
	e.Button = button
	this.sendMouseEvent(e)
}

func (this *App) MoveMouseRelative(x, y int32) {
	e := keys.NewMouseEvent()
	e.Event = keys.MOUSE_EVENT_TYPE_MOVE_RELATIVE
	e.X = x
	e.Y = y
	this.sendMouseEvent(e)
}

func (this *App) HasWindow(str string) bool {
	return this.hasWindow(str)
}

func (this *App) IsActiveWindow(str string) bool {
	return this.isForegroundWindow(str)
}

func (this *App) GetWindowRect(str string) (int32, int32, int32, int32, error) {
	return this.getWindowRect(str)
}

func (this *App) GetDesktopRect() (int32, int32, int32, int32) {
	return this.getDesktopRect()
}

func (this *App) RemoveTask(name string) {
	this.taskMutex.Lock()
	defer this.taskMutex.Unlock()
	delete(this.tasks, name)
}

func (this *App) GetTask(name string) ITask {
	this.taskMutex.Lock()
	defer this.taskMutex.Unlock()
	return this.tasks[name]
}

func (this *App) SendMouseEvent(event *keys.MouseEvent) {
	if !this.enabled {
		return
	}
	this.sendMouseEvent(event)
}

func (this *App) SetOnKeyboardEvent(f KeyEventHandler) {
	this.keyEventHandler = f
}

func (this *App) SetOnMouseEvent(f MouseEventHandler) {
	this.mouseEventHandler = f
}

func (this *App) GetMouseX() int32 {
	return this.mouseX
}

func (this *App) GetMouseY() int32 {
	return this.mouseY
}

func (this *App) emitMouseEvent(e *keys.MouseEvent) {
	if e.Event == keys.MOUSE_EVENT_TYPE_UP {
		this.setMouseStatus(e.Button, false)
	} else if e.Event == keys.MOUSE_EVENT_TYPE_DOWN  {
		this.setMouseStatus(e.Button, true)
	}
	if e.Event == keys.MOUSE_EVENT_TYPE_MOVE {
		this.mouseX = e.X
		this.mouseY = e.Y
	}
	if !this.enabled {
		return
	}
	if this.mouseEventHandler != nil {
		this.mouseEventHandler(e)
	}
}

func (this *App) setKeyStatus(key keys.KEY, v bool)bool {
	this.resetMutex.Lock()
	defer this.resetMutex.Unlock()
	if this.keyStatus[key] != v {
		this.keyStatus[key] = v
		return true
	}
	return false
}

func (this *App) setMouseStatus(key keys.MOUSE_BUTTON, v bool)bool {
	this.resetMutex.Lock()
	defer this.resetMutex.Unlock()
	if this.mouseStatus[key] != v {
		this.mouseStatus[key] = v
		return true
	}
	return false
}
func (this *App) clearKeyStatus(key keys.KEY, v bool) {
	this.resetMutex.Lock()
	defer this.resetMutex.Unlock()
	this.keyStatus[key] = v
}

func (this *App) emitKeyEvent(e *keys.KeyEvent) {
	if e.Event == keys.KEY_EVENT_TYPE_UP {
		if this.setKeyStatus(e.Code, false) {
			if this.keyEventHandler != nil {
				this.keyEventHandler(e)
			}
		}
	} else {
		if this.setKeyStatus(e.Code, true) {
			if !this.enabled {
				return
			}
			if this.keyEventHandler != nil {
				this.keyEventHandler(e)
			}
		}
	}

}
