package vcltool

import "github.com/ying32/govcl/vcl"

type FrameContainer struct {
	*vcl.TFrame
	ConfigAble
}

func NewFrameContainer(owner vcl.IComponent) (root *FrameContainer) {
	vcl.CreateResFrame(owner, &root)
	return
}

func (this *FrameContainer) setup(){

}

func (this *FrameContainer) OnCreate(){
	this.setup()
}

func (this *FrameContainer) OnDestroy(){

}