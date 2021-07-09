package vcltool

import (
	"github.com/nomos/go-lokas/log"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type FrameContainer struct {
	*vcl.TFrame
	ConfigAble
}

func NewFrameContainer(owner vcl.IComponent) (root *FrameContainer) {
	vcl.CreateResFrame(owner, &root)
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