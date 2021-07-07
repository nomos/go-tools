package vcltool

import (
	"github.com/nomos/go-lokas/log"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TTreeViewFrame struct {
	*vcl.TFrame
	Tree *vcl.TTreeView
	ConfigAble
}

func NewTreeViewFrame(owner vcl.IComponent) (root *TTreeViewFrame) {
	vcl.CreateResFrame(owner, &root)
	return
}

func (this *TTreeViewFrame) setup() {
	log.Errorf("TTreeViewFrame setup")
	this.SetAlign(types.AlClient)
	this.Tree = vcl.NewTreeView(this)
	this.Tree.SetParent(this)
	this.Tree.SetAlign(types.AlClient)
}

func (this *TTreeViewFrame) OnCreate() {
	this.setup()
}

func (this *TTreeViewFrame) OnDestroy() {

}
