package pjson

import (
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
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
	leftPanel:=ui.CreatePanel(types.AlLeft,this)
	leftPanel.Constraints().SetMinWidth(300)
	leftPanel.SetWidth(350)
	this.textEdit=ui.NewMemoFrame(leftPanel)
	this.textEdit.OnCreate()
	this.textEdit.SetAlign(types.AlClient)
	this.textEdit.SetParent(leftPanel)
	this.textEdit.BorderSpacing().SetLeft(6)
	this.textEdit.BorderSpacing().SetRight(6)
	rightPanel:=ui.CreatePanel(types.AlClient,this)
	line2:=ui.CreateLine(types.AlTop,48,rightPanel)
	ui.CreateSpeedBtn("folder",icons.GetImageList(32,32),line2)
	this.tree = treeview.New(rightPanel)
	this.tree.OnCreate()
	this.tree.SetAlign(types.AlClient)
	this.tree.SetParent(rightPanel)
	this.tree.BorderSpacing().SetBottom(6)
	this.tree.BorderSpacing().SetRight(6)
	rightPanel.Constraints().SetMinWidth(100)
	this.valueEditor = ui.NewValueEditorFrame(leftPanel)
	this.valueEditor.SetHeight(200)
	this.valueEditor.OnCreate()
	this.valueEditor.SetParent(leftPanel)
	this.valueEditor.SetAlign(types.AlBottom)
	ui.CreateSplitter(leftPanel,types.AlBottom,6)

}


func (this *JsonEditor) OnCreate(){
	this.setup()
}

func (this *JsonEditor) OnDestroy() {
	this.tree.OnDestroy()
	this.textEdit.OnDestroy()
	this.valueEditor.OnDestroy()
}
