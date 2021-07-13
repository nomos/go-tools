package ui

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ IFrame = (*ValueEditorFrame)(nil)

type ValueEditorFrame struct {
	*vcl.TFrame
	ConfigAble

	dropDown  *vcl.TComboBox
	valueEdit *MemoFrame
	keyEdit   *MemoFrame

	prettyCheck *vcl.TCheckBox
}

func NewValueEditorFrame(owner vcl.IComponent,option... FrameOption) (root *ValueEditorFrame)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *ValueEditorFrame) setup(){
	line1:=CreateLine(types.AlTop,0,6,24,this)

	this.keyEdit = NewMemoFrame(line1)
	this.keyEdit.BorderSpacing().SetLeft(6)
	this.keyEdit.OnCreate()

	this.dropDown = vcl.NewComboBox(line1)
	this.dropDown.SetParent(line1)
	this.dropDown.SetAlign(types.AlLeft)
	this.dropDown.BorderSpacing().SetTop(4)
	this.dropDown.BorderSpacing().SetBottom(4)
	this.dropDown.SetLeft(11)
	this.dropDown.SetStyle(types.CsDropDownList)
	this.dropDown.SetWidth(100)
	this.dropDown.Items().Add("Json")

	this.valueEdit = NewMemoFrame(this)
	this.valueEdit.BorderSpacing().SetAround(6)
	this.valueEdit.OnCreate()

	this.prettyCheck = CreateCheckBox("",line1)
	this.prettyCheck.BorderSpacing().SetBottom(4)
}

func (this *ValueEditorFrame) OnCreate() {
	this.setup()
}

func (this *ValueEditorFrame) OnDestroy() {
	this.valueEdit.OnDestroy()
	this.keyEdit.OnDestroy()
}
