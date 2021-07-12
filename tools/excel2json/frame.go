package excel2json

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"runtime"
)

type Excel2JsonFrame struct {
	*vcl.TFrame
	ui.ConfigAble

	ExcelEdit *vcl.TLabeledEdit
	TsEdit *vcl.TLabeledEdit
	ExcelButton *vcl.TSpeedButton
	TsButton *vcl.TSpeedButton
	HelpButton *vcl.TSpeedButton
	GenerateButton *vcl.TButton
	IndieFolderCheck *vcl.TCheckBox
}

func NewExcel2JsonFrame(owner vcl.IComponent,option... ui.FrameOption) (root *Excel2JsonFrame)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *Excel2JsonFrame) setup(){
	this.SetOnClick(func(sender vcl.IObject) {
		this.SetFocus()
		this.BringToFront()
	})
	imgList:=icons.GetImageList(32,32)
	line5:=this.createLine(32)
	this.createLine(10)
	line4:=this.createLine(32)
	line3:=this.createLine(24)
	this.createLine(10)
	line2:=this.createLine(32)
	line1:=this.createLine(24)
	this.createText("Excel路径",line1)
	this.createText("Ts路径",line3)
	this.ExcelEdit=this.createEdit(line2)
	this.TsEdit=this.createEdit(line4)
	this.ExcelButton=this.createSpeedBtn("folder",imgList,line2)
	this.TsButton=this.createSpeedBtn("folder",imgList,line4)
	this.HelpButton = this.createSpeedBtn("help",imgList,line5)
	this.createSeg(120,line5)
	this.GenerateButton = this.createButton("生成",line5)
	this.IndieFolderCheck = this.createCheckBox("嵌入",line5)

}

func (this *Excel2JsonFrame) OnCreate(){
	this.setup()
	openDirDialog := vcl.NewSelectDirectoryDialog(this)
	openDirDialog.SetTitle("打开文件夹")
	if this.getExcelPath()!="" {
		this.setExcelPath(this.getExcelPath())
	}
	if this.getDistPath()!="" {
		this.setDistPath(this.getDistPath())
	}
	this.ExcelEdit.SetOnChange(func(sender vcl.IObject) {

	})
	this.TsEdit.SetOnChange(func(sender vcl.IObject) {

	})

	this.ExcelButton.SetOnClick(func(sender vcl.IObject) {
		if this.getExcelPath()!="" {
			openDirDialog.SetInitialDir(this.getExcelPath())
		}
		if openDirDialog.Execute() {
			p := openDirDialog.FileName()
			this.GetLogger().Warn("选择Excel路径"+p)
			this.setExcelPath(p)
		}
	})
	this.TsButton.SetOnClick(func(sender vcl.IObject) {
		if this.getDistPath()!="" {
			openDirDialog.SetInitialDir(this.getDistPath())
		}
		if openDirDialog.Execute() {
			p := openDirDialog.FileName()
			this.GetLogger().Warn("选择导出路径"+p)
			this.setDistPath(p)
		}
	})
	this.IndieFolderCheck.SetOnClick(func(sender vcl.IObject) {
		this.setEmbed(this.IndieFolderCheck.Checked())
	})
	this.GenerateButton.SetOnClick(func(sender vcl.IObject) {
		log.Warn("click")
		if this.getExcelPath()=="" {
			this.ExcelEdit.SetFocus()
			return
		}
		if this.getDistPath()=="" {
			this.TsEdit.SetFocus()
			return
		}
		this.GetLogger().Info("开始生成Json配置...")
		Excel2JsonMiniGame(this.getExcelPath(),this.getDistPath(),this.GetLogger(),this.getEmbed())
	})
	this.HelpButton.SetOnClick(func(sender vcl.IObject) {
		defer func() {
			if r := recover(); r != nil {
				if err,ok:=r.(error);ok {
					log.Error(err.Error())
					buf := make([]byte, 1<<16)
					runtime.Stack(buf, true)
					log.Error(string(buf))
				}
			}
		}()
		log.Warnf(this.GetLogger())
		this.GetLogger().Info("excel的前三行分别为:类型,描述,属性名")
		this.GetLogger().Info("excel表名首字母大写为导出的数据类,默认和其他表名不导出")
		this.GetLogger().Info("属性名必须为英文,且必须包括id字段")
		this.GetLogger().Info("目前支持的类型及类型格式:")
		this.GetLogger().Info("基本类型int,float,string")
		this.GetLogger().Info("数组:[]int,[]float,[]string")
		this.GetLogger().Info("数组格式 [1,2,3,4,5] 或 [1.1,3.4,5.66] 或 [aa,bbb,ccc]")
		this.GetLogger().Info("字符字典类型:[string]string,[string]int,[string]float")
		this.GetLogger().Info("字符字典格式 [aaa:a1,bbb:b2,ccc:c3] 或 [a:1,b:2,c:3] 或 [e:0.5,b:0.6,c:9.3]")
		this.GetLogger().Info("数字字典类型:[int]string,[int]int,[int]float")
		this.GetLogger().Info("数字字典格式 [1:a1,2:b2,3:c3] 或 [1:1,2:2,3:3] 或 [1:0.5,2:0.6,3:9.3]")
		this.GetLogger().Info("导出到cocos勾选嵌入Ts")
	})
}



func (this *Excel2JsonFrame) getEmbed()bool {
	return this.Config().GetBool("embbed")
}

func (this *Excel2JsonFrame) setEmbed(v bool) {
	this.Config().Set("embbed",v)
	this.IndieFolderCheck.SetChecked(v)
}

func (this *Excel2JsonFrame) getExcelPath()string {
	return this.Config().GetString("excel_path")
}

func (this *Excel2JsonFrame) getDistPath()string {
	return this.Config().GetString("dist_path")
}

func (this *Excel2JsonFrame) setExcelPath(p string){
	this.ExcelEdit.SetText(p)
	this.Config().Set("excel_path",p)
}

func (this *Excel2JsonFrame) setDistPath(p string){
	this.TsEdit.SetText(p)
	this.Config().Set("dist_path",p)
}

func (this *Excel2JsonFrame) OnDestroy(){

}



func (this *Excel2JsonFrame) Name()string{
	return "Excel2JsonFrame"
}

func (this *Excel2JsonFrame) createSeg(width int32,parent vcl.IWinControl){
	frame:=vcl.NewPanel(this)
	frame.SetAlign(types.AlLeft)
	frame.SetWidth(width)
	frame.SetParent(parent)
	frame.SetBevelInner(0)
	frame.SetBevelOuter(0)
}

func (this *Excel2JsonFrame) createCheckBox(s string,parent vcl.IWinControl)*vcl.TCheckBox{
	ret:=vcl.NewCheckBox(parent)
	ret.SetParent(parent)
	ret.SetCaption(s)
	ret.BorderSpacing().SetTop(4)
	ret.BorderSpacing().SetBottom(4)
	ret.SetAlign(types.AlLeft)
	return ret
}

func (this *Excel2JsonFrame) createLine(height int32)*vcl.TPanel{
	frame:=vcl.NewPanel(this)
	frame.SetAlign(types.AlTop)
	frame.SetHeight(height)
	frame.BorderSpacing().SetAround(6)
	frame.BorderSpacing().SetLeft(6)
	frame.SetBevelInner(0)
	frame.SetBevelOuter(0)
	frame.SetCaption("")
	frame.SetParent(this)
	return frame
}

func (this *Excel2JsonFrame) createText(c string,parent vcl.IWinControl){
	t:=vcl.NewLabel(this)
	t.SetCaption(c)
	t.SetAlign(types.AlLeft)
	t.SetParent(parent)
}

func (this *Excel2JsonFrame) createButton(s string,parent vcl.IWinControl)*vcl.TButton{
	btn:=vcl.NewButton(this)
	btn.SetAlign(types.AlLeft)
	btn.SetParent(parent)
	btn.SetCaption(s)
	btn.SetWidth(80)
	btn.SetHeight(32)
	return btn
}

func (this *Excel2JsonFrame) createSpeedBtn(s string,img *icons.ImageList,parent vcl.IWinControl)*vcl.TSpeedButton{
	btn:=vcl.NewSpeedButton(this)
	btn.SetAlign(types.AlLeft)
	btn.SetParent(parent)
	btn.SetImages(img.ImageList())
	btn.SetWidth(32)
	btn.SetHeight(32)
	btn.SetImageIndex(img.GetImageIndex(s))
	return btn
}

func (this *Excel2JsonFrame) createEdit(parent vcl.IWinControl)*vcl.TLabeledEdit{
	ret:=vcl.NewLabeledEdit(parent)
	ret.SetParent(parent)
	ret.BorderSpacing().SetLeft(12)
	ret.BorderSpacing().SetTop(2)
	ret.BorderSpacing().SetBottom(2)
	ret.SetWidth(250)
	ret.SetAlign(types.AlLeft)
	ret.SetText("")
	return ret
}