// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-log/log"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TChartFrameFields struct {
}


func (this *TChartFrame) OnCreate(){
	this.Main.SetOnPaint(func(sender vcl.IObject) {
		//canvas := this.Main.Canvas()
		//canvas.MoveTo(10, 10)
		//canvas.LineTo(50, 10)
		//s := "这是一段文字"
		//canvas.Font().SetColor(colors.ClRed) // red
		//canvas.Font().SetSize(20)
		//style := canvas.Font().Style()
		//canvas.Brush().SetStyle(types.BsClear)
		//canvas.Font().SetStyle(style.Include(types.FsBold, types.FsItalic))
		//canvas.TextOut(100, 30, s)
		//r := types.TRect{0, 0, 80, 80}
		//
		//// 计算文字
		////fmt.Println("TfSingleLine: ", types.TfSingleLine)
		//s = "由于现有第三方的Go UI库不是太庞大就是用的不习惯，或者组件太少。"
		//canvas.TextRect2(&r, s, types.NewSet(types.TfCenter, types.TfVerticalCenter, types.TfSingleLine))
		////fmt.Println("r: ", r, ", s: ", s)
		//
		//s = "测试输出"
		//r = types.TRect{0, 0, 80, 80}
		//// brush
		//canvas.Brush().SetColor(colors.ClGreen)
		//canvas.FillRect(r)
	})
	this.SetOnResize(func(sender vcl.IObject) {
		log.Warnf("resize",this.Width(),this.Height())
	})

	this.Main.SetOnMouseMove(func(sender vcl.IObject, shift types.TShiftState, x, y int32) {
		log.WithFields(log.Fields{
			"x":x,
			"y":y,
		}).Info("SetOnMouseMove")
	})
}


func (this *TChartFrame) OnDestroy() {

}

func (this *TChartFrame) OnMainClick(sender vcl.IObject) {

}

