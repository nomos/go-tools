package pjson

import (
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/treeview"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ ui.IFrame = (*JsonEditor)(nil)

type JsonEditor struct {
	*vcl.TFrame
	ui.ConfigAble

	textEdit *vcl.TMemo
	tree *treeview.TreeView
}

func NewEditor(owner vcl.IComponent,option... ui.FrameOption) (root *JsonEditor) {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *JsonEditor) setup(){
	this.textEdit = vcl.NewMemo(this)
	this.textEdit.SetParent(this)
	this.textEdit.SetAlign(types.AlLeft)

}


func (this *JsonEditor) OnCreate(){
	this.setup()
}

func (this *JsonEditor) OnDestroy() {

}
