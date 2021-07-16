package pics

import (
	"bytes"
	"github.com/nomos/go-lokas/log"
	"golang.org/x/image/draw"
	"image"
	"image/png"
)

func Resize(src image.Image,width,height int)image.Image{
	bound := src.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	log.Warnf(dx,dy)
	// 缩略图的大小
	distBound:=image.Rect(0,0,width,height)
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	// 产生缩略图,等比例缩放
	draw.ApproxBiLinear.Scale(dst,distBound,src,bound,draw.Over,nil)
	return dst
}


func ResizePng(data []byte,width,height int)[]byte{
	img,_:=png.Decode(bytes.NewBuffer(data))
	img2:=Resize(img,width,height)
	rData:=make([]byte,0,1024*1024*8)
	ret:=bytes.NewBuffer(rData)
	png.Encode(ret,img2)
	return ret.Bytes()
}
