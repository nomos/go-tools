package treeview

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"sync"
)

type Frame struct {
	*vcl.TFrame
	Tree *vcl.TTreeView
	ui.ConfigAble
	Root ui.ITreeNode
	mu sync.Mutex
}

func New(owner vcl.IComponent,option... ui.FrameOption) (root *Frame) {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *Frame) setup() {
	log.Errorf("Frame setup")
	this.SetAlign(types.AlClient)
	this.Tree = vcl.NewTreeView(this)
	this.Tree.SetParent(this)
	this.Tree.SetAlign(types.AlClient)
	this.bindCallbacks()
}

func (this *Frame) bindCallbacks(){
	//完成tree->node抽象->update
	//testcase tree.update node
}

func (this *Frame) Lock(){
	this.mu.Lock()
}

func (this *Frame) Unlock() {
	this.mu.Unlock()
}

func (this *Frame) UpdateTree(node ui.ITreeNode){
	if node!=nil&&this.Root!=node {
		this.Root = node
	}
	this.rebuildTree()
}

//重新生成树
func (this *Frame) rebuildTree(){
	if this.Root == nil {
		this.Tree.Items().Clear()
		return
	}

}

func (this *Frame) OnCreate() {
	this.setup()
}

func (this *Frame) OnDestroy() {

}
