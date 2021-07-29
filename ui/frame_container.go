package ui

import (
	"github.com/nomos/go-lokas/log"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ IFrame = (*FrameContainer)(nil)

type FrameContainer struct {
	*vcl.TFrame
	ConfigAble

	frames []IFrame
}

func NewFrameContainer(owner vcl.IComponent,option... FrameOption) (root *FrameContainer) {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *FrameContainer) AddFrame(frame IFrame){
	frame.SetParent(this)
	frame.OnCreate()
	this.frames = append(this.frames, frame)
}

func (this *FrameContainer) setup(){
	this.frames = []IFrame{}
	this.SetAlign(types.AlClient)
}

func (this *FrameContainer) OnCreate(){
	this.setup()
	log.Warnf("OnCreate")
}

func (this *FrameContainer) OnDestroy(){
	for _,f:=range this.frames {
		f.OnDestroy()
	}
	log.Warnf("OnDestroy")
}

func (this *FrameContainer) OnEnter(){
	for _,f:=range this.frames {
		f.OnEnter()
	}
}

func (this *FrameContainer) OnExit(){
	for _,f:=range this.frames {
		f.OnExit()
	}
}

func (this *FrameContainer) Clear(){
	for _,f:=range this.frames {
		f.Clear()
	}
}