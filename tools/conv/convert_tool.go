package conv

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/convert"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

var _ ui.IFrame = (*ConvertTool)(nil)

type ConvertTool struct {
	*vcl.TFrame
	ui.ConfigAble

	typeSelectA *ui.EditLabel
	typeSelectB *ui.EditLabel
	ValueA *ui.EditLabel
	ValueB *ui.EditLabel
}

func NewConvertTool(owner vcl.IWinControl,option... ui.FrameOption) (root *ConvertTool)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *ConvertTool) setup(){
	this.SetAlign(types.AlClient)
	this.SetConfig(this.Config().Sub("convert"))
	line2:=ui.CreateLine(types.AlTop,44,this)
	line22:=ui.CreateLine(types.AlLeft,130,line2)
	ui.CreateSeg(20,line2)
	line21:=ui.CreateLine(types.AlLeft,130,line2)
	line1:=ui.CreateLine(types.AlTop,44,this)
	line12:=ui.CreateLine(types.AlLeft,130,line1)
	ui.CreateSeg(20,line1)
	line11:=ui.CreateLine(types.AlLeft,130,line1)
	this.typeSelectA = ui.NewEditLabel(line11,"类型",130,ui.EDIT_TYPE_ENUM)
	this.typeSelectA.OnCreate()
	this.typeSelectA.SetEnums(convert.ALL_ENC_TYPES)
	this.typeSelectA.SetAlign(types.AlLeft)
	this.typeSelectA.SetParent(line11)
	this.typeSelectB = ui.NewEditLabel(line12,"类型",130,ui.EDIT_TYPE_ENUM)
	this.typeSelectB.OnCreate()
	this.typeSelectB.SetEnums(convert.ALL_ENC_TYPES)
	this.typeSelectB.SetAlign(types.AlLeft)
	this.typeSelectB.SetParent(line12)

	this.ValueA = ui.NewEditLabel(line21,"输入",130,ui.EDIT_TYPE_STRING)
	this.ValueA.OnCreate()
	this.ValueA.SetAlign(types.AlLeft)
	this.ValueA.SetParent(line21)
	this.ValueB = ui.NewEditLabel(line22,"输出",130,ui.EDIT_TYPE_STRING)
	this.ValueB.OnCreate()
	this.ValueB.SetAlign(types.AlLeft)
	this.ValueB.SetParent(line22)
	this.ValueB.SetEnabled(false)
	if t:=this.Config().GetString("input_type");t!=""{
		this.typeSelectA.SetEnum(convert.ALL_ENC_TYPES.GetEnumByString(t).Enum())
	}
	if t:=this.Config().GetString("output_type");t!=""{
		this.typeSelectB.SetEnum(convert.ALL_ENC_TYPES.GetEnumByString(t).Enum())
	}
	this.ValueA.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.convert()
	}
	this.ValueA.SetChangeOnEdit(true)
	this.typeSelectB.SetEnums(getConvertAble(this.typeSelectA.Enum().(convert.TYPE)))
	this.typeSelectA.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Config().Set("input_type",this.typeSelectA.Enum().ToString())
		this.typeSelectB.SetEnums(getConvertAble(this.typeSelectA.Enum().(convert.TYPE)))
		this.convert()
	}
	this.typeSelectB.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Config().Set("output_type",this.typeSelectB.Enum().ToString())
		this.convert()
	}
}

func (this *ConvertTool) convert(){
	f:=getConvertFunc(this.typeSelectA.Enum().(convert.TYPE),this.typeSelectB.Enum().(convert.TYPE))
	src:=this.ValueA.String()
	if src == "" {
		this.ValueB.SetString("")
		return
	}
	if f!=nil {
		res,err:=f(src)
		if err != nil {
			log.Error(err.Error())
			this.ValueB.SetString("error")
			this.ValueB.SetColor(colors.ClRed)
			return
		}
		this.ValueB.SetString(res)
		this.ValueB.SetColor(colors.ClWhite)
	}
}

func (this *ConvertTool) OnCreate(){
	this.setup()
}

func (this *ConvertTool) OnDestroy(){

}

func (this *ConvertTool) OnEnter(){

}

func (this *ConvertTool) OnExit(){

}

func (this *ConvertTool) Clear(){

}



