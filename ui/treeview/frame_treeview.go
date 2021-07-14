package treeview

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"sync"
)

type TreeView struct {
	*vcl.TFrame
	Tree *vcl.TTreeView
	ui.ConfigAble
	Root           ui.ITreeSchema
	selectedSchema ui.ITreeSchema
	onUpdate       func(view *TreeView,parent *vcl.TTreeNode, schema ui.ITreeSchema)
	mu             sync.Mutex
	menus          map[string]*vcl.TMenuItem
	BuildMenuFunc  func(schema ui.ITreeSchema)*vcl.TPopupMenu
	OnEditedFunc   func(schema ui.ITreeSchema)
	KeyEditTag     bool
}

func New(owner vcl.IComponent, option ...ui.FrameOption) (root *TreeView) {
	vcl.CreateResFrame(owner, &root)
	for _, o := range option {
		o(root)
	}
	return
}

func (this *TreeView) setup() {
	this.SetAlign(types.AlClient)
	this.Tree = vcl.NewTreeView(this)
	this.Tree.SetParent(this)
	this.Tree.SetAlign(types.AlClient)
	this.Tree.SetImages(icons.GetImageList(16,16).ImageList())
	this.bindCallbacks()
}

func (this *TreeView) bindCallbacks() {
	this.Tree.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		log.Warnf("SetOnMouseDown")
		if util.IsNil(this.Root) {
			log.Error("root is empty")
			return
		}
		log.Warnf("SetOnMouseDown1")
		if button == types.MbRight {
			log.Warnf("SetOnMouseDown2")
			//根据点击位置获取节点
			node := this.Tree.GetNodeAt(x, y)
			if node == nil {
				log.Warnf("node nil")
				return
			}
			log.Warnf("SetOnMouseDown2 32" )
			this.selectedSchema = this.GetSchemaByNode(node)
			node.SetSelected(true)
			//TODO:设置编辑器
			p:=vcl.Mouse.CursorPos()
			menu:=this.BuildMenuFunc(this.selectedSchema)
			menu.Popup(p.X,p.Y)
		} else if button == types.MbLeft {
			log.Warnf("SetOnMouseDown3")
			node := this.Tree.GetNodeAt(x, y)
			if node == nil {
				log.Warnf("SetOnMouseDown3 nil")
				return
			}
			if this.selectedSchema != nil {
				next_node := this.GetNodeBySchema(this.selectedSchema)
				if next_node == node {
					return
				}
			}
			this.selectedSchema = this.GetSchemaByNode(node)
			node.SetText(this.selectedSchema.String())
			//TODO:设置编辑器
		}
	})
	//完成tree->node抽象->update
	//testcase tree.update node
}

func (this *TreeView) AddNode(node *vcl.TTreeNode, schema ui.ITreeSchema) *vcl.TTreeNode {
	newNode := this.Tree.Items().AddChild(node, schema.String())
	schema.SetNode(newNode)
	newNode.SetImageIndex(icons.GetImageList(16,16).GetImageIndex(schema.Image()))
	newNode.SetSelectedIndex(icons.GetImageList(16,16).GetImageIndex(schema.Image()))
	newNode.SetStateIndex(icons.GetImageList(16,16).GetImageIndex(schema.Image()))
	return newNode
}

func (this *TreeView) GetNodeBySchema(s ui.ITreeSchema) *vcl.TTreeNode {
	return s.Node()
	defer this.Unlock()
	this.Lock()
	if this.Root != nil {
		rootTree := s.GetRootTree()
		item := this.Tree.TopItem()
		index := 0
		for {
			if index > len(rootTree)-1 {
				return item
			}
			if int32(rootTree[index]) > item.Count()-1 {
				return nil
			}
			item = item.Item(int32(rootTree[index]))
			index++
		}
	}
	return nil
}

func (this *TreeView) Lock() {
	this.mu.Lock()
}

func (this *TreeView) Unlock() {
	this.mu.Unlock()
}

func (this *TreeView) GetSchemaByNode(node *vcl.TTreeNode) ui.ITreeSchema {
	rootTree := make([]int, 0)
	for {
		if node.Parent() == nil {
			break
		}
		rootTree = append(rootTree, int(node.Index()))
		node = node.Parent()
	}
	schema := this.Root
	for i := len(rootTree) - 1; i >= 0; i-- {
		schema = schema.Children()[rootTree[i]]
	}
	return schema
}

func (this *TreeView) ContainSchema(parent ui.ITreeSchema, schema ui.ITreeSchema) bool {
	if parent == nil {
		parent = this.Root
		return false
	}
	if parent == nil {
		return false
	}
	for _, c := range parent.Children() {
		if c == schema {
			return true
		}
		hasSchema := this.ContainSchema(c, schema)
		if hasSchema {
			return true
		}
	}
	return false
}

func (this *TreeView) UpdateTree(schema ui.ITreeSchema) {
	log.Warnf("UpdateTree")
	var parent *vcl.TTreeNode
	if schema == nil || schema == this.Root {
		if this.Root == nil {
			return
		}
		this.Clear()
		schema = this.Root
		parent = this.Root.Node()
	} else if this.Root == nil {
		this.Clear()
		this.Root = schema
		parent = this.Root.Node()
	} else {
		if this.ContainSchema(this.Root, schema) {
			parent=schema.Parent().Node()
			schema.Node().Free()
		} else {
			this.Clear()
			this.Root = schema
			parent = this.Root.Node()
		}
	}
	this.Tree.Items().BeginUpdate()
	this.onUpdate(this,parent, schema)
	this.Tree.Items().EndUpdate()
}

func (this *TreeView) SetOnUpdate(f func(view *TreeView,parent *vcl.TTreeNode, schema ui.ITreeSchema)) {
	this.onUpdate = f
}

func (this *TreeView) GetSelectSchema()ui.ITreeSchema{
	return this.selectedSchema
}

func (this *TreeView) SetSelectSchema(s ui.ITreeSchema){
	this.selectedSchema = s
	//TODO编辑器
	node:=this.GetNodeBySchema(s)
	if node!=nil&&!node.Selected() {
		node.SetSelected(true)
	}
}

func (this *TreeView) OnCreate() {
	this.setup()
}

func (this *TreeView) OnDestroy() {

}

func (this *TreeView) Clear() {
	this.Tree.Items().Clear()
}
