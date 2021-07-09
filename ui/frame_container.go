package ui

import (
	"github.com/nomos/go-lokas/log"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type FrameContainer struct {
	*vcl.TFrame
	ConfigAble
}

func NewFrameContainer(owner vcl.IComponent,option... FrameOption) (root *FrameContainer) {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *FrameContainer) setup(){
	this.SetAlign(types.AlClient)
}

func (this *FrameContainer) OnCreate(){
	this.setup()
	log.Warnf("OnCreate")
}

func (this *FrameContainer) OnDestroy(){
	log.Warnf("OnDestroy")
}