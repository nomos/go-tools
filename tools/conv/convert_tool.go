package conv

import (
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ ui.IFrame = (*ConvertTool)(nil)

type ConvertTool struct {
	*vcl.TFrame
	ui.ConfigAble

	typeSelectA *ui.EditLabel
	typeSelectB *ui.EditLabel
	ValueA *ui.EditLabel
	ValueB *ui.EditLabel
}

func NewConvertTool(owner vcl.IWinControl,option... ui.FrameOption) (root *ConvertTool)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *ConvertTool) setup(){
	this.SetAlign(types.AlClient)
	line2:=ui.CreateLine(types.AlTop,44,this)
	line22:=ui.CreateLine(types.AlLeft,120,line2)
	ui.CreateSeg(20,line2)
	line21:=ui.CreateLine(types.AlLeft,120,line2)
	line1:=ui.CreateLine(types.AlTop,44,this)
	line12:=ui.CreateLine(types.AlLeft,120,line1)
	ui.CreateSeg(20,line1)
	line11:=ui.CreateLine(types.AlLeft,120,line1)
	this.typeSelectA = ui.NewEditLabel(line11,"类型",120,ui.EDIT_TYPE_ENUM)
	this.typeSelectA.OnCreate()
	this.typeSelectA.SetEnums(ALL_ENC_TYPES)
	this.typeSelectA.SetAlign(types.AlLeft)
	this.typeSelectA.SetParent(line11)
	this.typeSelectB = ui.NewEditLabel(line12,"类型",120,ui.EDIT_TYPE_ENUM)
	this.typeSelectB.OnCreate()
	this.typeSelectB.SetEnums(ALL_ENC_TYPES)
	this.typeSelectB.SetAlign(types.AlLeft)
	this.typeSelectB.SetParent(line12)

	this.ValueA = ui.NewEditLabel(line21,"输入",120,ui.EDIT_TYPE_STRING)
	this.ValueA.OnCreate()
	this.ValueA.SetAlign(types.AlLeft)
	this.ValueA.SetParent(line21)
	this.ValueB = ui.NewEditLabel(line22,"输出",120,ui.EDIT_TYPE_STRING)
	this.ValueB.OnCreate()
	this.ValueB.SetAlign(types.AlLeft)
	this.ValueB.SetParent(line22)
	this.ValueB.SetEnabled(false)


}

func (this *ConvertTool) OnCreate(){
	this.setup()
}

func (this *ConvertTool) OnDestroy(){

}

func (this *ConvertTool) OnEnter(){

}

func (this *ConvertTool) OnExit(){

}

func (this *ConvertTool) Clear(){

}



