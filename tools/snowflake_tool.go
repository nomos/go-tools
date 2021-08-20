package tools

import (
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ ui.IFrame = (*SnowflakeTool)(nil)

type SnowflakeTool struct {
	*vcl.TFrame
	ui.ConfigAble
}

func NewSnowflakeTool(owner vcl.IWinControl,option... ui.FrameOption) (root *SnowflakeTool)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *SnowflakeTool) setup(){
	this.SetAlign(types.AlClient)
	line4:=ui.CreateLine(types.AlTop,44,this)
	line3:=ui.CreateLine(types.AlTop,44,this)
	line2:=ui.CreateLine(types.AlTop,44,this)
	line1:=ui.CreateLine(types.AlTop,44,this)
	idEdit:=ui.NewEditLabel(line1,"ID",280,ui.EDIT_TYPE_INTERGER)
	idEdit.SetParent(line1)
	idEdit.OnCreate()
	idEdit.SetChangeOnEdit(true)
	tsEdit:=ui.NewEditLabel(line2,"时间",280,ui.EDIT_TYPE_STRING)
	tsEdit.SetParent(line2)
	tsEdit.OnCreate()
	tsEdit.SetChangeOnEdit(true)
	nodeEdit:=ui.NewEditLabel(line3,"节点",280,ui.EDIT_TYPE_INTERGER)
	nodeEdit.SetParent(line3)
	nodeEdit.OnCreate()
	nodeEdit.SetChangeOnEdit(true)
	seedEdit:=ui.NewEditLabel(line4,"序号",280,ui.EDIT_TYPE_INTERGER)
	seedEdit.SetParent(line4)
	seedEdit.OnCreate()
	seedEdit.SetChangeOnEdit(true)
	idEdit.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		v:=idEdit.Int()
		id:= util.ID(v)
		tsEdit.SetString(util.FormatTimeToISOString(id.Time()))
		nodeEdit.SetInt(int(id.Node()))
		seedEdit.SetInt(int(id.Step()))
	}
}

func (this *SnowflakeTool) OnCreate(){
	this.setup()
}

func (this *SnowflakeTool) OnDestroy(){

}

func (this *SnowflakeTool) OnEnter(){

}

func (this *SnowflakeTool) OnExit(){

}

func (this *SnowflakeTool) Clear(){

}


