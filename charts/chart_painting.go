package charts

import "github.com/ying32/govcl/vcl"

type IChartPainting interface {
	Draw()
	DrawBorder()
	DrawArea()
}

type ChartPainting struct {
	Canvas *vcl.TCanvas
	ChartBorder *ChartBorder
	ChartFrame IChartFramePainting


}

type ChartKLine struct {

}

type ChartLine struct {

}