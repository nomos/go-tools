// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
    "github.com/nomos/go-log/log"
    "github.com/nomos/go-tools/tools/excel2json_tz"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
)

//::private::
type TExcel2JsonTzFields struct {
    ConfigAble
}

func (this *TExcel2JsonTz) OnCreate(){
    log.Warnf("TExcel2JsonTz OnCreate")
    openDirDialog := vcl.NewSelectDirectoryDialog(this)
    openDirDialog.SetOptions(openDirDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
    openDirDialog.SetTitle("打开文件夹")
    if this.getExcelPath()!="" {
        this.setExcelPath(this.getExcelPath())
    }
    if this.getDistPath()!="" {
        this.setDistPath(this.getDistPath())
    }
    this.SetExportType(this.GetExportType())

    this.OpenExcelDirButton.SetOnClick(func(sender vcl.IObject) {
        if this.getExcelPath()!="" {
            openDirDialog.SetInitialDir(this.getExcelPath())
        }
        if openDirDialog.Execute() {
            p := openDirDialog.FileName()
            this.log.Warn("选择Excel路径"+p)
            this.setExcelPath(p)
        }
    })
    this.OpenDistDirButton.SetOnClick(func(sender vcl.IObject) {
        if this.getDistPath()!="" {
            openDirDialog.SetInitialDir(this.getDistPath())
        }
        if openDirDialog.Execute() {
            p := openDirDialog.FileName()
            this.log.Warn("选择导出路径"+p)
            this.setDistPath(p)
        }
    })
    this.ExportSelect.SetOnChange(func(sender vcl.IObject) {
        this.SetExportType(excel2json_tz.GetExportType(this.ExportSelect.Text()))
    })
    this.GenerateButton.SetOnClick(func(sender vcl.IObject) {
        if this.getExcelPath()=="" {
            this.ExcelDirLabel.SetFocus()
            return
        }
        if this.getDistPath()=="" {
            this.DistDirLabel.SetFocus()
            return
        }
        this.log.Info("开始生成Json配置...")
        //excel2json.Excel2JsonMiniGame(this.getExcelPath(),this.getDistPath(),this.log,this.getEmbed())
    })
    this.HelpButton.SetOnClick(func(sender vcl.IObject) {
        this.log.Info("施工中.....")
    })
}

func (this *TExcel2JsonTz) SetExportType(t excel2json_tz.ExportType){
    this.conf.Set("exportType",t.String())
    this.ExportSelect.SetText(t.String())
}

func (this *TExcel2JsonTz) GetExportType()excel2json_tz.ExportType{
    return excel2json_tz.GetExportType(this.conf.GetString("exportType"))
}

func (this *TExcel2JsonTz) getExcelPath()string {
    return this.conf.GetString("excel_path")
}

func (this *TExcel2JsonTz) getDistPath()string {
    return this.conf.GetString("dist_path")
}

func (this *TExcel2JsonTz) setExcelPath(p string){
    this.ExcelDirLabel.SetText(p)
    this.conf.Set("excel_path",p)
}

func (this *TExcel2JsonTz) setDistPath(p string){
    this.DistDirLabel.SetText(p)
    this.conf.Set("dist_path",p)
}

func (this *TExcel2JsonTz) SetConsole(logger *TConsoleShell) {
    this.log = logger
}

func (this *TExcel2JsonTz) OnDestroy(){

}

