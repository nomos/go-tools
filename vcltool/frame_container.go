package vcltool

import (
	"github.com/nomos/go-lokas/util"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type IConfigAble interface {
	SetConfig(config *util.AppConfig)
	Config()*util.AppConfig
	SetSheetName(s string)
	SheetName()string
}

type ConfigAble struct {
	conf *util.AppConfig
	sheetName string
}

func (this *ConfigAble) SetConfig(config *util.AppConfig) {
	this.conf = config
}

func (this *ConfigAble) Config()*util.AppConfig{
	return this.conf
}

func (this *ConfigAble) SetSheetName(s string ){
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
	Free()
}

type FrameContainer struct {
	iframes map[string]IFrame
	webviewFrames map[string]*TWebViewFrame
	pageControl *vcl.TPageControl
	component vcl.IComponent
}

func NewFrameContainer(self vcl.IComponent,pageControl *vcl.TPageControl)*FrameContainer{
	ret:=&FrameContainer{
		iframes:      make(map[string]IFrame),
		webviewFrames:         make(map[string]*TWebViewFrame),
		pageControl:  pageControl,
		component:    self,
	}

	return ret
}

func (this *FrameContainer) OnCreate(){
	this.pageControl.SetOnChanging(func(sender vcl.IObject, allowChange *bool) {
		go func() {
			vcl.ThreadSync(func() {
				control:=this.pageControl.Controls(this.pageControl.ActivePageIndex())
				if control == nil {
					return
				}
				sheet:=vcl.AsTabSheet(control)
				if sheet==nil {
					return
				}
			})
		}()

	})
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


func (this *FrameContainer) AddIFrame(name string,frame IFrame,conf... *util.AppConfig)IFrame{
	if len(conf)>0 {
		frame.SetConfig(conf[0])
	}
	frame.SetSheetName(name)
	this.iframes[name] = frame
	sheet:=vcl.NewTabSheet(this.component)
	sheet.SetParent(this.pageControl)
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
	frame.SetParent(sheet)
	frame.OnCreate()
	frame.SetUrl(url)
}

