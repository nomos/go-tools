// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"time"
)

//::private::
type TWebViewFrameFields struct {
	ConfigAble
	url string
	webview *vcl.TMiniWebview
}

func (this *TWebViewFrame) SetUrl(url string) {
	this.url = url
	this.webview.Navigate(this.url)
	go func() {
		time.Sleep(time.Millisecond)
		vcl.ThreadSync(func() {
			this.webview.Realign()
		})
	}()
}

func (this *TWebViewFrame) Refresh(){
	this.webview.Refresh()
}

func (this *TWebViewFrame) OnCreate(){
	this.webview = vcl.NewMiniWebview(this)
	this.webview.SetParent(this.Main)
	this.webview.SetAlign(types.AlClient)
}

func (this *TWebViewFrame) OnDestroy(){

}
