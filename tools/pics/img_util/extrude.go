package img_util

import (
	"github.com/nomos/go-lokas/log"
	"image"
)

func ExtrudeImage(img image.Image, width, height int, extrude bool) (image.Image, error) {
	rectangle := img.Bounds()
	imgWidth := rectangle.Max.X - rectangle.Min.X
	imgHeight := rectangle.Max.Y - rectangle.Min.Y
	if !extrude {
		width += 2
		height += 2
	}
	widthSegs := imgWidth / width
	heightSegs := imgHeight / height

	if imgWidth%width > 0 {
		return nil, log.Error("width not fit")
	}
	if imgHeight%height > 0 {
		return nil, log.Error("height not fit")
	}
	outWidth := imgWidth + widthSegs*2
	outHeight := imgHeight + heightSegs*2
	if !extrude {
		outWidth = imgWidth - widthSegs*2
		outHeight = imgHeight - heightSegs*2
	}
	rectangle = image.Rect(0, 0, outWidth, outHeight)
	exportImage := image.NewNRGBA(rectangle)
	for y := 0; y < heightSegs; y++ {
		for x := 0; x < widthSegs; x++ {
			for y1 := 0; y1 < height; y1++ {
				for x1 := 0; x1 < width; x1++ {
					if extrude {

						extrudeFunc(x, y, x1, y1, width, height, img, exportImage)
					} else {

						unextrudeFunc(x, y, x1, y1, width, height, img, exportImage)
					}
				}
			}
		}
	}
	return exportImage, nil

}
func unextrudeFunc(x, y, x1, y1, width, height int, img image.Image, exportImage *image.NRGBA) {
	if y1 == 0 {
		//TOP
		return
	} else if y1 == height-1 {
		//BOTTOM
		return
	}
	distY := y*(height-2) + y1
	originY := y*height + y1
	if x1 == 0 {
		//LEFT
		return
	} else if x1 == width-1 {
		//RIGHT
		return
	}
	distX := x*(width-2) + x1
	originX := x*width + x1
	exportImage.Set(distX, distY, img.At(originX, originY))

}

func extrudeFunc(x, y, x1, y1, width, height int, img image.Image, exportImage *image.NRGBA) {
	if y1 == 0 {
		//TOP
		distYTop := y * (height + 2)
		distX := x*(width+2) + 1 + x1
		originYTop := y * height
		originX := x*width + x1
		exportImage.Set(distX, distYTop, img.At(originX, originYTop))
	} else if y1 == height-1 {
		//BOTTOM
		distYBottom := y*(height+2) + height + 1
		distX := x*(width+2) + 1 + x1
		originYBottom := y*height + y1
		originX := x*width + x1
		exportImage.Set(distX, distYBottom, img.At(originX, originYBottom))
	}
	distY := y*(height+2) + 1 + y1
	originY := y*height + y1
	if x1 == 0 {
		//LEFT
		distXLeft := x * (width + 2)
		originXLeft := x * width
		exportImage.Set(distXLeft, distY, img.At(originXLeft, originY))
	} else if x1 == width-1 {
		//RIGHT
		distXRight := x*(width+2) + width + 1
		originXRight := x*width + x1
		exportImage.Set(distXRight, distY, img.At(originXRight, originY))
	}
	distX := x*(width+2) + 1 + x1
	originX := x*width + x1
	exportImage.Set(distX, distY, img.At(originX, originY))

}
