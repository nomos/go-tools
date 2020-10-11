package charts

import "github.com/ying32/govcl/vcl"

type Chart struct {
	Frame *vcl.TFrame
	Canvas *vcl.TCanvas
}

func NewChart(frame *vcl.TFrame)*Chart{
	ret:=&Chart{
		Frame:  frame,
		Canvas: vcl.AsPaintBox(frame.Components(0)).Canvas(),
	}
	return ret
}



