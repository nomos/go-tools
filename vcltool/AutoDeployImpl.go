// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-log/log"
	"github.com/ying32/govcl/vcl"
	"regexp"
	"strconv"
	"strings"
)

type IDeployProcedure struct {

}

type DeployProcedure struct {
	Name string
	Type string
	Value string
}

type DeployFile struct {
	Name string
	FilePath string
	Procedures []IDeployProcedure
	Context map[string]string
	parent *TAutoDeploy
	sheet *vcl.TTabSheet
	frame *TDeployFrame
}

func NewDeployFile(s string,parent *TAutoDeploy)*DeployFile {
	ret:=&DeployFile{
		Name:       s,
		FilePath:   "",
		Procedures: make([]IDeployProcedure,0),
		parent:parent,
	}
	sheet:= vcl.NewTabSheet(parent.PageControl)
	sname:=strings.Replace(s,"未命名","new",-1)
	sheet.SetName(sname+"Sheet")
	sheet.SetCaption(s)
	frame := NewDeployFrame(sheet)
	frame.SetName(sname)
	sheet.SetParent(parent.PageControl)
	frame.LoadGlobalContext(parent.GlobalContext)
	frame.SetParent(sheet)
	ret.frame = frame
	ret.sheet = sheet
	ret.frame.file = ret
	return ret
}

type DeployFolder struct {
	Files map[string]*DeployFile
}
//::private::
type TAutoDeployFields struct {
	ConfigAble
	folders       []*DeployFolder
	GlobalContext map[string]string
	deployFiles []*DeployFile
	file          *DeployFile
}

func (this *TAutoDeploy) OnCreate(){
	this.deployFiles = make([]*DeployFile,0)
	this.GlobalContext = this.conf.GetStringMapString("context")
	if this.GlobalContext == nil {
		this.GlobalContext = make(map[string]string)
		this.conf.Set("context",this.GlobalContext)
	}
	this.initMenuActions()
	this.initFileActions()
	this.PageControl.SetOnChange(func(sender vcl.IObject) {

	})
}

func (this *TAutoDeploy) OnDestroy(){

}

func (this *TAutoDeploy) initMenuActions(){
	this.NewButton.SetOnClick(func(sender vcl.IObject) {
		this.NewFile()
	})
}

func (this *TAutoDeploy) initFileActions(){
	this.SaveButton.SetOnClick(func(sender vcl.IObject) {
		log.Warnf("Save")
		this.Save()
	})
}

func (this *TAutoDeploy) SetGlobalContext(context map[string]string){
	this.GlobalContext = context
	this.conf.Set("context",this.GlobalContext)
}

func (this *TAutoDeploy) NewFile(){
	name:=this.checkFileName("未命名")
	this.file=NewDeployFile(name,this)
	this.deployFiles = append(this.deployFiles, this.file)
}

func (this *TAutoDeploy) checkFileName(s string)string {
	reg1:=regexp.MustCompile(s+`([0-9]?)`)
	maxNum:=0
	duplicate:=false
	for _,file:=range this.deployFiles {
		if file.Name==s {
			duplicate = true
		}
		num,_:=strconv.Atoi(reg1.ReplaceAllString(file.Name,"$1"))
		if num>maxNum {
			maxNum = num
		}
	}
	if duplicate {
		log.Warnf("duplicate",maxNum)
		return s+strconv.Itoa(maxNum+1)
	}
	return s
}

func (this *TAutoDeploy) Save(){

}

func (this *TAutoDeploy) OnLeftPanelClick(sender vcl.IObject) {

}

