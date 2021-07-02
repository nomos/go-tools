package vcltool

import (
    "github.com/nomos/go-lokas/log"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
)

type TWebViewFrame struct {
    *vcl.TFrame
    Main         *vcl.TPanel
    
    ConfigAble
    url string
    webview *vcl.TMiniWebview
    initiated bool
}


func NewWebViewFrame(owner vcl.IComponent) (root *TWebViewFrame)  {
    vcl.CreateResFrame(owner, &root)
    return
}


func (this *TWebViewFrame) SetUrl(url string) {
    this.url = url
}

func (this *TWebViewFrame) Refresh(){
    this.webview.Refresh()
}

func (this *TWebViewFrame) setup(){
    log.Errorf("TWebViewFrame setup")
    this.SetAlign(types.AlClient)
    this.Main = vcl.NewPanel(this)
    this.Main.SetParent(this)
    this.Main.SetAlign(types.AlClient)
}

func (this *TWebViewFrame) OnCreate(){
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

func (this *TWebViewFrame) OnDestroy(){

}
