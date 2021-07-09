// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-tools/tools/excel2json"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TExcel2JsonMiniGameFrameFields struct {
	ui.ConfigAble
}

func (this *TExcel2JsonMiniGameFrame) OnCreate(){
	openDirDialog := vcl.NewSelectDirectoryDialog(this)
	openDirDialog.SetOptions(openDirDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	openDirDialog.SetTitle("打开文件夹")
	if this.getExcelPath()!="" {
		this.setExcelPath(this.getExcelPath())
	}
	if this.getDistPath()!="" {
		this.setDistPath(this.getDistPath())
	}
	this.setEmbed(this.getEmbed())

	this.OpenExcelDirButton.SetOnClick(func(sender vcl.IObject) {
		if this.getExcelPath()!="" {
			openDirDialog.SetInitialDir(this.getExcelPath())
		}
		if openDirDialog.Execute() {
			p := openDirDialog.FileName()
			this.GetLogger().Warn("选择Excel路径"+p)
			this.setExcelPath(p)
		}
	})
	this.OpenDistDirButton.SetOnClick(func(sender vcl.IObject) {
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
			this.ExcelDirLabel.SetFocus()
			return
		}
		if this.getDistPath()=="" {
			this.DistDirLabel.SetFocus()
			return
		}
		this.GetLogger().Info("开始生成Json配置...")
		excel2json.Excel2JsonMiniGame(this.getExcelPath(),this.getDistPath(),this.GetLogger(),this.getEmbed())
	})
	this.HelpButton.SetOnClick(func(sender vcl.IObject) {
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

func (this *TExcel2JsonMiniGameFrame) getEmbed()bool {
	return this.Config().GetBool("embbed")
}

func (this *TExcel2JsonMiniGameFrame) setEmbed(v bool) {
	this.Config().Set("embbed",v)
	this.IndieFolderCheck.SetChecked(v)
}


func (this *TExcel2JsonMiniGameFrame) getExcelPath()string {
	return this.Config().GetString("excel_path")
}

func (this *TExcel2JsonMiniGameFrame) getDistPath()string {
	return this.Config().GetString("dist_path")
}

func (this *TExcel2JsonMiniGameFrame) setExcelPath(p string){
	this.ExcelDirLabel.SetText(p)
	this.Config().Set("excel_path",p)
}

func (this *TExcel2JsonMiniGameFrame) setDistPath(p string){
	this.DistDirLabel.SetText(p)
	this.Config().Set("dist_path",p)
}

func (this *TExcel2JsonMiniGameFrame) SetConsole(logger *ui.ConsoleShell) {
	this.SetLogger(logger)
}

func (this *TExcel2JsonMiniGameFrame) OnDestroy(){

}
