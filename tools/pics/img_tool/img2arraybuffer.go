package img_tool

import (
    "github.com/atotto/clipboard"
    "github.com/nomos/go-lokas/log"
    "github.com/nomos/go-lokas/promise"
    "github.com/nomos/go-tools/tools/pics/img_png"
    "github.com/nomos/go-tools/ui"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
    "path"
    "strings"
    "time"
)

type TImage2ArrayBuffer struct {
    *vcl.TFrame
    DropDownPanel     *vcl.TPanel
    Label1            *vcl.TLabel
    dropFile string
    entered bool
    ui.ConfigAble
}

func (this *TImage2ArrayBuffer) OnEnter() {

}

func (this *TImage2ArrayBuffer) OnExit() {

}

func (this *TImage2ArrayBuffer) Clear() {

}

func NewImage2ArrayBuffer(owner vcl.IComponent,option... ui.FrameOption) (root *TImage2ArrayBuffer)  {
    vcl.CreateResFrame(owner, &root)
    for _,o:=range option {
        o(root)
    }
    return
}

var _ ui.IFrame = (*TImage2ArrayBuffer)(nil)

func (this *TImage2ArrayBuffer) setup(){
    this.SetAlign(types.AlClient)
    this.DropDownPanel = ui.CreatePanel(types.AlClient,this)
    this.DropDownPanel.BorderSpacing().SetAround(24)
    this.Label1 = ui.CreateText("拖动图片到此处",this.DropDownPanel)
    this.Label1.SetAlign(types.AlNone)
    this.Label1.AnchorHorizontalCenterTo(this.DropDownPanel)
    this.Label1.AnchorVerticalCenterTo(this.DropDownPanel)
}

func (this *TImage2ArrayBuffer) OnCreate(){
    this.setup()
    this.DropDownPanel.SetOnMouseEnter(func(sender vcl.IObject) {
        log.Info("enter drop down")
        this.entered = true
        filePath:=this.dropFile
        if filePath=="" {
            return
        }

        w,h,data,err:=img_png.Png2CompressedBase64(filePath)
        if err != nil {
            this.GetLogger().Error(err.Error())
            return
        }
        this.GetLogger().Infof("width:",w)
        this.GetLogger().Infof("height:",h)
        this.GetLogger().Infof("data:",data)
        fileName:=strings.Replace(path.Base(filePath),".png","",-1)
        clipboard.WriteAll(fileName+" : `"+data+"`,")
        this.GetLogger().Info("已拷贝到剪切板")
        //if path.Ext(filePath) == ".png" {
        //    outPath:= strings.Replace(filePath,".png",".txt",1)
        //    err:=ioutil.WriteFile(outPath,[]byte(strconv.Itoa(w)+" "+strconv.Itoa(h)+"\n"+data),0666)
        //    if err != nil {
        //        this.GetLogger().Error(err.Error())
        //    }
        //}
    })
    this.DropDownPanel.SetOnMouseLeave(func(sender vcl.IObject) {
        this.entered = false
    })

    this.GetListener().On("DropFile", func(i ...interface{}) {
        log.Warnf("listener")
        time.Sleep(1000)
        paths:=i[0].([]string)
        this.GetLogger().Warnf("DropFile",paths)
        this.dropFile = paths[0]
        promise.SetTimeout(time.Second*1, func(timeout *promise.Timeout) {
            this.dropFile = ""
        })
    })
}

func (this *TImage2ArrayBuffer) OnDestroy(){

}

func (this *TImage2ArrayBuffer) OnLabel1Click(sender vcl.IObject) {

}
