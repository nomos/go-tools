package ui

import "github.com/ying32/govcl/vcl"

type PathSelectFrame struct {
	*vcl.TFrame
	ConfigAble
}

func NewPathSelectFrame(owner vcl.IComponent,option... FrameOption) (root *PathSelectFrame) {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *PathSelectFrame) setup(){
}

func (this *PathSelectFrame) OnCreate(){
	this.setup()
}

func (this *PathSelectFrame) OnDestroy(){

}

func (this *PathSelectFrame) OnEnter(){

}

func (this *PathSelectFrame) OnExit(){

}

func (this *PathSelectFrame) Clear(){

}


