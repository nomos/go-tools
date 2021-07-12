package pics

import (
	"github.com/nomos/go-lokas/log"
	"golang.org/x/image/draw"
	"image"
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
	draw.NearestNeighbor.Scale(dst,distBound,src,bound,draw.Over,nil)
	return dst
}
