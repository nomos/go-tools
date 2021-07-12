package ui

import (
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func CreateCheckBox(s string, parent vcl.IWinControl) *vcl.TCheckBox {
	ret := vcl.NewCheckBox(parent)
	ret.SetParent(parent)
	ret.SetCaption(s)
	ret.BorderSpacing().SetTop(4)
	ret.BorderSpacing().SetBottom(4)
	ret.SetAlign(types.AlLeft)
	return ret
}

func CreateLine(height int32,component vcl.IWinControl)*vcl.TPanel{
	frame:=vcl.NewPanel(component)
	frame.SetAlign(types.AlTop)
	frame.SetHeight(height)
	frame.BorderSpacing().SetAround(6)
	frame.BorderSpacing().SetLeft(6)
	frame.SetBevelInner(0)
	frame.SetBevelOuter(0)
	frame.SetCaption("")
	frame.SetParent(component)
	return frame
}



func CreateText(c string,parent vcl.IWinControl)*vcl.TLabel{
	t:=vcl.NewLabel(parent)
	t.SetCaption(c)
	t.SetAlign(types.AlLeft)
	t.SetParent(parent)
	return t
}




func CreateButton(s string,parent vcl.IWinControl)*vcl.TButton{
	btn:=vcl.NewButton(parent)
	btn.SetAlign(types.AlLeft)
	btn.SetParent(parent)
	btn.SetCaption(s)
	btn.SetWidth(80)
	btn.SetHeight(32)
	return btn
}


func CreateEdit(parent vcl.IWinControl)*vcl.TLabeledEdit{
	ret:=vcl.NewLabeledEdit(parent)
	ret.SetParent(parent)
	ret.BorderSpacing().SetLeft(12)
	ret.BorderSpacing().SetTop(2)
	ret.BorderSpacing().SetBottom(2)
	ret.SetWidth(250)
	ret.SetAlign(types.AlLeft)
	ret.SetText("")
	return ret
}

func CreateSeg(width int32,parent vcl.IWinControl){
	frame:=vcl.NewPanel(parent)
	frame.SetAlign(types.AlLeft)
	frame.SetWidth(width)
	frame.SetParent(parent)
	frame.SetBevelInner(0)
	frame.SetBevelOuter(0)
}

func CreateSpeedBtn(s string,img *icons.ImageList,parent vcl.IWinControl)*vcl.TSpeedButton{
	btn:=vcl.NewSpeedButton(parent)
	btn.SetAlign(types.AlLeft)
	btn.SetParent(parent)
	btn.SetImages(img.ImageList())
	btn.SetWidth(32)
	btn.SetHeight(32)
	btn.SetImageIndex(img.GetImageIndex(s))
	return btn
}