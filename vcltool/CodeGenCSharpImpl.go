// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-tools/ui"
	"github.com/nomos/promise"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TCodeGenCSharpFields struct {
	ui.ConfigAble

}

func (this *TCodeGenCSharp) OnCreate(){
	openDirDialog := vcl.NewSelectDirectoryDialog(this)
	openDirDialog.SetOptions(openDirDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	openDirDialog.SetTitle("打开文件夹")
	if this.getModelPath()!="" {
		this.setModelPath(this.getModelPath())
	}
	if this.getCSharpPath()!="" {
		this.setCSharpPath(this.getCSharpPath())
	}
	defer func() {
		if r := recover(); r != nil {
			log.Error(r.(error).Error())
		}
	}()
	this.OpenModelDirButton.SetOnClick(func(sender vcl.IObject) {
		if this.getModelPath()!="" {
			openDirDialog.SetInitialDir(this.getModelPath())
		}
		if openDirDialog.Execute() {
			p := openDirDialog.FileName()
			this.GetLogger().Warn("选择Model路径"+p)
			this.setModelPath(p)
		}
	})
	this.OpenDistDirButton.SetOnClick(func(sender vcl.IObject) {
		if this.getCSharpPath()!="" {
			openDirDialog.SetInitialDir(this.getCSharpPath())
		}
		if openDirDialog.Execute() {
			p := openDirDialog.FileName()
			this.GetLogger().Warn("选择导出路径"+p)
			this.setCSharpPath(p)
		}
	})
	this.GenerateButton.SetOnClick(func(sender vcl.IObject) {
		if this.getModelPath()=="" {
			this.ModelDirLabel.SetFocus()
			return
		}
		if this.getCSharpPath()=="" {
			this.DistDirLabel.SetFocus()
			return
		}
		this.GetLogger().Info("开始生成C#文件...")
		g := protocol.NewGenerator(protocol.GEN_GO)
		_,err:=g.LoadModelFolder(this.getModelPath()).Then(func(data interface{}) interface{} {
			g.LoadCsFolder(this.getCSharpPath())
			log.Warnf("LoadModelFolder")
			return promise.Resolve(nil)
		}).Then(func(data interface{}) interface{} {
			log.Info("start generate")
			err:=g.GenerateModel2Cs()
			if err != nil {
				log.Error(err.Error())
				return promise.Reject(nil)
			}
			return promise.Resolve(nil)
		}).Catch(func(err error) interface{} {
			if err != nil {
				log.Error(err.Error())
				return err
			}
			return nil
		}).Await()
		if err != nil {
			log.Error(err.Error())
			this.GetLogger().Error("生成出错:"+err.Error())
		} else {
			this.GetLogger().Info("生成C#文件成功!")
		}
		//excel2json.Excel2JsonMiniGame(this.getExcelPath(),this.getDistPath(),this.log,this.getEmbed())
	})
	this.HelpButton.SetOnClick(func(sender vcl.IObject) {
		this.GetLogger().Info("施工中.....")
	})
}
func (this *TCodeGenCSharp) getModelPath()string {
	return this.Config().GetString("model_path")
}

func (this *TCodeGenCSharp) getCSharpPath()string {
	return this.Config().GetString("csharp_path")
}

func (this *TCodeGenCSharp) setModelPath(p string){
	this.ModelDirLabel.SetText(p)
	this.Config().Set("model_path",p)
}

func (this *TCodeGenCSharp) setCSharpPath(p string){
	this.DistDirLabel.SetText(p)
	this.Config().Set("csharp_path",p)
}

func (this *TCodeGenCSharp) SetConsole(logger *ui.ConsoleShell) {
	this.SetLogger(logger)
}

func (this *TCodeGenCSharp) OnDestroy(){

}


