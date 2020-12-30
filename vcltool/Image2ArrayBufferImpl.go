// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
    "github.com/atotto/clipboard"
    "github.com/nomos/go-tools/pics/img_png"
    "github.com/ying32/govcl/vcl"
    _ "github.com/ying32/govcl/vcl/types"
    "path"
    "strings"
)

//::private::
type TImage2ArrayBufferFields struct {
    ConfigAble
}

func (this *TImage2ArrayBuffer) OnCreate(){
    this.listener.On("DropFile", func(i ...interface{}) {
        if this.container.IsFrameSelected(this) {
            paths:=i[0].([]string)
            this.log.Warnf("DropFile",paths)
            filePath:=paths[0]

            w,h,data,err:=img_png.Png2CompressedBase64(filePath)
            if err != nil {
                this.log.Error(err.Error())
                return
            }
            this.log.Infof("width:",w)
            this.log.Infof("height:",h)
            this.log.Infof("data:",data)
            fileName:=strings.Replace(path.Base(filePath),".png","",-1)
            clipboard.WriteAll(fileName+" : `"+data+"`")
            this.log.Info("已拷贝到剪切板")
            //if path.Ext(filePath) == ".png" {
            //    outPath:= strings.Replace(filePath,".png",".txt",1)
            //    err:=ioutil.WriteFile(outPath,[]byte(strconv.Itoa(w)+" "+strconv.Itoa(h)+"\n"+data),0666)
            //    if err != nil {
            //        this.log.Error(err.Error())
            //    }
            //}
        }
    })
}

func (this *TImage2ArrayBuffer) OnDestroy(){

}

func (this *TImage2ArrayBuffer) OnLabel1Click(sender vcl.IObject) {

}

