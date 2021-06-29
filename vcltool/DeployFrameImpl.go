// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-log/log"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TDeployFrameFields struct {
	ConfigAble
	selectItem    *vcl.TListItem
	file          *DeployFile
}

func (this *TDeployFrame) OnCreate(){

}

func (this *TDeployFrame) OnDestroy(){

}

func (this *TDeployFrame) Save(){

}

func (this *TDeployFrame) Close(){

}

func (this *TDeployFrame) initFileActions(){
	this.CloseButton1.SetOnClick(func(sender vcl.IObject) {
		log.Warnf("Close")
		this.Close()
	})
}


func (this *TDeployFrame) LoadGlobalContext(context map[string]string) {
	log.Warnf("LoadGlobalContext",context)
	for k,v:=range context {
		listItem:=this.GlobalContextList.Items().Add()
		listItem.SetCaption(k)
		listItem.SubItems().Add(v)
	}
}

func (this *TDeployFrame) initContextActions(){
	log.Warnf("initContextActions")
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
			this.KeyEdit.SetText(item.Caption())
			this.ValueEdit.SetText(item.SubItems().S(0))
			this.selectItem = item
		} else {
			this.selectItem = nil
		}
	})
	this.FileContextList.SetOnSelectItem(func(sender vcl.IObject, item *vcl.TListItem, selected bool) {
		if selected {
			this.KeyEdit.SetText(item.Caption())
			this.ValueEdit.SetText(item.SubItems().S(0))
			this.selectItem = item
		} else {
			this.selectItem = nil
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
	this.ContextPageControl.SetActivePageIndex(0)
}

func (this *TDeployFrame) IsGlobalContext()bool {
	return this.ContextPageControl.ActivePageIndex()==0
}

func (this *TDeployFrame) SaveGlobalContext(){
	log.Warnf("SaveGlobalContext")
	if this.IsGlobalContext() {
		var i int32
		this.file.parent.GlobalContext = make(map[string]string)
		for i=0;i<this.GlobalContextList.Items().Count();i++{
			item:=this.GlobalContextList.Items().Item(i)
			this.file.parent.GlobalContext[item.Caption()] = item.SubItems().S(0)
		}
		log.Warnf("setContext",this.file.parent.GlobalContext)
		this.conf.Set("context",this.file.parent.GlobalContext)
	}
}
