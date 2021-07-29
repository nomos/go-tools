package ui

import (
	"github.com/nomos/go-lokas/util"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type ValueEditorFrame struct {
	*vcl.TFrame
	ConfigAble

	dropDown  *vcl.TComboBox
	valueEdit *MemoFrame
	keyEdit   *MemoFrame

	schema ITreeSchema

	prettyCheck *vcl.TCheckBox

	OnSetSchema func(schema ITreeSchema)
	OnKeyChange func(schema ITreeSchema)
	OnValueChange func(schema ITreeSchema)
	OnKeyKeyDown func(key types.Char, shift types.TShiftState,schema ITreeSchema)
	OnValueKeyDown func(key types.Char, shift types.TShiftState,schema ITreeSchema)
	OnSelect func(schema ITreeSchema,s string)
	OnKeyExit func(schema ITreeSchema)
	OnValueExit func(schema ITreeSchema)
}

func NewValueEditorFrame(owner vcl.IComponent,option... FrameOption) (root *ValueEditorFrame)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *ValueEditorFrame) setup(){
	line1:=CreateLine(types.AlTop,0,6,24,this)

	this.keyEdit = NewMemoFrame(line1)
	this.keyEdit.BorderSpacing().SetLeft(6)
	this.keyEdit.OnCreate()

	this.dropDown = vcl.NewComboBox(line1)
	this.dropDown.SetParent(line1)
	this.dropDown.SetAlign(types.AlLeft)
	this.dropDown.BorderSpacing().SetTop(4)
	this.dropDown.BorderSpacing().SetBottom(4)
	this.dropDown.SetLeft(11)
	this.dropDown.SetStyle(types.CsDropDownList)
	this.dropDown.SetWidth(100)

	this.valueEdit = NewMemoFrame(this)
	this.valueEdit.BorderSpacing().SetAround(6)
	this.valueEdit.OnCreate()
	CreateSeg(3,line1)
	this.prettyCheck = CreateCheckBox("",line1)
	this.prettyCheck.BorderSpacing().SetBottom(4)
}

func (this *ValueEditorFrame) OnCreate() {
	this.setup()
	this.setCallbacks()
}

func (this *ValueEditorFrame) setCallbacks(){
	this.prettyCheck.SetChecked(this.conf.GetBool("value_edit_format"))
	this.prettyCheck.SetOnChange(func(sender vcl.IObject) {
		this.conf.Set("value_edit_format",this.Pretty())
		if !util.IsNil(this.schema) {
			this.SetSchema(this.schema)
		}
	})
	this.keyEdit.Memo.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if this.OnKeyKeyDown==nil {
			return
		}
		this.OnKeyKeyDown(*key,shift,this.schema)
	})
	this.keyEdit.Memo.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if this.OnValueKeyDown==nil {
			return
		}
		this.OnValueKeyDown(*key,shift,this.schema)
	})
	this.keyEdit.Memo.SetOnExit(func(sender vcl.IObject) {
		if this.OnKeyExit==nil {
			return
		}
		this.OnKeyExit(this.schema)
	})
	this.keyEdit.Memo.SetOnChange(func(sender vcl.IObject) {
		if this.OnKeyChange==nil {
			return
		}
		this.OnKeyChange(this.schema)
	})
	this.valueEdit.Memo.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if this.OnValueKeyDown==nil {
			return
		}
		this.OnValueKeyDown(*key,shift,this.schema)
	})
	this.valueEdit.Memo.SetOnExit(func(sender vcl.IObject) {
		if this.OnValueExit==nil {
			return
		}
		this.OnValueExit(this.schema)
	})
	this.valueEdit.Memo.SetOnChange(func(sender vcl.IObject) {
		if this.OnValueChange==nil {
			return
		}
		this.OnValueChange(this.schema)
	})
	this.dropDown.SetOnSelect(func(sender vcl.IObject) {
		if this.OnSelect==nil {
			return
		}
		if this.schema == nil {
			return
		}
		this.OnSelect(this.schema,this.dropDown.Text())
	})
}

func (this *ValueEditorFrame) GetType()string{
	return this.dropDown.Text()
}

func (this *ValueEditorFrame) AddType(s string){
	this.dropDown.Items().Add(s)
}

func (this *ValueEditorFrame) SetSchema(s ITreeSchema){
	this.schema = s
	this.OnSetSchema(s)
}

func (this *ValueEditorFrame) OnDestroy() {
	this.valueEdit.OnDestroy()
	this.keyEdit.OnDestroy()
}

func (this *ValueEditorFrame) Clear(){
	this.schema = nil
	this.keyEdit.SetText("")
	this.valueEdit.SetText("")
	this.dropDown.SetText("")
}

func (this *ValueEditorFrame) SetKey(s string){
	this.keyEdit.SetText(s)
}

func (this *ValueEditorFrame) SetType(s string){
	vcl.ThreadSync(func() {
		this.dropDown.SetText(s)
	})
}

func (this *ValueEditorFrame) SetValue(s string){
	this.valueEdit.SetText(s)
}

func (this *ValueEditorFrame) Value()string {
	return this.valueEdit.Text()
}

func (this *ValueEditorFrame) Key()string {
	return this.keyEdit.Text()
}

func (this *ValueEditorFrame) Pretty()bool{
	return this.prettyCheck.Checked()
}

func (this *ValueEditorFrame) OnEnter() {

}

func (this *ValueEditorFrame) OnExit() {

}