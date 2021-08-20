package img_tool

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/tools/pics/img_dds"
	"github.com/nomos/go-tools/ui"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"image"
	"image/png"
	"os"
	"strings"
)

var _ ui.IFrame = (*DDSConverter)(nil)

type DDSConverter struct {
	*vcl.TFrame
	ui.ConfigAble
}

func NewDDSConverter(owner vcl.IWinControl,option... ui.FrameOption) (root *DDSConverter)  {
	vcl.CreateResFrame(owner, &root)
	for _,o:=range option {
		o(root)
	}
	return
}

func (this *DDSConverter) setup(){
	this.SetAlign(types.AlClient)
	this.SetConfig(this.Config().Sub("dds"))
	line2:=ui.CreateLine(types.AlTop,44,this)
	line1:=ui.CreateLine(types.AlTop,44,this)
	dirBtn:=ui.NewOpenPathBar(line1,"DDS目录",280)
	dirBtn.OnCreate()
	dirBtn.SetParent(line1)
	transBtn:=ui.CreateButton("转换DDS",line2)
	tag:=ui.NewEditLabel(line2,"Bgrflip",100,ui.EDIT_TYPE_BOOL)
	tag.OnCreate()
	tag.SetParent(line2)
	tag.SetAlign(types.AlLeft)
	dirBtn.SetOpenDirDialog("DDS目录")
	if this.Config().GetString("dds_folder")!="" {
		dirBtn.SetPath(this.Config().GetString("dds_folder"))
	}
	tag.SetBool(this.Config().GetBool("bgr_flip"))
	dirBtn.OnEdit = func(s string) {
		if util.IsFileExist(s) {
			this.Config().Set("dds_folder",s)
		}
	}
	dirBtn.OnOpen = func(s string) {
		if util.IsFileExist(s) {
			this.Config().Set("dds_folder",s)
		}
	}
	transBtn.SetOnClick(func(sender vcl.IObject) {
		this.convertDDSFolder(this.Config().GetString("dds_folder"),this.Config().GetBool("bgr_flip"))
	})
	tag.OnValueChange = func(label *ui.EditLabel, editType ui.EDIT_TYPE, value interface{}) {
		this.Config().Set("bgr_flip",tag.Bool())
	}
}

func (this *DDSConverter) convertDDSFolder(filePath string,flip bool)error{
	files,err:=util.WalkDirFiles(filePath,true)
	if err != nil {
		return err
	}
	file1:=util.FilterFileWithExt(files,".dds")
	for _,file:=range file1 {
		this.convertDDS(file,flip)
	}
	return nil
}

func (this *DDSConverter) convertDDS(filePath string,flip bool){
	file, err := os.Open(filePath)
	if err != nil {
		log.Panicf(err)
	}
	defer file.Close()
	tex, err := img_dds.Decode(file,flip)
	if err != nil {
		log.Errorf("decode errror", err)
		return
	}
	this.GetLogger().Infof("decoded",tex.Width, tex.Height, len(tex.Data),filePath)
	image := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{tex.Width, tex.Height}})
	image.Pix = tex.Data
	file1:=strings.Replace(filePath,".dds",".png",1)
	out, _ := os.Create(file1)
	png.Encode(out, image)
}

func (this *DDSConverter) OnCreate(){
	this.setup()
}

func (this *DDSConverter) OnDestroy(){

}

func (this *DDSConverter) OnEnter(){

}

func (this *DDSConverter) OnExit(){

}

func (this *DDSConverter) Clear(){

}



