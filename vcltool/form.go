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
	frames      []IFrame
	created    bool
	relocating bool
}

func NewForm() (root *Form) {
	vcl.Application.CreateForm(&root)
	return
}

func (this *Form) AddFrame(frame IFrame)IFrame {
	this.frames = append(this.frames, frame)
	frame.SetParent(this)
	frame.OnCreate()
	return frame
}

func (this *Form) setup() {
	this.frames = []IFrame{}
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
	this.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		if util.IsNil(this.Config()) || this.relocating {
			return
		}
		x := this.ClientOrigin().X
		y := this.ClientOrigin().Y
		this.Config().Set("posX", x)
		this.Config().Set("posY", y)
	})
	this.created = true
}

func (this *Form) OnFormDestroy(sender vcl.IObject) {
	for _,f:=range this.frames {
		f.OnDestroy()
	}
}
