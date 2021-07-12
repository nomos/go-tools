package ui

import (
	"github.com/nomos/go-events"
	"github.com/nomos/go-lokas/log"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"runtime"
)


type PageContainer struct {
	*vcl.TFrame
	ConfigAble

	log           log.ILogger
	listener      events.EventEmmiter
	iframes       map[string]IFrame
	webviewFrames map[string]*WebViewFrame
	pageControl   *vcl.TPageControl
	component     vcl.IComponent
	num           int
}

func NewPageContainer(owner vcl.IComponent,option... FrameOption) (root *PageContainer) {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *PageContainer) setup(){
	this.SetAlign(types.AlClient)
	this.SetColor(colors.ClSysDefault)
	this.iframes = make(map[string]IFrame)
	this.webviewFrames = make(map[string]*WebViewFrame)
	this.pageControl = vcl.NewPageControl(this)
	this.pageControl.SetParent(this)
	this.pageControl.SetAlign(types.AlClient)
	this.listener = events.New()
}

func (this *PageContainer) SetLogger(log log.ILogger){
	this.log = log
}

func (this *PageContainer) SetListener(listener events.EventEmmiter){
	this.listener = listener
}

func (this *PageContainer) OnCreate() {
	this.setup()
	this.pageControl.SetOnChanging(func(sender vcl.IObject, allowChange *bool) {
		go func() {
			vcl.ThreadSync(func() {
				index := int(this.pageControl.ActivePageIndex())
				this.listener.Emit("page_change", index)
			})
		}()

	})
}

func (this *PageContainer) On(evt events.EventName, listener ...events.Listener) {
	this.listener.On(evt, listener...)
}

func (this *PageContainer) Emit(evt events.EventName, args ...interface{}) {
	this.listener.Emit(evt, args...)
}

func (this *PageContainer) IsFrameSelected(frame IFrame) bool {
	control := this.pageControl.Controls(this.pageControl.ActivePageIndex())
	if control == nil {
		return false
	}
	sheet := vcl.AsTabSheet(control)
	if sheet == nil {
		return false
	}
	return sheet.Name() == frame.Name()+"Sheet"
}

func (this *PageContainer) OnDestroy() {
	for _, frame := range this.iframes {
		frame.OnDestroy()
		frame.Free()
	}
	this.iframes = make(map[string]IFrame)
	for _, frame := range this.webviewFrames {
		frame.OnDestroy()
		frame.Free()
	}
	this.webviewFrames = make(map[string]*WebViewFrame)
}

func (this *PageContainer) GetIFrame(s string) IFrame {
	return this.iframes[s]
}


func (this *PageContainer) AddIFrame(name string, frame IFrame, opts ...FrameOption) IFrame {
	defer func() {
		if r := recover(); r != nil {
			if err,ok:=r.(error);ok {
				log.Error(err.Error())
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, true)
				log.Error(string(buf))
			}
		}
	}()
	for _, opt := range opts {
		opt(frame)
	}
	frame.SetEventEmitter(this.listener)
	frame.setSheetName(name)
	frame.setContainer(this)
	this.iframes[name] = frame
	sheet := vcl.NewTabSheet(this)
	sheet.SetParent(this.pageControl)
	frame.SetIndex(this.num)
	this.num++
	sheet.SetName(frame.Name() + "Sheet")
	sheet.SetCaption(name)
	sheet.SetAlign(types.AlClient)
	frame.SetParent(sheet)
	frame.OnCreate()
	return frame
}

func (this *PageContainer) AddWebView(name string, url string) {
	sheet := vcl.NewTabSheet(this)
	sheet.SetParent(this.pageControl)
	sheet.SetName(name + "Sheet")
	sheet.SetCaption(name)
	sheet.SetAlign(types.AlClient)
	frame := NewWebViewFrame(sheet)
	frame.setContainer(this)
	this.iframes[name] = frame
	frame.SetIndex(this.num)
	this.num++
	frame.SetParent(sheet)
	frame.OnCreate()
	frame.SetUrl(url)
}
