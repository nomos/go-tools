package img_tool

import (
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ ui.IFrame = (*DDSConverter)(nil)

type DDSConverter struct {
	*vcl.TFrame
	ui.ConfigAble
}

func NewDDSConverter(owner vcl.IWinControl,option... ui.FrameOption) (root *DDSConverter)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *DDSConverter) setup(){
	this.SetAlign(types.AlClient)
	this.SetConfig(this.Config().Sub("dds"))
	line3:=ui.CreateLine(types.AlTop,44,this)
	line2:=ui.CreateLine(types.AlTop,44,this)
	line1:=ui.CreateLine(types.AlTop,44,this)
	dirBtn:=ui.NewOpenPathBar(line1,"DDS目录",300)
	dirBtn.OnCreate()
	dirBtn.SetParent(line1)
	transBtn:=ui.CreateButton("转换DDS",line3)
	tag:=ui.NewEditLabel(line2,"Bgrflip",100,ui.EDIT_TYPE_BOOL)
	tag.OnCreate()
	tag.SetParent(line2)
	dirBtn.SetOpenDirDialog("DDS目录")
	dirBtn.OnEdit = func(s string) {
		
	}
	dirBtn.OnOpen = func(s string) {

	}
	transBtn.SetOnClick(func(sender vcl.IObject) {

	})
	tag.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		
	}

}

func (this *DDSConverter) OnCreate(){
	this.setup()
}

func (this *DDSConverter) OnDestroy(){

}

func (this *DDSConverter) OnEnter(){

}

func (this *DDSConverter) OnExit(){

}

func (this *DDSConverter) Clear(){

}



