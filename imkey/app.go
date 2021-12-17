package imkey

import (
	"github.com/kbinani/screenshot"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/colors"
	"github.com/nomos/go-lokas/util/events"
	"github.com/nomos/go-lokas/util/keys"
	"image"
	"sync"
	"time"
)

var instance *App
var once sync.Once

func Instance() *App {
	once.Do(func() {
		if instance == nil {
			instance = &App{
				EventEmmiter:events.New(),
				app: &app{},
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
	taskMutex         sync.Mutex
	resetMutex        sync.Mutex
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

func (this *Task) Sleep(duration time.Duration) {
	iterNum := duration / time.Millisecond
	var i time.Duration = 0
	for i = 0; i < iterNum; i++ {
		if this.taskOn {
			return
		}
		time.Sleep(time.Millisecond)
	}
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
	this.resetMutex.Lock()
	defer this.resetMutex.Unlock()
	for k, v := range this.keyStatus {
		if v {
			this.sendKeyEvent(k, keys.KEY_EVENT_TYPE_UP)
		}
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
	this.sendKeyEvent(key, keys.KEY_EVENT_TYPE_DOWN)
}

func (this *App) ReleaseKey(key keys.KEY) {
	if !this.enabled {
		return
	}
	this.sendKeyEvent(key, keys.KEY_EVENT_TYPE_UP)
}

func (this *App) IsKeyPressed(key keys.KEY) bool {
	return this.keyStatus[key]
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

func (this *App) AddTask(name string,task ITask){
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
	go task.Run()
}

func (this *App) ScreenShot()(*image.RGBA,error){
	return screenshot.CaptureDisplay(0)
}

func (this *App) ScreenCapture(x,y,w,h int)(*image.RGBA,error){
	return screenshot.Capture(x,y,w,h)
}

func (this *App) GetScreenPixel(x,y int)(*colors.Color,error){
	img,err:=screenshot.Capture(x,y,1,1)
	var ret colors.Color
	if err!=nil {
		return nil,err
	}
	ret = colors.NewColorRGBA(img.Pix[0],img.Pix[1],img.Pix[2],img.Pix[3])
	return &ret,nil
}

func (this *App) StopTask(name string) {
	this.taskMutex.Lock()
	defer this.taskMutex.Unlock()
	t := this.tasks[name]
	t.TaskOff()
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

func (this *App) HasWindow(str string) bool {
	return this.hasWindow(str)
}

func (this *App) SetActiveWindow(str string){
	this.setActiveWindow(str)
}

func (this *App) IsActiveWindow(str string)bool{
	return this.isActiveWindow(str)
}

func (this *App) GetWindowRect(str string) (int32,int32,int32,int32,error) {
	return this.getWindowRect(str)
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

func (this *App) emitMouseEvent(e *keys.MouseEvent) {
	if !this.enabled {
		return
	}
	if this.mouseEventHandler != nil {
		this.mouseEventHandler(e)
	}
}

func (this *App) emitKeyEvent(e *keys.KeyEvent) {
	if e.Event == keys.KEY_EVENT_TYPE_UP {
		this.keyStatus[e.Code] = false
	} else {
		this.keyStatus[e.Code] = true
	}
	if !this.enabled {
		return
	}
	if this.keyEventHandler != nil {
		this.keyEventHandler(e)
	}
}
