// 由res2go自动生成。
// 在这里写你的事件。

package vcltool

import (
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-tools/pics/img_png"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"strconv"
)

//::private::
type TImageCutterFields struct {
	ConfigAble
}

func (this *TImageCutter) OnCreate(){
	openDialog := vcl.NewOpenDialog(this)
	openDialog.SetOptions(openDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	openDialog.SetFilter("图片(*.png,*.*)|*.png;*.*|所有文件(*.*)|*.*")
	openDialog.SetTitle("打开文件")
	if this.conf.GetString("png_path") != "" {
		this.PngPath.SetText(this.conf.GetString("png_path"))
	}
	this.OpenPngButton.SetOnClick(func(sender vcl.IObject) {
		if this.conf.GetString("png_path") != "" {
			openDialog.SetInitialDir(this.conf.GetString("png_path"))
		}
		if openDialog.Execute() {
			p := openDialog.FileName()
			this.conf.Set("png_path", p)
			log.Warnf(p)
			this.PngPath.SetText(p)
		}
	})
	if this.conf.GetString("png_width") != "" {
		this.ImageWidth.SetText(this.conf.GetString("png_width"))
	}
	if this.conf.GetString("png_height") != "" {
		this.ImageHeight.SetText(this.conf.GetString("png_height"))
	}
	this.ImageWidth.SetOnChange(func(sender vcl.IObject) {
		text:=this.ImageWidth.Text()
		_,err:=strconv.Atoi(text)
		if err!=nil {
			this.ImageWidth.SetText("0")
			return
		}
		this.conf.Set("png_width",text)
	})
	this.ImageHeight.SetOnChange(func(sender vcl.IObject) {
		text:=this.ImageHeight.Text()
		_,err:=strconv.Atoi(text)
		if err!=nil {
			this.ImageHeight.SetText("0")
			return
		}
		this.conf.Set("png_height",text)
	})
	this.GenerateButton.SetOnClick(func(sender vcl.IObject) {
		p:=this.conf.GetString("png_path")
		width,err := strconv.Atoi(this.conf.GetString("png_width"))
		if err != nil||width<=0 {
			log.Errorf("width is not fit",width)
			return
		}
		height,err := strconv.Atoi(this.conf.GetString("png_height"))
		if err != nil||width<=0 {
			log.Errorf("height is not fit",height)
			return
		}
		if err != nil {
			return
		}
		err=img_png.SubImage(p,width,height)
		if err != nil {
			log.Error(err.Error())
		}
	})
}

func (this *TImageCutter) OnDestroy(){

}
