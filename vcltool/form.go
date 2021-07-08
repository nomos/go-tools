package vcltool

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type Form struct {
	*vcl.TForm
	ConfigAble
	frame IFrame
	created bool
}

func NewForm() *Form {
	var form *Form
	vcl.Application.CreateForm(&form)
	return form
}

func (this *Form) AddFrame(frame IFrame){
	this.frame = frame
	if this.created {
		this.frame.SetParent(this)
		this.frame.OnCreate()
	}
}

func (this *Form) setup(){

}

func (this *Form) SetConfig(config lokas.IConfig) {
	this.ConfigAble.SetConfig(config)
	pos:=this.conf.GetInt("position")
	log.Warnf("pos",pos)
	if pos == 0 {
	} else {
	}

}

func (this *Form) OnFormCreate(sender vcl.IObject) {
	log.Warnf("OnFormCreate")
	this.setup()
	this.SetAlign(types.AlClient)
	if this.frame != nil {
		this.frame.SetParent(this)
		this.frame.OnCreate()
	}
	this.SetOnActivate(func(sender vcl.IObject) {
		log.Warnf("SetOnActivate")
	})
	this.SetOnShow(func(sender vcl.IObject) {
		log.Warnf("SetOnShow")
	})
	this.created = true
}

func (this *Form) OnFormDestroy(sender vcl.IObject) {
	this.frame.OnDestroy()
}
