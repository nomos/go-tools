package ui

import (
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
)

type WebViewFrame struct {
    *vcl.TFrame
    Main         *vcl.TPanel
    
    ConfigAble
    url string
    webview *vcl.TMiniWebview
    initiated bool
}


func NewWebViewFrame(owner vcl.IComponent,option... FrameOption) (root *WebViewFrame)  {
    vcl.CreateResFrame(owner, &root)
    for _,o:=range option {
        o(root)
    }
    return
}


func (this *WebViewFrame) SetUrl(url string) {
    this.url = url
}

func (this *WebViewFrame) Refresh(){
    this.webview.Refresh()
}

func (this *WebViewFrame) setup(){
    this.SetAlign(types.AlClient)
    this.Main = vcl.NewPanel(this)
    this.Main.SetParent(this)
    this.Main.SetAlign(types.AlClient)
}

func (this *WebViewFrame) OnCreate(){
    this.setup()
    this.webview = vcl.NewMiniWebview(this)
    this.webview.SetParent(this.Main)
    this.webview.SetAlign(types.AlClient)
    this.container.On("page_change", func(i ...interface{}) {
        num:=i[0].(int)
        if this.index==num {
            if !this.initiated {
                this.webview.Navigate(this.url)
                this.initiated = true
            }
        }
    })
}

func (this *WebViewFrame) OnDestroy(){

}


func (this *WebViewFrame) OnEnter(){

}

func (this *WebViewFrame) OnExit(){

}

func (this *WebViewFrame) Clear(){

}
