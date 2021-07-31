package cocos

import (
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

var _ ui.IFrame = (*CocosBuilderFrame)(nil)

type CocosBuilderFrame struct {
	*vcl.TFrame
	ui.ConfigAble

	enginePathEdit  *ui.OpenPathBar
	projectPathEdit *ui.OpenPathBar
	outputPathEdit  *ui.OpenPathBar
}

func NewCocosBuilderFrame(owner vcl.IWinControl,option... ui.FrameOption) (root *CocosBuilderFrame)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *CocosBuilderFrame) setup(){
	this.SetAlign(types.AlClient)
	panel1:=ui.CreatePanel(types.AlClient,this)
	line3:=ui.CreateLine(types.AlTop,42,panel1)
	line2:=ui.CreateLine(types.AlTop,42,panel1)
	line1:=ui.CreateLine(types.AlTop,42,panel1)
	this.enginePathEdit = ui.NewOpenPathBar(line1,"引擎路径",480,ui.WithOpenDirDialog("引擎路径"))
	this.enginePathEdit.OnCreate()
	this.enginePathEdit.SetParent(line1)
	this.enginePathEdit.SetColor(colors.ClYellow)
	this.outputPathEdit = ui.NewOpenPathBar(line3,"导出路径",480,ui.WithOpenDirDialog("导出路径"))
	this.outputPathEdit.OnCreate()
	this.outputPathEdit.SetParent(line3)
	this.outputPathEdit.SetColor(colors.ClYellow)
	this.projectPathEdit = ui.NewOpenPathBar(line2,"项目路径",480,ui.WithOpenDirDialog("项目路径"))
	this.projectPathEdit.OnCreate()
	this.projectPathEdit.SetParent(line2)
	this.projectPathEdit.SetColor(colors.ClYellow)
	if s:=this.Config().GetString("engine_path");s!=""{
		this.enginePathEdit.SetPath(s)
	}
	if s:=this.Config().GetString("output_path");s!=""{
		this.outputPathEdit.SetPath(s)
	}
	if s:=this.Config().GetString("project_path");s!=""{
		this.projectPathEdit.SetPath(s)
	}
	this.enginePathEdit.OnOpen = func(s string) {
		if s== "" {
			return
		}
		this.Config().Set("engine_path",s)
	}
	this.projectPathEdit.OnOpen = func(s string) {
		if s== "" {
			return
		}
		this.Config().Set("project_path",s)
	}
	this.outputPathEdit.OnOpen = func(s string) {
		if s== "" {
			return
		}
		this.Config().Set("output_path",s)
	}
	this.enginePathEdit.OnEdit = func(s string) {
		if s== "" {
			return
		}
		if util.IsFileExist(s) {
			this.Config().Set("engine_path",s)
		}
	}
	this.projectPathEdit.OnEdit = func(s string) {
		if s== "" {
			return
		}
		if util.IsFileExist(s) {
			this.Config().Set("project_path",s)
		}
	}
	this.outputPathEdit.OnEdit = func(s string) {
		if s== "" {
			return
		}
		if util.IsFileExist(s) {
			this.Config().Set("output_path",s)
		}
	}
}

func (this *CocosBuilderFrame) OnCreate(){
	this.setup()
}

func (this *CocosBuilderFrame) OnDestroy(){

}

func (this *CocosBuilderFrame) OnEnter(){

}

func (this *CocosBuilderFrame) OnExit(){

}

func (this *CocosBuilderFrame) Clear(){

}


