package icons

import (
	"github.com/nomos/go-tools/tools/pics"
	"github.com/nomos/go-tools/tools/pics/img_bmp"
	"github.com/nomos/go-tools/ui/icons/icon_assets"
	"github.com/nomos/go-tools/ui/icons/pix_icon_assets"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"strconv"
	"sync"
)

var iconList = map[string][]byte{}
var _listForLit = map[string]*ImageList{}

func init(){
	pix_icon_assets.Assign(iconList)
	icon_assets.Assign(iconList)
}

func GetImage(control vcl.IWinControl,width,height int,s string)*vcl.TImage{
	ret:=vcl.NewImage(control)
	ret.SetHeight(15)
	rdata:=pics.ResizePng(iconList[s],width,height)
	ret.Picture().LoadFromBytes(rdata)
	return ret
}


func LoadData(img *vcl.TImage,s string){
	img.Picture().LoadFromBytes(iconList[s])
}



var _mu sync.Mutex


func GetImageList(width,height int32)*ImageList{
	_mu.Lock()
	defer _mu.Unlock()
	index:=strconv.Itoa(int(width))+"_"+strconv.Itoa(int(height))
	if _listForLit[index]==nil {
		_listForLit[index] = NewImageList(width,height)
	}
	return _listForLit[index]
}

type ImageList struct {
	vclList *vcl.TImageList
	idx int32
	idxMap map[string]int32
}

func NewImageList(width,height int32)*ImageList {
	ret:=&ImageList{
		idx: 0,
		idxMap: map[string]int32{},
	}
	ret.vclList = vcl.NewImageList(vcl.Application)
	ret.setup(width,height)
	return ret
}

func (this *ImageList) setup(width,height int32){
	this.vclList.SetWidth(width)
	this.vclList.SetHeight(height)
	for s,d:=range iconList {
		this.AddPng(d)
		this.idxMap[s] = this.idx
		this.idx++
	}
}

func (this *ImageList) GetImageIndex(s string)int32{
	return this.idxMap[s]
}

func (this *ImageList) GetPixImageIndex(s string)int32{
	return this.idxMap["pix_"+s]
}

func (this *ImageList) ImageList()*vcl.TImageList{
	return this.vclList
}

func (this *ImageList) AddPng(data []byte){
	imgBmp,_:=img_bmp.Png2Bmp(data)
	vclBmp:=vcl.NewBitmap()
	vclBmp.LoadFromBytes(imgBmp)
	this.vclList.AddMasked(vclBmp,types.TColor(0xFFFFFFFF))
}