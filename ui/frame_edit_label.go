package ui

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"reflect"
	"strconv"
)

var _ IFrame = (*EditLabel)(nil)

type EDIT_TYPE protocol.Enum

const (
	EDIT_TYPE_STRING EDIT_TYPE = iota
	EDIT_TYPE_INTERGER
	EDIT_TYPE_DECIMAL
	EDIT_TYPE_ENUM
	EDIT_TYPE_BOOL
)

type EditLabel struct {
	*vcl.TFrame
	ConfigAble

	editType EDIT_TYPE

	value string
	enumValue protocol.IEnum
	boolValue bool

	dirty bool

	name string
	width int32
	incr float64
	label *vcl.TLabel
	edit *vcl.TEdit
	plusBtn *vcl.TSpeedButton
	minusBtn *vcl.TSpeedButton
	mPanel *vcl.TPanel
	dPanel *vcl.TPanel
	numPanel *vcl.TPanel
	enumPanel *vcl.TComboBox
	boolPanel *vcl.TCheckBox
	enums []protocol.IEnum
	enumMap map[string]protocol.IEnum
	enumIntMap map[protocol.Enum]protocol.IEnum
	enabled bool
	color types.TColor
	OnValueChange func(label *EditLabel,editType EDIT_TYPE,value interface{})
}

func NewEditLabel(owner vcl.IWinControl,name string,width int32,t EDIT_TYPE,option... FrameOption) (root *EditLabel)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	root.enabled = true
	root.editType = t
	root.name = name
	root.width = width
	root.incr = 1.0
	root.enums = []protocol.IEnum{}
	return
}

func (this *EditLabel) MarkDirty(){
	this.dirty = true
}

func (this *EditLabel) Dirty()bool{
	return this.dirty
}

func (this *EditLabel) Clean(){
	this.dirty = false
}

func (this *EditLabel) SetEnabled(v bool){
	this.enabled = v
	this.updateEnabled()
}

func (this *EditLabel) updateEnabled(){
	this.label.SetEnabled(this.enabled)
	if !this.enabled {
		this.label.SetColor(colors.ClGray)
	}
	//if this.boolPanel!=nil {
	//	this.boolPanel.SetEnabled(this.enabled)
	//}
	//if this.edit!=nil {
	//	this.edit.SetEnabled(this.enabled)
	//}
	//if this.mPanel!=nil {
	//	this.mPanel.SetEnabled(this.enabled)
	//}
	//if this.enumPanel!=nil {
	//	this.enumPanel.SetEnabled(this.enabled)
	//}
}

func (this *EditLabel) setup(){
	this.SetWidth(this.width)
	this.Constraints().SetMaxWidth(types.TConstraintSize(this.width))
	this.Constraints().SetMinWidth(types.TConstraintSize(this.width))
	this.SetHeight(32)
	this.dPanel=CreateLine(types.AlClient,0,0,22,this)
	line1:=CreateLine(types.AlTop,0,0,10,this)
	this.label = CreateText(this.name,line1)
	this.label.BorderSpacing().SetLeft(2)
	this.label.Font().SetSize(7)
	this.label.Font().SetColor(colors.ClWhite)
	this.label.SetColor(colors.ClBlack)
	this.createNumPanel()
	this.createEnumPanel()
	this.createBoolPanel()
	this.SetType(this.editType)
	this.updateEnumsUI()
	this.updateEnabled()
}

func (this *EditLabel) createBoolPanel(){
	this.boolPanel = vcl.NewCheckBox(this.dPanel)
	this.boolPanel.SetAlign(types.AlClient)
	this.boolPanel.SetParent(this.dPanel)
	this.boolPanel.SetOnClick(func(sender vcl.IObject) {
		if this.boolPanel.Checked() {
			this.SetBool(true,true)
		} else {
			this.SetBool(false,true)
		}
	})
}

func (this *EditLabel) createNumPanel(){
	this.numPanel = vcl.NewPanel(this.dPanel)
	this.numPanel.SetAlign(types.AlClient)
	this.edit = CreateEdit(types.TConstraintSize(this.width),this.numPanel)
	this.edit.BorderSpacing().SetLeft(0)
	this.edit.BorderSpacing().SetTop(0)
	this.edit.BorderSpacing().SetBottom(0)
	this.mPanel=vcl.NewPanel(this.numPanel)
	this.mPanel.SetAlign(types.AlRight)
	this.mPanel.SetParent(this.numPanel)
	this.mPanel.SetHeight(22)
	this.mPanel.SetWidth(14)
	this.mPanel.Constraints().SetMaxWidth(14)
	this.mPanel.Constraints().SetMinWidth(14)
	this.plusBtn = CreateSpeedBtn("sort_up",icons.GetImageList(14,12),this.mPanel)
	this.plusBtn.SetAlign(types.AlTop)
	this.plusBtn.SetHeight(11)
	this.plusBtn.SetFlat(true)
	this.minusBtn = CreateSpeedBtn("sort_down",icons.GetImageList(14,12),this.mPanel)
	this.minusBtn.SetAlign(types.AlClient)
	this.minusBtn.SetHeight(11)
	this.minusBtn.SetFlat(true)
	this.minusBtn.SetOnClick(func(sender vcl.IObject) {
		if !this.enabled {
			return
		}
		switch this.editType {
		case EDIT_TYPE_DECIMAL:
			data:=this.Float()
			this.SetFloat(data-this.incr,true,true)
		case EDIT_TYPE_INTERGER:
			data:=this.Int()
			this.SetInt(data-int(this.incr),true)
		}
	})
	this.plusBtn.SetOnClick(func(sender vcl.IObject) {
		if !this.enabled {
			return
		}
		switch this.editType {
		case EDIT_TYPE_DECIMAL:
			data:=this.Float()
			this.SetFloat(data+this.incr,true)
		case EDIT_TYPE_INTERGER:
			data:=this.Int()
			this.SetInt(data+int(this.incr),true)
		}
	})
	this.edit.SetOnExit(func(sender vcl.IObject) {
		if !this.enabled {
			if this.edit.Text()!=this.value {
				this.edit.SetText(this.value)
			}
			return
		}
		if this.dirty {
			switch this.editType {
			case EDIT_TYPE_INTERGER:
				this.OnValueChange(this,this.editType,this.Int())
				this.dirty = false
			case EDIT_TYPE_STRING:
				this.OnValueChange(this,this.editType,this.String())
				this.dirty = false
			case EDIT_TYPE_DECIMAL:
				this.OnValueChange(this,this.editType,this.Float())
				this.dirty = false
			}
		}
	})
	this.edit.SetOnChange(func(sender vcl.IObject) {
		if !this.enabled {
			if this.edit.Text()!=this.value {
				this.edit.SetText(this.value)
			}
			return
		}
		text:=this.edit.Text()
		switch this.editType {
		case EDIT_TYPE_STRING:
			this.SetString(text)
		case EDIT_TYPE_DECIMAL:
			if !isDouble(text) {
				text = "0.0"
				this.SetString(text)
			} else {
				this.SetString(text)
			}
		case EDIT_TYPE_INTERGER:
			if !isInt(text) {
				text = "0"
				this.SetString(text)
			} else {
				this.SetString(text)
			}
		default:
			log.Panic("type is not support")
		}
	})
}

func (this *EditLabel) createEnumPanel(){
	this.enumPanel = vcl.NewComboBox(this.dPanel)
	this.enumPanel.SetAlign(types.AlClient)
	this.enumPanel.SetParent(this.dPanel)
	this.enumPanel.SetOnSelect(func(sender vcl.IObject) {
		if !this.enabled {
			this.SetEnum(this.enumValue.Enum())
		}
		text:=this.enumPanel.Text()
		if v,ok := this.enumMap[text];ok {
			this.SetEnum(v.Enum(),true)
		}
	})
}

func (this *EditLabel) SetType(t EDIT_TYPE){
	this.editType = t
	if this.editType==EDIT_TYPE_ENUM {
		this.enumPanel.SetParent(this.dPanel)
		this.numPanel.SetParent(nil)
		this.boolPanel.SetParent(nil)
	} else if this.editType==EDIT_TYPE_BOOL {
		this.boolPanel.SetParent(this.dPanel)
		this.numPanel.SetParent(nil)
		this.enumPanel.SetParent(nil)
	} else {
		this.numPanel.SetParent(this.dPanel)
		this.enumPanel.SetParent(nil)
		this.boolPanel.SetParent(nil)
	}
	if this.editType!=EDIT_TYPE_STRING {
		this.mPanel.SetVisible(true)
	} else{
		this.mPanel.SetVisible(false)
	}
}


func (this *EditLabel) SetEnums(enums []protocol.IEnum){
	if enums!= nil {
		this.enums = enums
		this.enumMap = map[string]protocol.IEnum{}
		this.enumIntMap = map[protocol.Enum]protocol.IEnum{}
		for _,e:=range this.enums {
			this.enumMap[e.ToString()] = e
			this.enumIntMap[e.Enum()] = e
		}
	} else {
		this.enums = []protocol.IEnum{}
		this.enumMap = map[string]protocol.IEnum{}
		this.enumIntMap = map[protocol.Enum]protocol.IEnum{}
	}
	this.updateEnumsUI()
}

func (this *EditLabel) updateEnumsUI(){
	if this.enumPanel!=nil {
		this.enumPanel.Clear()
		for _,e:=range this.enums {
			this.enumPanel.Items().Add(e.ToString())
		}
		if len(this.enums)>0 {
			this.enumPanel.SetItemIndex(0)
		}
	}
}

func (this *EditLabel) SetIncrement(incr float64){
	this.incr = incr
}

func (this *EditLabel) SetColor(c types.TColor){
	this.color = c
	this.label.SetColor(c)
	this.label.Font().SetColor(c.RGB(255-c.R(),255-c.G(),255-c.B()))
}

func isDouble(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isInt(s string) bool {
	_, err := strconv.ParseInt(s,10,64)
	return err == nil
}

func (this *EditLabel) SetEnum(enum protocol.Enum,edited...interface{}){
	if v,ok:=this.enumIntMap[enum];ok {
		this.enumPanel.SetText(v.ToString())
		this.enumValue = v
	}
	if !util.IsNil(this.enumValue)&&this.enumValue.Enum() == enum.Enum() {
		return
	}
	if len(edited)>0&&this.OnValueChange!=nil {
		this.OnValueChange(this,this.editType,enum)
	}
}

func (this *EditLabel) SetOnValueChange(f func(label *EditLabel,editType EDIT_TYPE,value interface{})){
	this.OnValueChange = f
}

func (this *EditLabel) SetString(v string,edited...interface{}){
	if this.editType==EDIT_TYPE_ENUM {
		return
	}
	if this.value == v {
		return
	}
	this.value = v
	if this.edit.Text()!=this.value {
		this.edit.SetText(this.value)
	}
	if this.OnValueChange!=nil {
		this.MarkDirty()
	}
}

func (this *EditLabel) SetBool(v bool,edited...interface{}){
	if this.boolValue == v {
		return
	}
	this.boolValue = v
	this.boolPanel.SetChecked(v)
	if len(edited)>0&&this.OnValueChange!=nil {
		this.OnValueChange(this,this.editType,v)
	}
}

func (this *EditLabel) SetInt(v int,edited...interface{}){
	this.SetString(strconv.Itoa(v),edited...)
}

func (this *EditLabel) SetFloat(v float64,edited...interface{}){
	this.SetString(strconv.FormatFloat(v,'f',2,64),edited...)
}

func (this *EditLabel) Int()int{
	if this.value == "" {
		log.Panic("value is not int")
		return 0
	}
	ret, err := strconv.ParseInt(this.value,10,64)
	if err != nil {
		log.Panic(err.Error())
		return 0
	}
	return int(ret)
}

func (this *EditLabel) Float()float64 {
	if this.value == "" {
		log.Panic("value is not float")
		return 0
	}
	ret, err := strconv.ParseFloat(this.value, 64)
	if err != nil {
		log.Panic(err.Error())
		return 0
	}
	return ret
}

func (this *EditLabel) Bool()bool {
	return this.boolValue
}

func (this *EditLabel) String()string{
	return this.value
}

func (this *EditLabel) Set(v interface{}){
	t:=reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.String:
		if this.editType==EDIT_TYPE_STRING {
			this.SetString(v.(string))
		}
	case reflect.Bool:
		if this.editType==EDIT_TYPE_BOOL {
			this.SetBool(v.(bool))
		}
	case reflect.Int:
		if this.editType==EDIT_TYPE_INTERGER {
			this.SetInt(v.(int))
		}
	case reflect.Float64:
		if this.editType==EDIT_TYPE_DECIMAL {
			this.SetFloat(v.(float64))
		}
	case reflect.Int32:
		if this.editType==EDIT_TYPE_ENUM {
			this.SetEnum(v.(protocol.Enum))
		}
	default:
		log.Panic("unrecognized type:"+t.Kind().String())
	}
}

func (this *EditLabel) OnCreate(){
	this.setup()
}

func (this *EditLabel) OnDestroy(){

}



