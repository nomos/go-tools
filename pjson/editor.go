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

	textEdit *ui.MemoFrame
	valueEditor *ui.ValueEditorFrame
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
	this.SetAlign(types.AlClient)
	ui.CreateSplitter(this,types.AlLeft,6)
	//this.textEdit = vcl.NewMemo(this)
	//this.textEdit.SetParent(this)
	//this.textEdit.SetAlign(types.AlLeft)
	this.textEdit=ui.NewMemoFrame(this)
	this.textEdit.OnCreate()
	this.textEdit.SetAlign(types.AlLeft)
	this.textEdit.SetParent(this)
	rightPanel:=vcl.NewPanel(this)
	rightPanel.SetParent(this)
	rightPanel.SetAlign(types.AlClient)
	ui.CreateSplitter(rightPanel,types.AlTop,6)
	this.tree = treeview.New(rightPanel)
	this.tree.OnCreate()
	this.tree.SetAlign(types.AlTop)
	this.tree.SetParent(rightPanel)

	this.valueEditor = ui.NewValueEditorFrame(rightPanel)
	this.valueEditor.OnCreate()
	this.valueEditor.SetParent(rightPanel)
	this.valueEditor.SetAlign(types.AlClient)

}


func (this *JsonEditor) OnCreate(){
	this.setup()
}

func (this *JsonEditor) OnDestroy() {
	this.tree.OnDestroy()
	this.textEdit.OnDestroy()
	this.valueEditor.OnDestroy()
}
