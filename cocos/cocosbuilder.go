package cocos

import (
	"bufio"
	"fmt"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/protocol"
	"io"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

var _cocosPlatforms = map[protocol.Enum]string{}
func _addPlatform(enum protocol.Enum,s string)CocosPlatform{
	_cocosPlatforms[enum] = s
	return CocosPlatform(enum)
}
func (this CocosPlatform) ToString()string{
	return _cocosPlatforms[this.Enum()]
}
func (this CocosPlatform) Enum()protocol.Enum{
	return protocol.Enum(this)
}
func _addOrientation(enum protocol.Enum,s string)CocosOrientation{
	_cocosOrientations[enum] = s
	return CocosOrientation(enum)
}
func (this CocosOrientation) ToString()string{
	return _cocosOrientations[this.Enum()]
}
func (this CocosOrientation) Enum()protocol.Enum{
	return protocol.Enum(this)
}

func _addWebOrientation(enum protocol.Enum,s string)CocosWebOrientation{
	_cocosWebOrientations[enum] = s
	return CocosWebOrientation(enum)
}
func (this CocosWebOrientation) ToString()string{
	return _cocosWebOrientations[this.Enum()]
}
func (this CocosWebOrientation) Enum()protocol.Enum{
	return protocol.Enum(this)
}

var _cocosOrientations = map[protocol.Enum]string{}
var _cocosWebOrientations = map[protocol.Enum]string{}

type CocosModules string

type CocosPlatform protocol.Enum

type CocosOrientation protocol.Enum

type CocosWebOrientation protocol.Enum

var (
	CC_WEB_MOBILE  CocosPlatform = _addPlatform(1,"web-mobile")
	CC_WEB_DESKTOP CocosPlatform = _addPlatform(2,"web-desktop")
	CC_ANDROID CocosPlatform = _addPlatform(3,"android")
	CC_IOS CocosPlatform = _addPlatform(4,"ios")
	CC_MAC CocosPlatform = _addPlatform(5,"mac")
	CC_WIN32 CocosPlatform = _addPlatform(6,"win32")
	CC_WECHATGAME CocosPlatform = _addPlatform(7,"wechatgame")
	ALL_CC_PLATFORM protocol.IEnumCollection = []protocol.IEnum{CC_WEB_MOBILE,CC_WEB_DESKTOP,CC_ANDROID,CC_IOS,CC_MAC,CC_WIN32,CC_WECHATGAME}
)

var (
	CC_PORTRAIT       CocosOrientation = _addOrientation(1, "portrait")
	CC_UPSIDEDOWN     CocosOrientation = _addOrientation(2,"upsideDown")
	CC_LANDSCAPELEFT  CocosOrientation = _addOrientation(3,"landscapeLeft")
	CC_LANDSCALERIGHT CocosOrientation = _addOrientation(4,"landscapeRight")
	ALL_CC_ORIENTATION protocol.IEnumCollection = []protocol.IEnum{CC_PORTRAIT,CC_UPSIDEDOWN,CC_LANDSCAPELEFT,CC_LANDSCALERIGHT}
)

var (
	CC_WEB_PORTRAIT       CocosWebOrientation = _addWebOrientation(1, "portrait")
	CC_WEB_LANDSCAPE     CocosWebOrientation = _addWebOrientation(2,"landscape")
	CC_WEB_AUTO  CocosWebOrientation = _addWebOrientation(3,"auto")
	ALL_CC_WEB_ORIENTATION protocol.IEnumCollection = []protocol.IEnum{CC_WEB_PORTRAIT,CC_WEB_LANDSCAPE,CC_WEB_AUTO}
)

const (

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

)





type CocosBuildOption struct {
	EnginePath      string
	Path            string
	BuildPath       string
	ExcludedModules []CocosModules
	Platform        CocosPlatform
	Debug           bool
	PreviewWidth    int
	PreviewHeight   int
	SourceMaps      bool
	StartScene      string
	Orientation     CocosOrientation
	WebOrientation  CocosWebOrientation
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
	if conf.Platform==CC_WEB_MOBILE {
		params += fmt.Sprintf("webOrientation=%s;", conf.WebOrientation.ToString())
	} else {
		//params += fmt.Sprintf("orientation=%s;", conf.Orientation.ToString())
	}
	params += fmt.Sprintf("platform=%s;", conf.Platform.ToString())
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
