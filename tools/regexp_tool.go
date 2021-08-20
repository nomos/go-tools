package tools

import (
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	regexp "regexp"
)

var _ ui.IFrame = (*RegExpTool)(nil)

type RegExpTool struct {
	*vcl.TFrame
	ui.ConfigAble

	regexpEdit *ui.EditLabel
	srcEdit *ui.EditLabel
	findEdit *ui.EditLabel

	replaceEdit *ui.EditLabel
	resultEdit *ui.EditLabel

	regSelect *ui.EditLabel
}

func NewRegExpTool(owner vcl.IWinControl,option... ui.FrameOption) (root *RegExpTool)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *RegExpTool) setup(){
	this.SetAlign(types.AlClient)
	line4:=ui.CreateLine(types.AlTop,44,this)
	line3:=ui.CreateLine(types.AlTop,44,this)
	line2:=ui.CreateLine(types.AlTop,44,this)

	line2_r:=ui.CreateLine(types.AlClient,280,line2)
	line2_l:=ui.CreateLine(types.AlLeft,280,line2)
	line3_r:=ui.CreateLine(types.AlClient,280,line3)
	line3_l:=ui.CreateLine(types.AlLeft,280,line3)
	line1:=ui.CreateLine(types.AlTop,44,this)
	this.regexpEdit = ui.NewEditLabel(line1,"regexp",540,ui.EDIT_TYPE_STRING)
	this.regexpEdit.SetParent(line1)
	this.regexpEdit.OnCreate()
	this.regexpEdit.SetChangeOnEdit(true)
	this.srcEdit = ui.NewEditLabel(line2_l,"src",260,ui.EDIT_TYPE_STRING)
	this.srcEdit.SetParent(line2_l)
	this.srcEdit.OnCreate()
	this.srcEdit.SetChangeOnEdit(true)
	this.findEdit = ui.NewEditLabel(line2_r,"findstring",260,ui.EDIT_TYPE_STRING)
	this.findEdit.SetParent(line2_r)
	this.findEdit.OnCreate()
	this.findEdit.SetChangeOnEdit(true)
	this.replaceEdit = ui.NewEditLabel(line3_l,"replace",260,ui.EDIT_TYPE_STRING)
	this.replaceEdit.SetParent(line3_l)
	this.replaceEdit.OnCreate()
	this.replaceEdit.SetChangeOnEdit(true)
	this.resultEdit = ui.NewEditLabel(line3_r,"result",260,ui.EDIT_TYPE_STRING)
	this.resultEdit.SetParent(line3_r)
	this.resultEdit.OnCreate()
	this.resultEdit.SetChangeOnEdit(true)
	this.regSelect = ui.NewEditLabel(line4,"regexp list",260,ui.EDIT_TYPE_ENUM)
	this.regSelect.SetParent(line4)
	this.regSelect.OnCreate()
	this.regSelect.SetEnums(protocol.GetAllLineRegExpEnums())
	this.regSelect.SetChangeOnEdit(true)

	this.regSelect.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		regexpMap := protocol.GetLineRegExps()
		regexpress,ok:= regexpMap[this.regSelect.Enum().ToString()]
		if ok {
			this.regexpEdit.SetString(regexpress.String(),true)
		}
	}
	this.regexpEdit.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Regexp()
	}
	this.replaceEdit.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Regexp()
	}
	this.srcEdit.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Regexp()
	}
}

func  (this *RegExpTool) Regexp(){
	regexp1, err := regexp.Compile(this.regexpEdit.String())
	if err != nil {
		this.replaceEdit.Font().SetColor(colors.ClRed)
		this.findEdit.Font().SetColor(colors.ClRed)
	} else {
		this.replaceEdit.Font().SetColor(colors.ClGreen)
		srcStr := this.srcEdit.String()
		dstStr := regexp1.FindString(srcStr)
		resultStr:=regexp1.ReplaceAllString(srcStr,this.replaceEdit.String())
		this.findEdit.SetString(dstStr)
		if dstStr == srcStr {
			this.findEdit.Font().SetColor(colors.ClGreen)
		} else {
			this.findEdit.Font().SetColor(colors.ClRed)
		}
		this.resultEdit.SetString(resultStr)
		this.resultEdit.Font().SetColor(colors.ClGreen)
	}
}

func (this *RegExpTool) OnCreate(){
	this.setup()
}

func (this *RegExpTool) OnDestroy(){

}

func (this *RegExpTool) OnEnter(){

}

func (this *RegExpTool) OnExit(){

}

func (this *RegExpTool) Clear(){

}


