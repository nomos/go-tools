package pjson

import (
	"encoding/json"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/nomos/go-tools/ui/treeview"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

var _ ui.IFrame = (*JsonEditor)(nil)

type JsonEditor struct {
	*vcl.TFrame
	ui.ConfigAble

	lastTreeItem       *vcl.TTreeNode
	textEdit *ui.MemoFrame
	valueEditor *ui.ValueEditorFrame
	tree *treeview.TreeView
	
	schema *Schema
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
	line2:=ui.CreateLine(types.AlTop,0,0,32,rightPanel)
	line2.BorderSpacing().SetBottom(3)
	line3:=ui.CreateLine(types.AlTop,6,0,32,leftPanel)
	line3.BorderSpacing().SetBottom(3)
	ui.CreateSeg(3,line2)
	ui.CreateSpeedBtn("sort_up",icons.GetImageList(32,32),line2)
	ui.CreateSeg(3,line2)
	ui.CreateSpeedBtn("sort_down",icons.GetImageList(32,32),line2)
	ui.CreateSeg(3,line2)
	ui.CreateSpeedBtn("null_icon",icons.GetImageList(32,32),line2)
	ui.CreateSeg(3,line2)
	ui.CreateSpeedBtn("object_box",icons.GetImageList(32,32),line2)
	ui.CreateSeg(3,line2)
	ui.CreateSpeedBtn("array_box",icons.GetImageList(32,32),line2)
	ui.CreateSeg(3,line2)
	ui.CreateSpeedBtn("boolean_icon",icons.GetImageList(32,32),line2)
	ui.CreateSeg(3,line2)
	ui.CreateSpeedBtn("number_icon",icons.GetImageList(32,32),line2)
	ui.CreateSeg(3,line2)
	ui.CreateSpeedBtn("string_icon",icons.GetImageList(32,32),line2)

	ui.CreateSpeedBtn("cancel",icons.GetImageList(32,32),line3)
	ui.CreateSeg(3,line3)
	ui.CreateSpeedBtn("save",icons.GetImageList(32,32),line3)
	ui.CreateSeg(3,line3)
	ui.CreateSpeedBtn("folder",icons.GetImageList(32,32),line3)
	this.tree = treeview.New(rightPanel)
	this.tree.OnCreate()
	this.tree.SetAlign(types.AlClient)
	this.tree.SetParent(rightPanel)
	this.tree.BorderSpacing().SetBottom(6)
	this.tree.BorderSpacing().SetRight(6)
	rightPanel.Constraints().SetMinWidth(400)
	rightPanel.SetWidth(400)
	this.valueEditor = ui.NewValueEditorFrame(leftPanel)
	this.valueEditor.SetHeight(200)
	this.valueEditor.OnCreate()
	this.valueEditor.SetParent(leftPanel)
	this.valueEditor.SetAlign(types.AlBottom)
	ui.CreateSplitter(leftPanel,types.AlBottom,6)
}


func (this *JsonEditor) OnCreate(){
	this.setup()
	this.textEdit.SetOnChange(func(sender vcl.IObject) {
		this.parseString()
	})

	this.tree.SetOnUpdate(this.ParseNode)
}

func (this *JsonEditor) OnDestroy() {
	this.tree.OnDestroy()
	this.textEdit.OnDestroy()
	this.valueEditor.OnDestroy()
}

func (this *JsonEditor) parseString() {
	text := this.textEdit.Text()
	i := orderedmap.New()
	err := json.Unmarshal([]byte(text), i)
	if err != nil {
		this.textEdit.Font().SetColor(colors.ClRed)
		return
	}
	this.textEdit.Font().SetColor(colors.ClSysDefault)
	this.schema = NewSchema()
	this.schema.Unmarshal("", -1, *i)
	this.tree.UpdateTree(this.schema)
}


func (this *JsonEditor) Clear() {
	this.tree.Clear()
	this.lastTreeItem = nil
}


func (this *JsonEditor) ParseNode(view *treeview.TreeView,parent *vcl.TTreeNode, schema ui.ITreeSchema) {
	switch schema.(*Schema).Type {
	case Object, Array:
		node := view.AddNode(parent, schema)
		for _, v := range schema.Children() {
			this.ParseNode(view,node, v)
		}
	case String:
		view.AddNode(parent, schema)
	case Number:
		view.AddNode(parent, schema)
	case Boolean:
		view.AddNode(parent, schema)
	case Null:
		view.AddNode(parent, schema)
	default:

	}
}