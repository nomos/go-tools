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
		for {
			if util.IsNil(this.Root) {
				log.Error("root is empty")
				break
			}
			if button == types.MbRight {
				node := this.Tree.GetNodeAt(x, y)
				if node == nil {
					break
				}
			} else if button == types.MbLeft {
				node := this.Tree.GetNodeAt(x, y)
				if this.selectedSchema != nil {
					next_node := this.GetNodeBySchema(this.selectedSchema)
					if node == nil {
						return
					}
					if next_node == node {
						return
					}
				}
			}
			break
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

func (this *TreeView) OnCreate() {
	this.setup()
}

func (this *TreeView) OnDestroy() {

}

func (this *TreeView) Clear() {
	this.Tree.Items().Clear()
}
