// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/ying32/govcl/vcl"
    _ "github.com/ying32/govcl/vcl/types"
)


//::private::
type TQuantTradeFrameFields struct {
	ConfigAble
	chartFrame *TChartFrame
}

func (this *TQuantTradeFrame) OnCreate(){
	this.chartFrame = NewChartFrame(this)
	this.chartFrame.SetParent(this.MainChartPanel)
	this.chartFrame.OnCreate()

}

func (this *TQuantTradeFrame) OnDestroy(){
	this.chartFrame.OnDestroy()
}

func (this *TQuantTradeFrame) OnEdit1Change(sender vcl.IObject) {

}

