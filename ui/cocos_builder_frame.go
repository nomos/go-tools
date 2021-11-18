package ui

import (
	"encoding/json"
	"errors"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/tools/cocos"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var _ IFrame = (*CocosBuilderFrame)(nil)




type CocosBuilderFrame struct {
	*vcl.TFrame
	ConfigAble

	scenes map[string]string
	enginePathEdit     *OpenPathBar
	projectPathEdit    *OpenPathBar
	outputPathEdit     *OpenPathBar
	startSceneEdit     *EditLabel
	webOrientationEdit *EditLabel
	platformEdit       *EditLabel
	md5Check           *EditLabel
	debugCheck         *EditLabel
	generateBtn        *vcl.TButton
}

func NewCocosBuilderFrame(owner vcl.IWinControl,option...FrameOption) (root *CocosBuilderFrame)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *CocosBuilderFrame) setup(){
	this.SetAlign(types.AlClient)
	panel1:= CreatePanel(types.AlClient,this)
	line8:= CreateLine(types.AlTop,32,panel1)
	line4:= CreateLine(types.AlTop,42,panel1)
	line7:= CreateLine(types.AlTop,42,panel1)
	line6:= CreateLine(types.AlTop,42,panel1)
	line5:= CreateLine(types.AlTop,42,panel1)
	line3:= CreateLine(types.AlTop,42,panel1)
	line2:= CreateLine(types.AlTop,42,panel1)
	line1:= CreateLine(types.AlTop,42,panel1)
	this.generateBtn = CreateButton("打包",line8)
	line41:= CreateLine(types.AlLeft,80,line4)
	line42:= CreateLine(types.AlLeft,80,line4)
	this.startSceneEdit = NewEditLabel(line7,"开始场景",280, EDIT_TYPE_ENUM)
	this.startSceneEdit.SetParent(line7)
	this.startSceneEdit.OnCreate()
	this.webOrientationEdit = NewEditLabel(line5,"Web旋转方向",120, EDIT_TYPE_ENUM)
	this.webOrientationEdit.SetEnums(cocos.ALL_CC_WEB_ORIENTATION)
	this.webOrientationEdit.SetParent(line5)
	this.webOrientationEdit.OnCreate()
	this.platformEdit = NewEditLabel(line6,"平台",120, EDIT_TYPE_ENUM)
	this.platformEdit.SetEnums(cocos.ALL_CC_PLATFORM)
	this.platformEdit.SetParent(line6)
	this.platformEdit.OnCreate()
	this.md5Check = NewEditLabel(line41,"md5",120, EDIT_TYPE_BOOL)
	this.md5Check.SetParent(line41)
	this.md5Check.OnCreate()
	this.md5Check.SetAlign(types.AlLeft)
	this.debugCheck = NewEditLabel(line42,"debug",120, EDIT_TYPE_BOOL)
	this.debugCheck.SetParent(line42)
	this.debugCheck.OnCreate()
	this.debugCheck.SetAlign(types.AlLeft)
	this.enginePathEdit = NewOpenPathBar(line1,"引擎路径",480, WithOpenDirDialog("引擎路径"))
	this.enginePathEdit.OnCreate()
	this.enginePathEdit.SetParent(line1)
	this.outputPathEdit = NewOpenPathBar(line3,"导出路径",480, WithOpenDirDialog("导出路径"))
	this.outputPathEdit.OnCreate()
	this.outputPathEdit.SetParent(line3)
	this.projectPathEdit = NewOpenPathBar(line2,"项目路径",480, WithOpenDirDialog("项目路径"))
	this.projectPathEdit.OnCreate()
	this.projectPathEdit.SetParent(line2)
	if s:=this.Config().GetString("engine_path");s!=""{
		this.enginePathEdit.SetPath(s)
	}
	if s:=this.Config().GetString("output_path");s!=""{
		this.outputPathEdit.SetPath(s)
	}
	if s:=this.Config().GetString("project_path");s!=""{
		if util.IsFileExist(s) {
			this.findScenes()
			this.projectPathEdit.SetPath(s)
		}
	}
	this.md5Check.SetBool(this.Config().GetBool("md5"))
	this.debugCheck.SetBool(this.Config().GetBool("debug"))
	this.md5Check.OnValueChange = func(label *EditLabel, editType EDIT_TYPE, value interface{}) {
		this.Config().Set("md5",this.md5Check.Bool())
	}
	this.debugCheck.OnValueChange = func(label *EditLabel, editType EDIT_TYPE, value interface{}) {
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
			this.findScenes()
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
	if this.Config().GetString("platform")!="" {
		e:=cocos.ALL_CC_PLATFORM.GetEnumByString(this.Config().GetString("platform"))
		if !util.IsNil(e) {
			this.platformEdit.SetEnum(e.Enum())
		}
	}
	if this.Config().GetString("web_orientation")!="" {
		e:=cocos.ALL_CC_WEB_ORIENTATION.GetEnumByString(this.Config().GetString("web_orientation"))
		if !util.IsNil(e) {
			this.webOrientationEdit.SetEnum(e.Enum())
		}
	}
	this.startSceneEdit.OnValueChange = func(label *EditLabel, editType EDIT_TYPE, value interface{}) {
		log.Warnf("OnValueChange",value.(protocol.IEnum).ToString())
		this.Config().Set("start_scene",value.(protocol.IEnum).ToString())
	}
	this.platformEdit.OnValueChange = func(label *EditLabel, editType EDIT_TYPE, value interface{}) {
		this.Config().Set("platform",value.(protocol.IEnum).ToString())
	}
	this.webOrientationEdit.OnValueChange = func(label *EditLabel, editType EDIT_TYPE, value interface{}) {
		this.Config().Set("web_orientation",value.(protocol.IEnum).ToString())
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
		go func() {
			defer func() {
				vcl.ThreadSync(func() {
					this.generateBtn.SetEnabled(true)
				})
			}()
			vcl.ThreadSync(func() {
				this.generateBtn.SetEnabled(false)
			})
			err:=cocos.BuildCocos(&cocos.CocosBuildOption{
				Path: projectPath,
				EnginePath: enginePath,
				BuildPath:      exportPath,
				ExcludedModules: []cocos.CocosModules{
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
				StartScene:    this.getSceneUUID(),
				Platform:      this.platformEdit.Enum().(cocos.CocosPlatform),
				Debug:         this.Config().GetBool("debug"),
				PreviewWidth:  960,
				PreviewHeight: 640,
				WebOrientation:   this.webOrientationEdit.Enum().(cocos.CocosWebOrientation),
				Md5Cache:      this.Config().GetBool("md5"),
			},this.GetLogger())
			if err != nil {
				this.GetLogger().Error(err.Error())
			}
		}()
	})
}

func (this *CocosBuilderFrame) OpenProject()error{
	enginePath:=this.Config().GetString("engine_path")
	projectPath:=this.Config().GetString("project_path")
	if projectPath==""{
		this.GetLogger().Warnf("路径为空",projectPath)
		return errors.New("路径为空")
	}
	err:=cocos.OpenCocosProject(&cocos.CocosBuildOption{
		Path: projectPath,
		EnginePath: enginePath,
	},this.GetLogger())
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return err
}

func (this *CocosBuilderFrame) findScenes(){
	this.startSceneEdit.Clear()
	scenes:=[]string{}
	util.WalkDirFilesWithFunc(path.Join(this.Config().GetString("project_path"),"assets"), func(filePath string, file os.FileInfo) bool {
		if util.IsFileWithExt(file.Name(),".fire") {
			scenes = append(scenes, filePath)
		}
		return false
	},true)
	for _,s:=range scenes {
		this.loadSceneId(s)
	}
	if s:=this.Config().GetString("start_scene");s!=""{
		this.startSceneEdit.SetEnumString(s)
	}

}

type sceneJson struct {
	UUID string
}

func (this *CocosBuilderFrame) getSceneUUID()string{
	return this.scenes[this.startSceneEdit.String()]
}

func (this *CocosBuilderFrame) loadSceneId(p string){
	if this.scenes==nil {
		this.scenes = map[string]string{}
	}
	var scene *sceneJson
	d,_:=ioutil.ReadFile(p+".meta")
	err:=json.Unmarshal(d,&scene)
	if err!=nil {
		log.Error(err.Error())
	}
	p = strings.ReplaceAll(p,this.Config().GetString("project_path"),"")
	this.scenes[p] = scene.UUID
	this.startSceneEdit.AddStringEnum(p)
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


