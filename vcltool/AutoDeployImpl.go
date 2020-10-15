// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"encoding/json"
	"github.com/nomos/go-log/log"
	"github.com/ying32/govcl/vcl"
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
	frame *TDeployFrame
	sheet *vcl.TTabSheet
}

func NewDeployFile(s string,parent *TAutoDeploy)*DeployFile {
	ret:=&DeployFile{
		Name:       s,
		FilePath:   "",
		Procedures: make([]IDeployProcedure,0),
		parent:parent,
	}
	return ret
}

type DeployFolder struct {
	Files map[string]*DeployFile
}
//::private::
type TAutoDeployFields struct {
	ConfigAble
	console       *TConsoleShell
	folders       []*DeployFolder
	GlobalContext map[string]string
	file          *DeployFile
	selectItem    *vcl.TListItem
}

func (this *TAutoDeploy) OnCreate(){
	this.Panel.Hide()
	this.GlobalContext = this.conf.GetStringMapString("context")
	if this.GlobalContext == nil {
		this.GlobalContext = make(map[string]string)
		this.conf.Set("context",this.GlobalContext)
	}
	this.GlobalContext["aaaa"] = "bbbb"
	this.GlobalContext["cccc"] = "dddd"
	this.console = NewConsoleShell(this)
	this.console.SetParent(this.BottomPanel)
	this.console.OnCreate()
	this.initMenuActions()
	this.initContextActions()
	this.initFileActions()
}

func (this *TAutoDeploy) OnDestroy(){
}

func (this *TAutoDeploy) initMenuActions(){
	this.NewButton.SetOnClick(func(sender vcl.IObject) {
		this.NewFile()
	})
}


func (this *TAutoDeploy) LoadGlobalContext(context map[string]string) {
	log.Warnf("LoadGlobalContext",context)
	for k,v:=range context {
		log.Warnf("list",k,v)
		//this.GlobalContextList.Items().Add()
		//listItem:=this.GlobalContextList.Items().Add()
		//listItem.SubItems().Add(k)
		//listItem.SubItems().Add(v.(string))
	}
}

func (this *TAutoDeploy) initFileActions(){
	this.SaveButton.SetOnClick(func(sender vcl.IObject) {
		log.Warnf("Save")
		this.Save()
	})
	this.CloseButton1.SetOnClick(func(sender vcl.IObject) {
		log.Warnf("Close")
		this.Close()
	})
}

func (this *TAutoDeploy) initContextActions(){
	log.Warnf("initContextActions")
	this.LoadGlobalContext(this.GlobalContext)
	this.ContextAdd.SetOnClick(func(sender vcl.IObject) {
		this.selectItem = nil
		if this.ContextPageControl.ActivePageIndex() == 0 {
			log.Warnf("GlobalContextList.Items().Add()")
			item:=this.GlobalContextList.Items().Add()
			item.SetCaption("")
			item.SubItems().Add("")
			item.MakeVisible(true)
			item.SetSelected(true)
		} else {
			log.Warnf("GlobalContextList.Items().Add()")
			item:=this.FileContextList.Items().Add()
			item.SetCaption("")
			item.SubItems().Add("")
			item.MakeVisible(true)
			item.SetSelected(true)
		}
	})
	this.GlobalContextList.SetOnSelectItem(func(sender vcl.IObject, item *vcl.TListItem, selected bool) {
		if selected {
			log.Warnf("GlobalContext selet",item.Index())
			this.KeyEdit.SetText(item.Caption())
			this.ValueEdit.SetText(item.SubItems().S(0))
			this.selectItem = item
		} else {
			this.selectItem = nil
			log.Warnf("GlobalContext unselect",item.Index())
		}
	})
	this.FileContextList.SetOnSelectItem(func(sender vcl.IObject, item *vcl.TListItem, selected bool) {
		if selected {
			log.Warnf("FileContext selet",item.Index())
			this.KeyEdit.SetText(item.Caption())
			this.ValueEdit.SetText(item.SubItems().S(0))
			this.selectItem = item
		} else {
			this.selectItem = nil
			log.Warnf("FileContext unselect",item.Index())
		}
	})
	this.KeyEdit.SetOnChange(func(sender vcl.IObject) {
		if this.selectItem!=nil {
			this.selectItem.SetCaption(this.KeyEdit.Text())
			this.SaveGlobalContext()
		}
	})
	this.ValueEdit.SetOnChange(func(sender vcl.IObject) {
		if this.selectItem!=nil {
			this.selectItem.SubItems().SetS(0,this.ValueEdit.Text())
			this.SaveGlobalContext()

		}
	})
}

func (this *TAutoDeploy) IsGlobalContext()bool {
	return this.ContextPageControl.ActivePageIndex()==0
}

func (this *TAutoDeploy) SaveGlobalContext(){
	log.Warnf("SaveGlobalContext")
	if this.IsGlobalContext() {
		var i int32
		for i=0;i<this.GlobalContextList.Items().Count();i++{
			item:=this.GlobalContextList.Items().Item(i)
			this.GlobalContext = make(map[string]string)
			this.GlobalContext[item.Caption()] = item.SubItems().S(0)
		}
		log.Warnf("setCOntext",this.GlobalContext)
		jsonstr,_:=json.Marshal(this.GlobalContext)
		this.conf.Set("context",jsonstr)
	}
}

func (this *TAutoDeploy) Save(){

}

func (this *TAutoDeploy) Close(){
	this.Panel.Hide()
}

func (this *TAutoDeploy) SetGlobalContext(context map[string]string){
	this.GlobalContext = context
	this.conf.Set("context",this.GlobalContext)
}

func (this *TAutoDeploy) NewFile(){
	this.Panel.Show()
	this.file=NewDeployFile("未命名",this)
	this.FileName.SetText("未命名")
}

func (this *TAutoDeploy) OnLeftPanelClick(sender vcl.IObject) {

}

func (this *TAutoDeploy) CloseFile(file *DeployFile) {
}

