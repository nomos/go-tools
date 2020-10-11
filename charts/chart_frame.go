package charts

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type ChartBorder struct {
	Frame       *vcl.TFrame
	Left        int32
	Right       int32
	Top         int32
	Bottom      int32
	TitleHeight int32
	TopSpace    int32
	BottomSpace int32
}

func NewChartBorder(frame *vcl.TFrame) *ChartBorder {
	ret:=&ChartBorder{
		Frame:       frame,
		Left:        50,
		Right:       80,
		Top:         50,
		Bottom:      50,
		TitleHeight: 24,
		TopSpace:    0,
		BottomSpace: 0,
	}
	return ret
}

func (this *ChartBorder) GetChartWidth() int32 {
	return this.Frame.Width()
}

func (this *ChartBorder) GetWidth()int32 {
	return this.Frame.Width() - this.Left - this.Right
}

func (this *ChartBorder) GetChartHeight() int32 {
	return this.Frame.Height()
}

func (this *ChartBorder) GetHeight()int32 {
	return this.Frame.Height() - this.Top - this.Bottom
}

//去掉标题的高度, 上下间距
func (this *ChartBorder) GetHeightEx()int32 {
	return this.Frame.Height() - this.Top - this.Bottom - this.TitleHeight - this.TopSpace - this.BottomSpace
}

func (this *ChartBorder) GetLeft() int32 {
	return this.Left
}

func (this *ChartBorder) GetRight() int32 {
	return this.Frame.Width() - this.Right
}

func (this *ChartBorder) GetTop() int32 {
	return this.Top
}

func (this *ChartBorder) GetTopTitle()int32{
	return this.Top + this.TitleHeight
}

func (this *ChartBorder) GetTopEx() int32 {
	return this.Top + this.TitleHeight + this.TopSpace
}

func (this *ChartBorder) GetBottom() int32 {
	return this.Frame.Height() - this.Bottom
}

func (this *ChartBorder) GetBottomEx()int32 {
	return this.Frame.Height() - this.Bottom - this.BottomSpace
}

func (this *ChartBorder) GetTitleHeight()int32 {
	return this.TitleHeight
}

type IChartFramePainting interface {
	Draw()
	DrawFrame()
	DrawBorder()
	DrawTitleBG()
}

type ChartFramePainting struct {
	Canvas       *vcl.TCanvas
	ChartBorder *ChartBorder
	IsShowBorder bool
	PenBorder    types.TColor
	TitleBGColor types.TColor
}

func (this *ChartFramePainting) Draw() {

}

func (this *ChartFramePainting) DrawFrame() {

}

func (this *ChartFramePainting) DrawBorder() {
	if !this.IsShowBorder {
		return
	}
	this.Canvas.Pen().SetColor(this.PenBorder)
	this.Canvas.FrameRect(types.TRect{
		Left:   this.ChartBorder.GetLeft(),
		Top:    this.ChartBorder.GetTop(),
		Right:  this.ChartBorder.GetRight(),
		Bottom: this.ChartBorder.GetBottom(),
	})
}

func (this *ChartFramePainting) DrawTitleBG() {
	if this.ChartBorder.TitleHeight <= 0 {return}
	this.Canvas.Brush().SetColor(this.TitleBGColor)
	this.Canvas.FillRect(types.TRect{
		Left:   this.ChartBorder.GetLeft(),
		Top:    this.ChartBorder.GetTop(),
		Right:  this.ChartBorder.GetRight(),
		Bottom: this.ChartBorder.GetTopTitle(),
	})
}
