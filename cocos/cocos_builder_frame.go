package cocos

import (
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var _ ui.IFrame = (*CocosBuilderFrame)(nil)

type CocosBuilderFrame struct {
	*vcl.TFrame
	ui.ConfigAble

	enginePathEdit  *ui.OpenPathBar
	projectPathEdit *ui.OpenPathBar
	outputPathEdit  *ui.OpenPathBar
	md5Check *ui.EditLabel
	debugCheck *ui.EditLabel
	generateBtn *vcl.TButton
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
	line4:=ui.CreateLine(types.AlTop,42,panel1)
	line3:=ui.CreateLine(types.AlTop,42,panel1)
	line2:=ui.CreateLine(types.AlTop,42,panel1)
	line1:=ui.CreateLine(types.AlTop,42,panel1)
	this.generateBtn = ui.CreateButton("打包",line4)
	line41:=ui.CreateLine(types.AlLeft,80,line4)
	line42:=ui.CreateLine(types.AlLeft,80,line4)
	this.md5Check = ui.NewEditLabel(line41,"md5",120,ui.EDIT_TYPE_BOOL)
	this.md5Check.SetParent(line41)
	this.md5Check.OnCreate()
	this.md5Check.SetAlign(types.AlLeft)
	this.debugCheck = ui.NewEditLabel(line42,"debug",120,ui.EDIT_TYPE_BOOL)
	this.debugCheck.SetParent(line42)
	this.debugCheck.OnCreate()
	this.debugCheck.SetAlign(types.AlLeft)
	this.enginePathEdit = ui.NewOpenPathBar(line1,"引擎路径",480,ui.WithOpenDirDialog("引擎路径"))
	this.enginePathEdit.OnCreate()
	this.enginePathEdit.SetParent(line1)
	this.outputPathEdit = ui.NewOpenPathBar(line3,"导出路径",480,ui.WithOpenDirDialog("导出路径"))
	this.outputPathEdit.OnCreate()
	this.outputPathEdit.SetParent(line3)
	this.projectPathEdit = ui.NewOpenPathBar(line2,"项目路径",480,ui.WithOpenDirDialog("项目路径"))
	this.projectPathEdit.OnCreate()
	this.projectPathEdit.SetParent(line2)
	if s:=this.Config().GetString("engine_path");s!=""{
		this.enginePathEdit.SetPath(s)
	}
	if s:=this.Config().GetString("output_path");s!=""{
		this.outputPathEdit.SetPath(s)
	}
	if s:=this.Config().GetString("project_path");s!=""{
		this.projectPathEdit.SetPath(s)
	}
	this.md5Check.SetBool(this.Config().GetBool("md5"))
	this.debugCheck.SetBool(this.Config().GetBool("debug"))
	this.md5Check.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Config().Set("md5",this.md5Check.Bool())
	}
	this.debugCheck.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Config().Set("debug",this.debugCheck.Bool())
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
	this.generateBtn.SetOnClick(func(sender vcl.IObject) {
		this.GetLogger().Clear()
		enginePath:=this.Config().GetString("engine_path")
		projectPath:=this.Config().GetString("project_path")
		exportPath:=this.Config().GetString("output_path")
		if projectPath==""||exportPath=="" {
			this.GetLogger().Warnf("路径为空",projectPath,exportPath)
			return
		}
		this.GetLogger().Info("开始打包>>>>")
		go func() {
			defer func() {
				vcl.ThreadSync(func() {
					this.generateBtn.SetEnabled(true)
				})
			}()
			vcl.ThreadSync(func() {
				this.generateBtn.SetEnabled(false)
			})
			err:=BuildCocos(&CocosBuildOption{
				Path: projectPath,
				EnginePath: enginePath,
				BuildPath:      exportPath,
				ExcludedModules: []CocosModules{
					//tools.CCMO_COLLIDER,
					//tools.CCMO_DRANGONBONES,
					//tools.CCMO_GEOMUTILS,
					//tools.CCMO_INTERSECTION,
					//tools.CCMO_LABELEFFECT,
					//tools.CCMO_SPINESKELETON,
					//tools.CCMO_STUDIOCOMPONENT,
					//tools.CCMO_TILEMAP,
					//tools.CCMO_VIDEOPLAYER,
					//tools.CCMO_WEBVIEW,
					//tools.CCMO_3D,
					//tools.CCMO_VIDEOPLAYER,
				},
				StartScene: "bee9a278-2b8a-412c-9cec-b8759a57d6b8",
				Platform:        CC_WEB_MOBILE,
				Debug:           this.Config().GetBool("debug"),
				PreviewWidth:    960,
				PreviewHeight:   640,
				WebOrientation:  CC_LANDSCAPELEFT,
				Md5Cache:        this.Config().GetBool("md5"),
			},this.GetLogger())
			if err != nil {
				this.GetLogger().Error(err.Error())
			}
		}()
	})
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


