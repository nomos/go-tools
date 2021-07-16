package ui

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"runtime"
	"sync"
)

type TreeView struct {
	*vcl.TFrame
	Tree *vcl.TTreeView
	ConfigAble
	Root           ITreeSchema
	selectedSchema ITreeSchema
	mu             sync.Mutex
	menus          map[string]*vcl.TMenuItem
	OnUpdateFunc   func(view *TreeView, parent *vcl.TTreeNode, schema ITreeSchema)
	BuildMenuFunc  func(schema ITreeSchema) *vcl.TPopupMenu
	OnEditedFunc   func(schema ITreeSchema, s *string)
	OnEditingFunc  func(schema ITreeSchema)
	OnChangingFunc func()
	OnSelectFunc   func(schema ITreeSchema)
	building       bool
}

func NewTreeView(owner vcl.IComponent, option ...FrameOption) (root *TreeView) {
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
			schema.SetCollapsed()
		}
	})
	this.Tree.SetOnExpanded(func(sender vcl.IObject, node *vcl.TTreeNode) {
		if this.building {
			return
		}
		schema := this.GetSchemaByNode(node)
		if schema != nil {
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
		if button == types.MbRight {
			//根据点击位置获取节点
			node := this.Tree.GetNodeAt(x, y)
			if node == nil {
				log.Warnf("node nil")
				this.OnSelectFunc(nil)
				return
			}
			this.selectedSchema = this.GetSchemaByNode(node)
			node.SetSelected(true)
			this.OnSelectFunc(this.selectedSchema)
			p := vcl.Mouse.CursorPos()
			menu := this.BuildMenuFunc(this.selectedSchema)
			menu.Popup(p.X, p.Y)
		} else if button == types.MbLeft {
			node := this.Tree.GetNodeAt(x, y)
			if node == nil {
				this.OnSelectFunc(nil)
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
			this.OnSelectFunc(this.selectedSchema)
		}
	})
	//完成tree->node抽象->update
	//testcase tree.update node
}

func (this *TreeView) AddNode(node *vcl.TTreeNode, schema ITreeSchema) *vcl.TTreeNode {
	newNode := this.Tree.Items().AddChild(node, schema.String())
	schema.SetNode(newNode)
	newNode.SetImageIndex(icons.GetImageList(16, 16).GetImageIndex(schema.Image()))
	newNode.SetSelectedIndex(icons.GetImageList(16, 16).GetImageIndex(schema.Image()))
	newNode.SetStateIndex(icons.GetImageList(16, 16).GetImageIndex(schema.Image()))
	if node == nil {
		schema.SetExpanded()
	}
	return newNode
}

func (this *TreeView) GetNodeBySchema(s ITreeSchema) *vcl.TTreeNode {
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

func (this *TreeView) GetSchemaByNode(node *vcl.TTreeNode) ITreeSchema {
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

func (this *TreeView) ContainSchema(parent ITreeSchema, schema ITreeSchema) bool {
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

func (this *TreeView) UpdateTree(schema ITreeSchema) {
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

func (this *TreeView) RefreshExpand(schema ITreeSchema) {
	if schema == nil {
		schema = this.Root
	}
	if !schema.Collapse() {
		if schema.Node() != nil {
			schema.Node().Expand(false)
		}
	}
	for _, s := range schema.Children() {
		this.RefreshExpand(s)
	}
}

func (this *TreeView) GetSelectSchema() ITreeSchema {
	return this.selectedSchema
}

func (this *TreeView) SetSelectSchema(s ITreeSchema) {
	if util.IsNil(s) {
		this.Tree.SetSelected(nil)
		return
	}
	this.selectedSchema = s
	this.OnSelectFunc(this.selectedSchema)
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

func (this *TreeView) clearNodes(schema ITreeSchema) {
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
	this.Tree.Items().Clear()
	this.Root = nil
}

func (this *TreeView) ClearTreeNodes() {
	this.Tree.Items().Clear()
	this.clearNodes(this.Root)
}
