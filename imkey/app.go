package imkey

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/keys"
	"sync"
	"time"
)

var instance *App
var once sync.Once

func Instance() *App {
	once.Do(func() {
		if instance == nil {
			instance = &App{
				app: &app{},
			}
			instance.Init()
		}
	})
	return instance
}

type App struct {
	*app
	keyEventHandler   KeyEventHandler
	mouseEventHandler MouseEventHandler
	tasks             map[string]*Task
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
	this.App.removeTask(this.Name)
}

type TaskFunc func(task *Task)

func (this *App) Init() {
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
	this.stopAllTask()
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

func (this *App) StopTask(name string) {
	this.taskMutex.Lock()
	defer this.taskMutex.Unlock()
	t := this.tasks[name]
	t.taskOn = false
}

func (this *App) stopAllTask() {
	this.taskMutex.Lock()
	for _, v := range this.tasks {
		v.taskOn = false
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

func (this *App) removeTask(name string) {
	this.taskMutex.Lock()
	defer this.taskMutex.Unlock()
	delete(this.tasks, name)
}

func (this *App) getTask(name string) *Task {
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
