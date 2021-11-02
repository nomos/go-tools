package erpc

import (
	"encoding/json"
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/tools/cocos"
	"io/ioutil"
	"os"
	"path"
	"strings"
)


type sceneJson struct {
	UUID string
}

func loadSceneId(p string)string{

	var scene *sceneJson
	d,_:=ioutil.ReadFile(p+".meta")
	err:=json.Unmarshal(d,&scene)
	if err!=nil {
		log.Error(err.Error())
	}
	return scene.UUID
}

func init(){
	RegisterAdminFunc(LOAD_COCOS_PROJECT, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {

		projectPath:=params.String()
		scenes:=[]string{}
		util.WalkDirFilesWithFunc(path.Join(projectPath,"assets"), func(filePath string, file os.FileInfo) bool {
			if util.IsFileWithExt(file.Name(),".fire") {
				scenes = append(scenes, filePath)
			}
			return false
		},true)
		retStr:=strings.Join(scenes,"|")
		return []byte(retStr),nil
	})
	
	RegisterAdminFunc(BUILD_COCOS_PROJECT, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {

		enginePath:=params.String()
		projectPath:=params.String()
		exportPath:=params.String()
		startScene:=params.String()
		uuid:=loadSceneId(startScene)
		debug:=params.Bool()
		md5:=params.Bool()
		platformEnum:=params.Int32()
		webOrientation :=params.Int32()
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
			StartScene:    uuid,
			Platform:      cocos.CocosPlatform(platformEnum),
			Debug:         debug,
			PreviewWidth:  960,
			PreviewHeight: 640,
			WebOrientation:   cocos.CocosWebOrientation(webOrientation),
			Md5Cache:      md5,
		},logger)
		if err != nil {
			logger.Error(err.Error())
		}
		return []byte(""),nil
	})
}