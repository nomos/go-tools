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
	name string
	url string
	urls map[string]string
	webviews map[string]*vcl.TMiniWebview
	curWebview *vcl.TMiniWebview
}

func (this *TWebViewFrame) SetUrl(s string,url string) {
	this.name = s
	this.url = url
	this.urls[s] = url
	this.webviews[s] = vcl.NewMiniWebview(this)
	this.webviews[s].SetAlign(types.AlClient)
	this.webviews[s].SetParent(this.Main)
	this.webviews[s].Realign()
	this.webviews[s].Navigate(url)
	this.webviews[s].Hide()
}

func (this *TWebViewFrame) SheetName()string {
	return this.name
}

func (this *TWebViewFrame) Navigate(s string){
	if this.curWebview!=nil {
		this.curWebview.Hide()
	}
	if webview,ok:=this.webviews[s];ok {
		webview.Show()
		this.curWebview = webview
		webview.Realign()
	}
}

func (this *TWebViewFrame) OnCreate(){
	this.webviews = make(map[string]*vcl.TMiniWebview)
	this.urls = make(map[string]string)
}

func (this *TWebViewFrame) OnDestroy(){

}
