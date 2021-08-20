package tools

import (
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"strings"
)

var _ ui.IFrame = (*Upper2LowerTool)(nil)

type Upper2LowerTool struct {
	*vcl.TFrame
	ui.ConfigAble
}

func NewUpper2LowerTool(owner vcl.IWinControl,option... ui.FrameOption) (root *Upper2LowerTool)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *Upper2LowerTool) setup(){
	this.SetAlign(types.AlClient)
	top:=ui.CreateLine(types.AlTop,44,this)
	memo:=vcl.NewMemo(this)
	memo.SetParent(this)
	memo.SetAlign(types.AlClient)
	clearButton:=ui.CreateButton("清空",top)
	upper2lower:=ui.CreateButton("AZ-az",top)
	lower2upper:=ui.CreateButton("az-AZ",top)
	clearButton.SetOnClick(func(sender vcl.IObject) {
		memo.Clear()
	})
	upper2lower.SetOnClick(func(sender vcl.IObject) {
		text := memo.Lines().Text()
		text = strings.ToLower(text)
		memo.Clear()
		memo.Lines().SetText(text)
	})
	lower2upper.SetOnClick(func(sender vcl.IObject) {
		text := memo.Lines().Text()
		text = strings.ToUpper(text)
		memo.Clear()
		memo.Lines().SetText(text)
	})
}

func (this *Upper2LowerTool) OnCreate(){
	this.setup()
}

func (this *Upper2LowerTool) OnDestroy(){

}

func (this *Upper2LowerTool) OnEnter(){

}

func (this *Upper2LowerTool) OnExit(){

}

func (this *Upper2LowerTool) Clear(){

}



