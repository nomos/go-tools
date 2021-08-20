package img_tool

import (
    "github.com/nomos/go-lokas/log"
    "github.com/nomos/go-lokas/util"
    "github.com/nomos/go-tools/tools/pics/img_png"
    "github.com/nomos/go-tools/ui"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
    "strconv"
)

type TImageCutter struct {
    *vcl.TFrame
    OpenPngButton  *ui.OpenPathBar
    GenerateButton *vcl.TButton
    ImageWidth     *ui.EditLabel
    ImageHeight    *ui.EditLabel

    ui.ConfigAble
}

var _ ui.IFrame = (*TImageCutter)(nil)

func NewImageCutter(owner vcl.IComponent,option... ui.FrameOption) (root *TImageCutter)  {
    vcl.CreateResFrame(owner, &root)
    for _,o:=range option {
        o(root)
    }
    return
}

func (this *TImageCutter) OnEnter() {
}

func (this *TImageCutter) OnExit() {
}

func (this *TImageCutter) Clear() {
}

func (this *TImageCutter) setup(){
    this.SetAlign(types.AlClient)
    line2:=ui.CreateLine(types.AlTop,44,this)
    this.GenerateButton = ui.CreateButton("生成图片",line2)
    this.GenerateButton.BorderSpacing().SetTop(12)
    line22:=ui.CreateLine(types.AlLeft,100,line2)
    line21:=ui.CreateLine(types.AlLeft,100,line2)
    line1:=ui.CreateLine(types.AlTop,44,this)
    this.OpenPngButton = ui.NewOpenPathBar(line1,"图片路径",280)
    this.OpenPngButton.SetParent(line1)
    this.OpenPngButton.OnCreate()
    this.ImageHeight = ui.NewEditLabel(line21,"高度",100,ui.EDIT_TYPE_INTERGER)
    this.ImageHeight.SetAlign(types.AlLeft)
    this.ImageHeight.SetParent(line21)
    this.ImageHeight.OnCreate()
    this.ImageWidth = ui.NewEditLabel(line22,"宽度",100,ui.EDIT_TYPE_INTERGER)
    this.ImageWidth.SetAlign(types.AlLeft)
    this.ImageWidth.SetParent(line22)
    this.ImageWidth.OnCreate()

}

func (this *TImageCutter) OnCreate(){
    this.setup()
    this.OpenPngButton.SetOpenFileDialog("打开文件","图片(*.png,*.*)|*.png;*.*|所有文件(*.*)|*.*")
    if this.Config().GetString("png_path") != "" {
        this.OpenPngButton.SetPath(this.Config().GetString("png_path"))
    }
    this.OpenPngButton.OnOpen = func(s string) {
        if util.IsFileExist(s) {
            this.Config().Set("png_path", s)
            this.OpenPngButton.SetPath(s)
        }
    }
    this.OpenPngButton.OnEdit = func(s string) {
        if util.IsFileExist(s) {
            this.Config().Set("png_path", s)
            this.OpenPngButton.SetPath(s)
        }
    }
    if this.Config().GetString("png_width") != "" {
        this.ImageWidth.SetString(this.Config().GetString("png_width"))
    }
    if this.Config().GetString("png_height") != "" {
        this.ImageHeight.SetString(this.Config().GetString("png_height"))
    }
    this.ImageWidth.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
        text:=this.ImageWidth.String()
        _,err:=strconv.Atoi(text)
        if err!=nil {
            this.ImageWidth.SetString("0")
            return
        }
        this.Config().Set("png_width",text)
    }
    this.ImageHeight.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
        text:=this.ImageHeight.String()
        _,err:=strconv.Atoi(text)
        if err!=nil {
            this.ImageHeight.SetString("0")
            return
        }
        this.Config().Set("png_height",text)
    }
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
