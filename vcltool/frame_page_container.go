package vcltool

import (
	"github.com/nomos/go-events"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)



type PageContainer struct {
	log           log.ILogger
	listener      events.EventEmmiter
	iframes       map[string]IFrame
	webviewFrames map[string]*TWebViewFrame
	pageControl   *vcl.TPageControl
	component     vcl.IComponent
	num           int
}

func NewPageContainer(self vcl.IComponent, pageControl *vcl.TPageControl, log log.ILogger, listener events.EventEmmiter) *PageContainer {
	ret := &PageContainer{
		log:           log,
		listener:      listener,
		iframes:       make(map[string]IFrame),
		webviewFrames: make(map[string]*TWebViewFrame),
		pageControl:   pageControl,
		component:     self,
		num:           0,
	}

	return ret
}

func (this *PageContainer) On(evt events.EventName, listener ...events.Listener) {
	this.listener.On(evt, listener...)
}

func (this *PageContainer) Emit(evt events.EventName, args ...interface{}) {
	this.listener.Emit(evt, args...)
}

func (this *PageContainer) OnCreate() {
	this.pageControl.SetOnChanging(func(sender vcl.IObject, allowChange *bool) {
		go func() {
			vcl.ThreadSync(func() {
				index := int(this.pageControl.ActivePageIndex())
				this.listener.Emit("page_change", index)
			})
		}()

	})
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

func (this *PageContainer) Destroy() {
	for _, frame := range this.iframes {
		frame.OnDestroy()
		frame.Free()
	}
	this.iframes = make(map[string]IFrame)
	for _, frame := range this.webviewFrames {
		frame.OnDestroy()
		frame.Free()
	}
	this.webviewFrames = make(map[string]*TWebViewFrame)
}

func (this *PageContainer) GetIFrame(s string) IFrame {
	return this.iframes[s]
}

type FrameOption func(frame IFrame)

func WithConfig(conf lokas.IConfig) FrameOption {
	return func(frame IFrame) {
		frame.SetConfig(conf)
	}
}

func WithContent(s string, data interface{}) FrameOption {
	return func(frame IFrame) {
		frame.SetContent(s, data)
	}
}

func (this *PageContainer) AddIFrame(name string, frame IFrame, opts ...FrameOption) IFrame {
	for _, opt := range opts {
		opt(frame)
	}
	frame.SetEventEmitter(this.listener)
	frame.SetLogger(this.log)
	frame.setSheetName(name)
	frame.setContainer(this)
	this.iframes[name] = frame
	sheet := vcl.NewTabSheet(this.component)
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
	sheet := vcl.NewTabSheet(this.component)
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
