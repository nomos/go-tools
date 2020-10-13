package vcltool

import (
	"github.com/nomos/go-lokas/util"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type IConfigAble interface {
	SetConfig(config *util.AppConfig)
	Config()*util.AppConfig
}

type ConfigAble struct {
	conf *util.AppConfig
}

func (this *ConfigAble) SetConfig(config *util.AppConfig) {
	this.conf = config
}

func (this *ConfigAble) Config()*util.AppConfig{
	return this.conf
}

type IFrame interface {
	IConfigAble
	Name()string
	SheetName()string
	OnCreate()
	OnDestroy()
	SetParent(vcl.IWinControl)
	Free()
}

type FrameContainer struct {
	iframes map[string]IFrame
	urls map[string]string
	webviewFrame *TWebViewFrame
	pageControl *vcl.TPageControl
	component vcl.IComponent
}

func NewFrameContainer(self vcl.IComponent,pageControl *vcl.TPageControl)*FrameContainer{
	ret:=&FrameContainer{
		iframes:      make(map[string]IFrame),
		urls:         make(map[string]string),
		webviewFrame: NewWebViewFrame(self),
		pageControl:  pageControl,
		component:    self,
	}
	ret.webviewFrame.OnCreate()

	return ret
}

func (this *FrameContainer) OnCreate(){
	this.pageControl.SetOnChanging(func(sender vcl.IObject, allowChange *bool) {
		go func() {
			vcl.ThreadSync(func() {
				sheet:=vcl.AsTabSheet(this.pageControl.Controls(this.pageControl.ActivePageIndex()))
				if sheet==nil {
					return
				}
				if _,ok:=this.urls[sheet.Name()];ok {
					this.webviewFrame.SetParent(nil)
					this.webviewFrame.SetParent(sheet)
					this.webviewFrame.Navigate(sheet.Name())
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
	this.webviewFrame.OnDestroy()
	this.webviewFrame.Free()
}

func (this *FrameContainer) GetIFrame(s string)IFrame{
	return this.iframes[s]
}


func (this *FrameContainer) AddIFrame(frame IFrame,conf... *util.AppConfig)IFrame{
	if len(conf)>0 {
		frame.SetConfig(conf[0])
	}
	this.iframes[frame.SheetName()] = frame
	sheet:=vcl.NewTabSheet(this.component)
	sheet.SetParent(this.pageControl)
	sheet.SetName(frame.Name()+"Sheet")
	sheet.SetCaption(frame.SheetName())
	sheet.SetAlign(types.AlClient)
	frame.SetParent(sheet)
	frame.OnCreate()
	return frame
}

func (this *FrameContainer) AddWebView(sheetName string,url string) {
	this.urls[sheetName+"Sheet"] = url
	sheet:=vcl.NewTabSheet(this.component)
	sheet.SetParent(this.pageControl)
	sheet.SetName(sheetName+"Sheet")
	sheet.SetCaption(sheetName)
	sheet.SetAlign(types.AlClient)
	this.webviewFrame.SetUrl(sheetName+"Sheet",url)
}

