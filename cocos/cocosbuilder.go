package cocos

import (
	"bufio"
	"fmt"
	"github.com/nomos/go-lokas/log"
	"io"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

type CocosModules string

type CocosPlatform string

type CocosOrientation string

const (
	CC_WEB_MOBILE  CocosPlatform = "web-mobile"
	CC_WEB_DESKTOP CocosPlatform = "web-desktop"
	CC_ANDROID     CocosPlatform = "android"
	CC_IOS         CocosPlatform = "ios"
	CC_MAC         CocosPlatform = "mac"
	CC_WIN32       CocosPlatform = "win32"
	CC_WECHATGAME  CocosPlatform = "wechatgame"

	CCMO_LABEL               CocosModules = "Label"
	CCMO_COLLIDER            CocosModules = "Collider"
	CCMO_ACTION              CocosModules = "Action"
	CCMO_ANIMATION           CocosModules = "Animation"
	CCMO_CANVASRENDERER      CocosModules = "Canvas Renderer"
	CCMO_DYNAMICALTAS        CocosModules = "Dynamic Atlas"
	CCMO_DRANGONBONES        CocosModules = "DragonBones"
	CCMO_GEOMUTILS           CocosModules = "Geom Utils"
	CCMO_INTERSECTION        CocosModules = "Intersection"
	CCMO_EDITBOX             CocosModules = "EditBox"
	CCMO_LAYOUT              CocosModules = "Layout"
	CCMO_GRAPHICS            CocosModules = "Graphics"
	CCMO_LABELEFFECT         CocosModules = "Label Effect"
	CCMO_MASK                CocosModules = "Mask"
	CCMO_MESH                CocosModules = "Mesh"
	CCMO_MOTIONSTREAK        CocosModules = "MotionStreak"
	CCMO_NODEPOOL            CocosModules = "NodePool"
	CCMO_PHYSICS             CocosModules = "Physics"
	CCMO_PAGEVIEW            CocosModules = "PageView"
	CCMO_PAGEVIEWINDICATOR   CocosModules = "PageViewIndicator"
	CCMO_PROGRESSBAR         CocosModules = "ProgressBar"
	CCMO_PARTICALSYSTEM      CocosModules = "ParticleSystem"
	CCMO_RICHTEXT            CocosModules = "RichText"
	CCMO_RENDERTEXUTURE      CocosModules = "Renderer Texture"
	CCMO_SLIDER              CocosModules = "Slider"
	CCMO_SCROLLBAR           CocosModules = "ScrollBar"
	CCMO_SCROLLVIEW          CocosModules = "ScrollView"
	CCMO_SPINESKELETON       CocosModules = "Spine Skeleton"
	CCMO_STUDIOCOMPONENT     CocosModules = "StudioComponent"
	CCMO_TOGGLE              CocosModules = "Toggle"
	CCMO_TILEMAP             CocosModules = "TiledMap"
	CCMO_VIDEOPLAYER         CocosModules = "VideoPlayer"
	CCMO_WIDGET              CocosModules = "Widget"
	CCMO_WEBVIEW             CocosModules = "WebView"
	CCMO_WEBGLRENDERER       CocosModules = "WebGL Renderer"
	CCMO_3D                  CocosModules = "3D"
	CCMO_SUBCONTEXT          CocosModules = "SubContext"
	CCMO_TYPESCRIPTPOLLYFILL CocosModules = "TypeScript Polyfill"
	CCMO_SAFEAREA            CocosModules = "SafeArea"

	CC_PORTRAIT       CocosOrientation = "portrait"
	CC_UPSIDEDOWN     CocosOrientation = "upsideDown"
	CC_LANDSCAPELEFT  CocosOrientation = "landscapeLeft"
	CC_LANDSCALERIGHT CocosOrientation = "landscapeRight"
)

type CocosBuildOption struct {
	EnginePath string
	Path            string
	BuildPath       string
	ExcludedModules []CocosModules
	Platform        CocosPlatform
	Debug           bool
	PreviewWidth    int
	PreviewHeight   int
	SourceMaps      bool
	StartScene      string
	WebOrientation  CocosOrientation
	Md5Cache        bool
}

const COCOS_WIN_BUILDER = "CocosCreator.exe"
const COCOS_MAC_BUILDER = "CocosCreator.app/Contents/MacOS/CocosCreator"

func BuildCocos(conf *CocosBuildOption,writer io.Writer) error {
	var builderStr  = conf.EnginePath
	if runtime.GOOS == "darwin" {
		builderStr = path.Join(builderStr,COCOS_MAC_BUILDER)
	} else if runtime.GOOS == "windows" {
		builderStr = path.Join(builderStr,COCOS_WIN_BUILDER)
	}
	params := ""
	log.Infof("engine path",builderStr)
	_,err:=writer.Write([]byte("engine path"+builderStr+"\n"))
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if len(conf.ExcludedModules) > 0 {
		excludeString := "["
		for _, v := range conf.ExcludedModules {
			excludeString += string(v)
			excludeString += ","
		}
		excludeString = strings.TrimRight(excludeString, ",")
		excludeString += "]"
		params += fmt.Sprintf("excludedModules=%s;", excludeString)

	}
	params += fmt.Sprintf("platform=%s;", conf.Platform)
	params += fmt.Sprintf("debug=%v;", conf.Debug)
	params += fmt.Sprintf("sourceMaps=%v;", conf.SourceMaps)
	params += fmt.Sprintf("startScene=%s;", conf.StartScene)
	if conf.BuildPath != "" {
		params += fmt.Sprintf("buildPath=%s;", conf.BuildPath)
	}
	cmd := exec.Command(builderStr, "--force", "--path", conf.Path, "--build", params)
	stdout, _ := cmd.StdoutPipe()
	log.Infof(cmd.String())
	err=cmd.Start()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		if writer!=nil {
			writer.Write([]byte(line))
		}
	}
	err=cmd.Wait()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
