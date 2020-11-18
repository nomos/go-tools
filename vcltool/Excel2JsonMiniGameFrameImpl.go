// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-tools/tools/excel2json"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TExcel2JsonMiniGameFrameFields struct {
	ConfigAble
	logger log.ILogger
}

func (this *TExcel2JsonMiniGameFrame) OnCreate(){
	log.Warn("TExcel2JsonMiniGameFrame OnCreate")
	openDirDialog := vcl.NewSelectDirectoryDialog(this)
	openDirDialog.SetOptions(openDirDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	openDirDialog.SetTitle("打开文件夹")
	if this.getExcelPath()!="" {
		this.setExcelPath(this.getExcelPath())
	}
	if this.getDistPath()!="" {
		this.setDistPath(this.getDistPath())
	}
	this.setIndiv(this.getIndiv())

	this.OpenExcelDirButton.SetOnClick(func(sender vcl.IObject) {
		if this.getExcelPath()!="" {
			openDirDialog.SetInitialDir(this.getExcelPath())
		}
		if openDirDialog.Execute() {
			p := openDirDialog.FileName()
			this.logger.Warn("选择Excel路径"+p)
			this.setExcelPath(p)
		}
	})
	this.OpenDistDirButton.SetOnClick(func(sender vcl.IObject) {
		if this.getDistPath()!="" {
			openDirDialog.SetInitialDir(this.getDistPath())
		}
		if openDirDialog.Execute() {
			p := openDirDialog.FileName()
			this.logger.Warn("选择导出路径"+p)
			this.setDistPath(p)
		}
	})
	this.IndieFolderCheck.SetOnClick(func(sender vcl.IObject) {
		this.setIndiv(this.IndieFolderCheck.Checked())
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
		this.logger.Info("开始生成Json配置...")
		excel2json.Excel2JsonMiniGame(this.getExcelPath(),this.getDistPath(),this.logger)
	})
}

func (this *TExcel2JsonMiniGameFrame) getIndiv()bool {
	return this.conf.GetBool("indiv_folder")
}

func (this *TExcel2JsonMiniGameFrame) setIndiv(v bool) {
	this.conf.Set("indiv_folder",v)
	this.IndieFolderCheck.SetChecked(v)
}


func (this *TExcel2JsonMiniGameFrame) getExcelPath()string {
	return this.conf.GetString("excel_path")
}

func (this *TExcel2JsonMiniGameFrame) getDistPath()string {
	return this.conf.GetString("dist_path")
}

func (this *TExcel2JsonMiniGameFrame) setExcelPath(p string){
	this.ExcelDirLabel.SetText(p)
	this.conf.Set("excel_path",p)
}

func (this *TExcel2JsonMiniGameFrame) setDistPath(p string){
	this.DistDirLabel.SetText(p)
	this.conf.Set("dist_path",p)
}

func (this *TExcel2JsonMiniGameFrame) SetConsole(logger *TConsoleShell) {
	this.logger = logger
}

func (this *TExcel2JsonMiniGameFrame) OnDestroy(){

}
