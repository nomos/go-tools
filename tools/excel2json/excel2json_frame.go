package excel2json

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/go-tools/ui/icons"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ ui.IFrame = (*Excel2JsonFrame)(nil)

type Excel2JsonFrame struct {
	*vcl.TFrame
	ui.ConfigAble

	ExcelEdit *ui.OpenPathBar
	TsEdit *ui.OpenPathBar
	HelpButton *vcl.TSpeedButton
	GenerateButton *vcl.TButton
	IndieFolderCheck *vcl.TCheckBox
}

func (this *Excel2JsonFrame) OnEnter() {
	panic("implement me")
}

func (this *Excel2JsonFrame) OnExit() {
	panic("implement me")
}

func (this *Excel2JsonFrame) Clear() {
	panic("implement me")
}

func NewExcel2JsonFrame(owner vcl.IComponent,option... ui.FrameOption) (root *Excel2JsonFrame)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *Excel2JsonFrame) setup(){
	this.SetAlign(types.AlClient)
	imgList:=icons.GetImageList(32,32)
	line5:=ui.CreateLine(types.AlTop,32,this)
	line5.BorderSpacing().SetLeft(6)
	line5.BorderSpacing().SetAround(6)
	line1:=ui.CreateLine(types.AlTop,42,this)
	line2:=ui.CreateLine(types.AlTop,42,this)

	this.ExcelEdit=ui.NewOpenPathBar(line2,"Excel路径",480,ui.WithOpenDirDialog("Excel路径"))
	this.ExcelEdit.OnCreate()
	this.ExcelEdit.SetParent(line2)
	this.TsEdit=ui.NewOpenPathBar(line1,"Ts路径",480,ui.WithOpenDirDialog("Ts路径"))
	this.TsEdit.OnCreate()
	this.TsEdit.SetParent(line1)
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

	this.ExcelEdit.OnOpen = func(s string) {
		if this.getExcelPath()!="" {
			this.ExcelEdit.SetInitialDir(this.getExcelPath())
		}
		if s!="" {
			this.GetLogger().Warn("选择Excel路径"+s)
			this.setExcelPath(s)
		}
	}
	this.ExcelEdit.OnEdit = func(s string) {

	}
	this.TsEdit.OnOpen = func(s string) {
		if this.getDistPath()!="" {
			this.TsEdit.SetInitialDir(this.getDistPath())
		}
		if s!="" {
			this.GetLogger().Warn("选择导出路径"+s)
			this.setDistPath(s)
		}
	}
	this.TsEdit.OnEdit = func(s string) {

	}

	this.IndieFolderCheck.SetOnClick(func(sender vcl.IObject) {
		this.setEmbed(this.IndieFolderCheck.Checked())
	})

	this.GenerateButton.SetOnClick(func(sender vcl.IObject) {
		log.Warn("click")
		if this.getExcelPath()=="" {

			this.GetLogger().Warn("Excel路径为空")
			this.ExcelEdit.SetFocus()
			return
		}
		if this.getDistPath()=="" {
			this.GetLogger().Warn("导出路径为空")
			this.TsEdit.SetFocus()
			return
		}
		this.GetLogger().Info("开始生成Json配置...")
		Excel2JsonMiniGame(this.getExcelPath(),this.getDistPath(),this.GetLogger(),this.getEmbed())
	})

	this.HelpButton.SetOnClick(func(sender vcl.IObject) {
		defer func() {
			if r := recover(); r != nil {
				util.Recover(r,false)
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
	this.ExcelEdit.SetPath(p)
	this.Config().Set("excel_path",p)
}

func (this *Excel2JsonFrame) setDistPath(p string){
	this.TsEdit.SetPath(p)
	this.Config().Set("dist_path",p)
}

func (this *Excel2JsonFrame) OnDestroy(){

}



func (this *Excel2JsonFrame) Name()string{
	return "Excel2JsonFrame"
}