package vcltool

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/util"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type Form struct {
	*vcl.TForm
	ConfigAble
	frame      IFrame
	created    bool
	relocating bool
}

func NewForm(component vcl.IComponent) (root *Form) {
	vcl.Application.CreateForm(&root)
	return
}

func (this *Form) AddFrame(frame IFrame) {
	this.frame = frame
	if this.created {
		this.frame.SetParent(this)
		this.frame.OnCreate()
	}
}

func (this *Form) setup() {

}

func (this *Form) SetConfig(config lokas.IConfig) {
	this.relocating = true
	this.ConfigAble.SetConfig(config)
	posX:=this.Config().GetInt("posX")
	posY:=this.Config().GetInt("posY")
	vcl.ThreadSync(func() {
		this.SetLeft(int32(posX))
		this.SetTop(int32(posY))
		this.Show()
		this.relocating = false
	})
}

func (this *Form) OnFormCreate(sender vcl.IObject) {
	this.setup()
	this.SetAlign(types.AlClient)
	if this.frame != nil {
		this.frame.SetParent(this)
		this.frame.OnCreate()
	}
	this.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		if util.IsNil(this.Config()) || this.relocating {
			return
		}
		x := this.ClientOrigin().X
		y := this.ClientOrigin().Y
		if util.IsNil(this.Config()) {
			this.Config().Set("posX", x)
			this.Config().Set("posY", y)
		}
	})
	this.created = true
}

func (this *Form) OnFormDestroy(sender vcl.IObject) {
	if this.frame != nil {
		this.frame.OnDestroy()
	}
}
