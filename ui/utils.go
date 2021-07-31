package ui

import (
	"github.com/nomos/go-lokas/log"
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


func CreatePanel(align types.TAlign,component vcl.IWinControl)*vcl.TPanel{
	frame:=vcl.NewPanel(component)
	frame.SetBevelOuter(0)
	frame.SetBevelInner(0)
	frame.SetAlign(align)
	frame.SetCaption("")
	frame.SetParent(component)
	return frame
}
func CreateLine(align types.TAlign,size types.TConstraintSize,component vcl.IWinControl)*vcl.TPanel{
	frame:=vcl.NewPanel(component)
	frame.SetBevelOuter(0)
	frame.SetBevelInner(0)
	frame.SetAlign(align)
	if align==types.AlTop||align==types.AlBottom {
		frame.Constraints().SetMinHeight(size)
		frame.Constraints().SetMaxHeight(size)
	} else if align==types.AlLeft||align==types.AlRight {
		frame.Constraints().SetMinWidth(size)
		frame.Constraints().SetMaxWidth(size)
	} else {
		frame.SetWidth(int32(size))
		frame.SetHeight(int32(size))
	}
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

func CreateSplitter(parent vcl.IWinControl,align types.TAlign, size types.TConstraintSize)*vcl.TSplitter{
	splitter:=vcl.NewSplitter(parent)
	splitter.SetParent(parent)
	splitter.SetAlign(align)
	if align==types.AlLeft||align==types.AlRight {
		splitter.Constraints().SetMaxWidth(size)
		splitter.Constraints().SetMinWidth(size)
	} else if align==types.AlTop||align==types.AlBottom{
		splitter.Constraints().SetMaxHeight(size)
		splitter.Constraints().SetMinHeight(size)
	} else {
		log.Panic("wrong align")
	}
	return splitter
}

func CreateSeparator(align types.TAlign,parent vcl.IWinControl)*vcl.TBevel{
	splitter:=vcl.NewBevel(parent)
	splitter.SetParent(parent)
	splitter.SetAlign(align)
	if align==types.AlLeft||align==types.AlRight {
		splitter.Constraints().SetMaxWidth(2)
		splitter.Constraints().SetMinWidth(2)
	} else if align==types.AlTop||align==types.AlBottom{
		splitter.Constraints().SetMaxHeight(2)
		splitter.Constraints().SetMinHeight(2)
	} else {
		log.Panic("wrong align")
	}
	return splitter
}

func CreateEdit(width types.TConstraintSize,parent vcl.IWinControl)*vcl.TEdit{
	ret:=vcl.NewEdit(parent)
	ret.SetParent(parent)
	ret.BorderSpacing().SetLeft(12)
	ret.BorderSpacing().SetTop(2)
	ret.BorderSpacing().SetBottom(2)
	ret.Constraints().SetMaxWidth(width)
	ret.Constraints().SetMinWidth(width)
	ret.SetAlign(types.AlLeft)
	ret.SetText("")
	return ret
}

func CreateEditNoConstrain(parent vcl.IWinControl)*vcl.TEdit{
	ret:=vcl.NewEdit(parent)
	ret.SetParent(parent)
	ret.BorderSpacing().SetTop(2)
	ret.BorderSpacing().SetBottom(2)
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

func CreateImage(s string,parent vcl.IWinControl)*vcl.TImage{
	ret :=icons.GetImage(parent,32,32,s)
	ret.SetAlign(types.AlLeft)
	ret.SetParent(parent)
	ret.SetWidth(32)
	ret.SetHeight(32)
	return ret
}