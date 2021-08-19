package img_tool

import (
    "github.com/nomos/go-lokas/log"
    "github.com/nomos/go-tools/tools/pics/img_png"
    "github.com/nomos/go-tools/ui"
    "github.com/nomos/go-tools/ui/icons"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
    "strconv"
)

type TImageCutter struct {
    *vcl.TFrame
    PngPath        *vcl.TEdit
    OpenPngButton  *vcl.TSpeedButton
    Label1         *vcl.TLabel
    GenerateButton *vcl.TButton
    ImageWidth     *vcl.TEdit
    ImageHeight    *vcl.TEdit
    Label2         *vcl.TLabel
    Label3         *vcl.TLabel

    ui.ConfigAble
}


func NewImageCutter(owner vcl.IComponent) (root *TImageCutter)  {
    vcl.CreateResFrame(owner, &root)
    return
}

func (this *TImageCutter) setup(){
    this.SetAlign(types.AlClient)
    line2:=ui.CreateLine(types.AlTop,32,this)
    line1:=ui.CreateLine(types.AlTop,32,this)
    this.PngPath = ui.CreateEdit(200,line1)
    this.OpenPngButton = ui.CreateSpeedBtn("folder",icons.GetImageList(32,32),line1)
    this.Label1 = ui.CreateText("图片路径",line1)
    this.GenerateButton = ui.CreateButton("生成图片",line2)
    this.ImageHeight = ui.CreateEdit(100,line2)
    this.Label2 = ui.CreateText("高度",line2)
    this.ImageWidth = ui.CreateEdit(100,line2)
    this.Label3 = ui.CreateText("宽度",line2)
}

func (this *TImageCutter) OnCreate(){
    this.setup()
    openDialog := vcl.NewOpenDialog(this)
    openDialog.SetOptions(openDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
    openDialog.SetFilter("图片(*.png,*.*)|*.png;*.*|所有文件(*.*)|*.*")
    openDialog.SetTitle("打开文件")
    if this.Config().GetString("png_path") != "" {
        this.PngPath.SetText(this.Config().GetString("png_path"))
    }
    this.OpenPngButton.SetOnClick(func(sender vcl.IObject) {
        if this.Config().GetString("png_path") != "" {
            openDialog.SetInitialDir(this.Config().GetString("png_path"))
        }
        if openDialog.Execute() {
            p := openDialog.FileName()
            this.Config().Set("png_path", p)
            log.Warnf(p)
            this.PngPath.SetText(p)
        }
    })
    if this.Config().GetString("png_width") != "" {
        this.ImageWidth.SetText(this.Config().GetString("png_width"))
    }
    if this.Config().GetString("png_height") != "" {
        this.ImageHeight.SetText(this.Config().GetString("png_height"))
    }
    this.ImageWidth.SetOnChange(func(sender vcl.IObject) {
        text:=this.ImageWidth.Text()
        _,err:=strconv.Atoi(text)
        if err!=nil {
            this.ImageWidth.SetText("0")
            return
        }
        this.Config().Set("png_width",text)
    })
    this.ImageHeight.SetOnChange(func(sender vcl.IObject) {
        text:=this.ImageHeight.Text()
        _,err:=strconv.Atoi(text)
        if err!=nil {
            this.ImageHeight.SetText("0")
            return
        }
        this.Config().Set("png_height",text)
    })
    this.GenerateButton.SetOnClick(func(sender vcl.IObject) {
        p:=this.Config().GetString("png_path")
        width,err := strconv.Atoi(this.Config().GetString("png_width"))
        if err != nil||width<=0 {
            log.Errorf("width is not fit",width)
            return
        }
        height,err := strconv.Atoi(this.Config().GetString("png_height"))
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
