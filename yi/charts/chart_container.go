package charts

import (
	"github.com/nomos/go-events"
	"github.com/ying32/govcl/vcl"
)

type ChartContainer struct {
	events.EventEmmiter
	Frame *vcl.TFrame
	Canvas *vcl.TCanvas
}





