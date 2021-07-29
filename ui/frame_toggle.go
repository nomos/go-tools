package ui

import (
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ IFrame = (*ToggleFrame)(nil)

type ToggleFrame struct {
	*vcl.TFrame
	ConfigAble

	image *vcl.TImage
	checked bool

	OnChecked func(bool)
}

func NewToggleFrame(owner vcl.IWinControl,checked bool,option... FrameOption) (root *ToggleFrame)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	root.checked = checked
	return
}

func (this *ToggleFrame) setup(){
	this.SetAlign(types.AlLeft)
	this.SetWidth(100)
	this.image = vcl.NewImage(this)
	this.image.SetParent(this)
	this.image.SetAlign(types.AlClient)
	this.image.SetOnClick(func(sender vcl.IObject) {
		this.checked = !this.checked
		if this.checked {
			this.Check()
			if this.OnChecked!=nil {
				this.OnChecked(true)
			}
		} else {
			this.Uncheck()
			if this.OnChecked!=nil {
				this.OnChecked(false)
			}
		}
	})
	go this.updateCheck()
}

func (this *ToggleFrame) updateCheck(){
	if this.checked {
		go vcl.ThreadSync(func() {
			icons.LoadData(this.image,"toggle_on")
		})
	} else {
		go vcl.ThreadSync(func() {
			icons.LoadData(this.image,"toggle_off")
		})
	}
}

func (this *ToggleFrame) Check(){
	this.checked = true
	this.updateCheck()
}

func (this *ToggleFrame) Uncheck(){
	this.checked = false
	this.updateCheck()
}

func (this *ToggleFrame) OnCreate(){
	this.setup()
}

func (this *ToggleFrame) OnDestroy(){

}


func (this *ToggleFrame) OnEnter() {

}

func (this *ToggleFrame) OnExit() {

}

func (this *ToggleFrame) Clear() {
	this.Uncheck()
}

