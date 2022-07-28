package img_util

import (
	"github.com/nomos/go-lokas/util"
	"image"
	"math"
	"sort"
)

func SplitImage(img image.Image, width, height int) []image.Image {
	rectangle := img.Bounds()
	imgWidth := rectangle.Max.X - rectangle.Min.X
	imgHeight := rectangle.Max.Y - rectangle.Min.Y
	xSplit := int(math.Ceil(float64(imgWidth) / float64(width)))
	ySplit := int(math.Ceil(float64(imgHeight) / float64(height)))
	ret := make([]image.Image, 0)
	rgbImg := img.(*image.NRGBA)
	isEmptyImg := func(img image.Image) bool {
		ret := true
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				_, _, _, a := img.At(x, y).RGBA()
				if a > 0 {
					ret = false
				}
			}
		}
		return ret
	}
	for y := 0; y < ySplit; y++ {
		for x := 0; x < xSplit; x++ {
			x1 := util.Ternary((x+1)*width > imgWidth, imgWidth, (x+1)*width)
			y1 := util.Ternary((y+1)*height > imgHeight, imgHeight, (y+1)*height)
			if (x+1)*width > imgWidth {
				continue
			}
			if (y+1)*height > imgHeight {
				continue
			}

			img2 := rgbImg.SubImage(image.Rect(x*width, y*height, x1, y1)).(*image.NRGBA)
			if !isEmptyImg(img2) {
				ret = append(ret, img2)
			}
		}
	}
	lengthFunc := func(x, y int) float64 {
		return math.Pow(float64(x), 2) + math.Pow(float64(y), 2)
	}
	sort.Slice(ret, func(i, j int) bool {
		return lengthFunc(ret[i].Bounds().Min.X, ret[i].Bounds().Min.Y) < lengthFunc(ret[j].Bounds().Min.X, ret[j].Bounds().Min.Y)
	})
	return ret
}
