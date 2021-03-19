// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TWebViewFrameFields struct {
	ConfigAble
	url string
	webview *vcl.TMiniWebview
	initiated bool
}

func (this *TWebViewFrame) SetUrl(url string) {
	this.url = url
}

func (this *TWebViewFrame) Refresh(){
	this.webview.Refresh()
}

func (this *TWebViewFrame) OnCreate(){
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
