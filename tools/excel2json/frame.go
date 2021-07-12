package excel2json

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
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
	line5:=ui.CreateLine(32,this)
	ui.CreateLine(10,this)
	line4:=ui.CreateLine(32,this)
	line3:=ui.CreateLine(24,this)
	ui.CreateLine(10,this)
	line2:=ui.CreateLine(32,this)
	line1:=ui.CreateLine(24,this)
	ui.CreateText("Excel路径",line1)
	ui.CreateText("Ts路径",line3)
	this.ExcelEdit=ui.CreateEdit(line2)
	this.TsEdit=ui.CreateEdit(line4)
	this.ExcelButton=ui.CreateSpeedBtn("folder",imgList,line2)
	this.TsButton=ui.CreateSpeedBtn("folder",imgList,line4)
	this.HelpButton = ui.CreateSpeedBtn("help",imgList,line5)
	ui.CreateSeg(120,line5)
	this.GenerateButton = ui.CreateButton("生成",line5)
	this.IndieFolderCheck = ui.CreateCheckBox("嵌入",line5)

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