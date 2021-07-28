package ui

import (
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

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
	this.SetAlign(types.AlClient)
	this.image = vcl.NewImage(this)
	this.image.SetParent(this)
	this.image.SetAlign(types.AlClient)
	this.image.SetOnClick(func(sender vcl.IObject) {
		if this.checked {
			if this.OnChecked!=nil {
				this.OnChecked(true)
			}
			this.check()
		} else {
			if this.OnChecked!=nil {
				this.OnChecked(false)
			}
			this.uncheck()
		}
	})
}

func (this *ToggleFrame) updateCheck(){
	if this.checked {
		vcl.ThreadSync(func() {
			icons.LoadData(this.image,"toggle_on")
		})
	} else {
		vcl.ThreadSync(func() {
			icons.LoadData(this.image,"toggle_off")
		})
	}
}

func (this *ToggleFrame) check(){
	this.checked = true
	this.updateCheck()
}

func (this *ToggleFrame) uncheck(){
	this.checked = false
	this.updateCheck()
}

func (this *ToggleFrame) OnCreate(){
	this.setup()
}

func (this *ToggleFrame) OnDestroy(){

}


