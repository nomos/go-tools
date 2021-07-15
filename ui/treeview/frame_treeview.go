package treeview

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"runtime"
	"sync"
)

type TreeView struct {
	*vcl.TFrame
	Tree *vcl.TTreeView
	ui.ConfigAble
	Root           ui.ITreeSchema
	selectedSchema ui.ITreeSchema
	mu             sync.Mutex
	menus          map[string]*vcl.TMenuItem
	OnUpdateFunc   func(view *TreeView, parent *vcl.TTreeNode, schema ui.ITreeSchema)
	BuildMenuFunc  func(schema ui.ITreeSchema) *vcl.TPopupMenu
	OnEditedFunc   func(schema ui.ITreeSchema, s *string)
	OnEditingFunc  func(schema ui.ITreeSchema)
	OnChangingFunc func()
	OnSelectFunc   func(schema ui.ITreeSchema)
	building       bool
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
	this.Tree.SetImages(icons.GetImageList(16, 16).ImageList())
	this.bindCallbacks()
}

func (this *TreeView) bindCallbacks() {
	this.Tree.SetOnChanging(func(sender vcl.IObject, node *vcl.TTreeNode, allowChange *bool) {
		this.OnChangingFunc()
	})
	this.Tree.SetOnCollapsed(func(sender vcl.IObject, node *vcl.TTreeNode) {
		if this.building {
			return
		}
		schema := this.GetSchemaByNode(node)
		if schema != nil {
			log.Warnf("SetOnCollapsed", node.ToString(), schema.Image(), schema.Collapse(), schema.Key(), schema.Value())
			schema.SetCollapsed()
		}
	})
	this.Tree.SetOnExpanded(func(sender vcl.IObject, node *vcl.TTreeNode) {
		if this.building {
			return
		}
		schema := this.GetSchemaByNode(node)
		if schema != nil {
			log.Warnf("SetOnExpanded", node.ToString(), schema.Image(), schema.Collapse(), schema.Key(), schema.Value())
			schema.SetExpanded()
		}
	})
	this.Tree.SetOnEdited(func(sender vcl.IObject, node *vcl.TTreeNode, s *string) {
		schema := this.GetSchemaByNode(node)
		if schema != nil {
			this.OnEditedFunc(schema, s)
		}
	})
	this.Tree.SetOnEditing(func(sender vcl.IObject, node *vcl.TTreeNode, allowEdit *bool) {
		schema := this.GetSchemaByNode(node)
		if schema != nil {
			this.OnEditingFunc(schema)
		}
	})
	this.Tree.SetOnChange(func(sender vcl.IObject, node *vcl.TTreeNode) {
		schema := this.GetSchemaByNode(node)
		if schema != nil {
			schema.UpdateNode()
		}
	})
	this.Tree.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
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
			log.Warnf("SetOnMouseDown2 32")
			this.selectedSchema = this.GetSchemaByNode(node)
			node.SetSelected(true)
			//TODO:设置编辑器
			p := vcl.Mouse.CursorPos()
			menu := this.BuildMenuFunc(this.selectedSchema)
			menu.Popup(p.X, p.Y)
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
			//node.SetText(this.selectedSchema.String())
			//TODO:设置编辑器
		}
	})
	//完成tree->node抽象->update
	//testcase tree.update node
}

func (this *TreeView) AddNode(node *vcl.TTreeNode, schema ui.ITreeSchema) *vcl.TTreeNode {
	log.Warn(schema.String())
	newNode := this.Tree.Items().AddChild(node, schema.String())
	schema.SetNode(newNode)
	newNode.SetImageIndex(icons.GetImageList(16, 16).GetImageIndex(schema.Image()))
	newNode.SetSelectedIndex(icons.GetImageList(16, 16).GetImageIndex(schema.Image()))
	newNode.SetStateIndex(icons.GetImageList(16, 16).GetImageIndex(schema.Image()))
	log.Warnf("AddNode", newNode.ToString(), schema.Image(), schema.Collapse(), schema.Key(), schema.Value(), schema.String())
	if node == nil {
		schema.SetExpanded()
	}
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
	if node == nil {
		return nil
	}
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
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				log.Error(err.Error())
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, true)
				log.Error(string(buf))
			}
		}
	}()
	var parent *vcl.TTreeNode
	//if schema == nil || schema == this.Root {
	//	if this.Root == nil {
	//		return
	//	}
	//	this.Clear()
	//	schema = this.Root
	//	parent = this.Root.Node()
	//} else if this.Root == nil {
	//	this.Clear()
	//	this.Root = schema
	//	parent = this.Root.Node()
	//} else {
	//	if this.ContainSchema(this.Root, schema) {
	//		parent=schema.Parent().Node()
	//		if schema.Node()!=nil {
	//			schema.Node().Free()
	//		}
	//	} else {
	//		this.Clear()
	//		this.Root = schema
	//		parent = this.Root.Node()
	//	}
	//}
	this.Lock()
	defer this.Unlock()
	this.building = true
	this.ClearTreeNodes()
	if this.Root == nil {
		this.Root = schema
	}
	if schema == nil {
		schema = this.Root
	}
	this.Tree.Items().BeginUpdate()
	this.OnUpdateFunc(this, parent, schema)
	this.Tree.Items().EndUpdate()
	this.building = false
	this.RefreshExpand(this.Root)
}

func (this *TreeView) RefreshExpand(schema ui.ITreeSchema) {
	if schema == nil {
		schema = this.Root
	}
	if !schema.Collapse() {
		log.Warnf("expand2", schema.Image(), schema.Key(), schema.Value())
		if schema.Node() != nil {
			schema.Node().Expand(false)
		}
	}
	for _, s := range schema.Children() {
		this.RefreshExpand(s)
	}
}

func (this *TreeView) GetSelectSchema() ui.ITreeSchema {
	return this.selectedSchema
}

func (this *TreeView) SetSelectSchema(s ui.ITreeSchema) {
	if util.IsNil(s) {
		this.Tree.SetSelected(nil)
		return
	}
	this.selectedSchema = s
	//TODO编辑器
	node := this.GetNodeBySchema(s)
	if node != nil && !node.Selected() {
		node.SetSelected(true)
	}
}

func (this *TreeView) OnCreate() {
	this.setup()
}

func (this *TreeView) OnDestroy() {

}

func (this *TreeView) clearNodes(schema ui.ITreeSchema) {
	if schema == nil {
		schema = this.Root
	}
	if schema == nil {
		return
	}
	schema.SetNode(nil)
	for _, s := range schema.Children() {
		this.clearNodes(s)
	}
}

func (this *TreeView) Clear() {
	this.ClearTreeNodes()
	this.Root = nil
}

func (this *TreeView) ClearTreeNodes() {
	this.Tree.Items().Clear()
	this.clearNodes(this.Root)
}
