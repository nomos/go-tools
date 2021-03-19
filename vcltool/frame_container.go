package vcltool

import (
	"github.com/nomos/go-events"
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/util"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type IConfigAble interface {
	Config()*util.AppConfig
	SheetName()string
	SetConfig(config *util.AppConfig)
	SetContent(s string,data interface{})
	GetContent(s string)interface{}
	setSheetName(s string)
	setContainer( IFrameContainer)
}

type ConfigAble struct {
	content map[string]interface{}
	conf *util.AppConfig
	log log.ILogger
	sheetName string
	listener events.EventEmmiter
	container IFrameContainer
	index int
}

func (this *ConfigAble)SetContent(s string,data interface{}){
	if this.content==nil {
		this.content = map[string]interface{}{}
	}
	this.content[s] = data
}

func (this *ConfigAble)GetContent(s string)interface{}{
	if this.content==nil {
		this.content = map[string]interface{}{}
	}
	return this.content[s]
}

func (this *ConfigAble) setContainer(container IFrameContainer) {
	this.container = container
}

func (this *ConfigAble) SetEventEmitter(listener events.EventEmmiter) {
	this.listener = listener
}


func (this *ConfigAble)SetIndex(num int) {
	this.index = num
}


func (this *ConfigAble)GetIndex()int {
	return this.index
}

func (this *ConfigAble) SetLogger(log log.ILogger) {
	this.log = log
}

func (this *ConfigAble) SetConfig(config *util.AppConfig) {
	this.conf = config
}

func (this *ConfigAble) Config()*util.AppConfig{
	return this.conf
}

func (this *ConfigAble) setSheetName(s string ){
	this.sheetName = s
}

func (this *ConfigAble) SheetName()string {
	return this.sheetName
}

type IFrame interface {
	IConfigAble
	Name()string
	OnCreate()
	OnDestroy()
	SetParent(vcl.IWinControl)
	SetEventEmitter(emmiter events.EventEmmiter)
	SetLogger(logger log.ILogger)
	Free()
	SetIndex(num int)
	GetIndex()int
}

type IFrameContainer interface {
	IsFrameSelected(frame IFrame)bool
	On(evt events.EventName,listener...events.Listener)
	Emit(evt events.EventName,args...interface{})
}

type FrameContainer struct {
	log log.ILogger
	listener events.EventEmmiter
	iframes map[string]IFrame
	webviewFrames map[string]*TWebViewFrame
	pageControl *vcl.TPageControl
	component vcl.IComponent
	num int
}

func NewFrameContainer(self vcl.IComponent,pageControl *vcl.TPageControl,log log.ILogger,listener events.EventEmmiter)*FrameContainer{
	ret:=&FrameContainer{
		log:log,
		listener: listener,
		iframes:      make(map[string]IFrame),
		webviewFrames:         make(map[string]*TWebViewFrame),
		pageControl:  pageControl,
		component:    self,
		num:0,
	}

	return ret
}

func (this *FrameContainer) On(evt events.EventName,listener...events.Listener){
	this.listener.On(evt,listener...)
}

func (this *FrameContainer) Emit(evt events.EventName,args...interface{}){
	this.listener.Emit(evt,args...)
}

func (this *FrameContainer) OnCreate(){
	this.pageControl.SetOnChanging(func(sender vcl.IObject, allowChange *bool) {
		go func() {
			vcl.ThreadSync(func() {
				index:=int(this.pageControl.ActivePageIndex())
				this.listener.Emit("page_change",index)
			})
		}()

	})
}

func (this *FrameContainer) IsFrameSelected(frame IFrame)bool{
	control:=this.pageControl.Controls(this.pageControl.ActivePageIndex())
	if control == nil {
		return false
	}
	sheet:=vcl.AsTabSheet(control)
	if sheet==nil {
		return false
	}
	return sheet.Name() == frame.Name()+"Sheet"
}

func (this *FrameContainer) Destroy(){
	for _,frame:=range this.iframes {
		frame.OnDestroy()
		frame.Free()
	}
	this.iframes = make(map[string]IFrame)
	for _,frame:=range this.webviewFrames{
		frame.OnDestroy()
		frame.Free()
	}
	this.webviewFrames = make(map[string]*TWebViewFrame)
}

func (this *FrameContainer) GetIFrame(s string)IFrame{
	return this.iframes[s]
}

type FrameOption func(frame IFrame)

func WithConfig(conf *util.AppConfig)FrameOption{
	return func(frame IFrame) {
		frame.SetConfig(conf)
	}
}

func WithContent(s string,data interface{})FrameOption{
	return func(frame IFrame) {
		frame.SetContent(s,data)
	}
}

func (this *FrameContainer) AddIFrame(name string,frame IFrame,opts... FrameOption)IFrame{
	for _,opt:=range opts {
		opt(frame)
	}
	frame.SetEventEmitter(this.listener)
	frame.SetLogger(this.log)
	frame.setSheetName(name)
	frame.setContainer(this)
	this.iframes[name] = frame
	sheet:=vcl.NewTabSheet(this.component)
	sheet.SetParent(this.pageControl)
	frame.SetIndex(this.num)
	this.num++
	sheet.SetName(frame.Name()+"Sheet")
	sheet.SetCaption(name)
	sheet.SetAlign(types.AlClient)
	frame.SetParent(sheet)
	frame.OnCreate()
	return frame
}

func (this *FrameContainer) AddWebView(name string,url string) {
	sheet:=vcl.NewTabSheet(this.component)
	sheet.SetParent(this.pageControl)
	sheet.SetName(name+"Sheet")
	sheet.SetCaption(name)
	sheet.SetAlign(types.AlClient)
	frame:=NewWebViewFrame(sheet)
	frame.setContainer(this)
	this.iframes[name] = frame
	frame.SetIndex(this.num)
	this.num++
	frame.SetParent(sheet)
	frame.OnCreate()
	frame.SetUrl(url)
}

