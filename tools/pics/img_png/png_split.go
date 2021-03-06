package img_png

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"image"
	"image/png"
	"math"
	"os"
	"path"
	"sort"
	"strings"
)

func splitImage(img image.Image, width, height int) []image.Image {
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

func SplitImage(imgPath string, width, height int, logger log.ILogger) error {
	if logger == nil {
		logger = log.DefaultLogger()
	}
	img, err := ReadImageFile(imgPath)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	imgs := splitImage(img, width, height)
	ext := path.Ext(imgPath)
	path1 := strings.Replace(imgPath, ext, "", 1)
	digital := util.DigitalNum(len(imgs) + 1)
	if digital < 2 {
		digital = 2
	}
	logger.Warnf("digital", digital)
	exist, err := util.PathExists(path1)
	if err != nil {
		return err
	}
	if !exist {
		err = os.Mkdir(path1, 0777)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	for index, imga := range imgs {
		path2 := path.Join(path1, util.DigitalToString(index+1, digital)+".png")
		exist, err := util.PathExists(path2)
		if err != nil {
			return err
		}
		if !exist {
			_, err = os.Create(path2)
		}
		f, err := os.OpenFile(path2, os.O_RDWR, 0666)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		err = png.Encode(f, imga)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		err = f.Close()
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		logger.Warnf("????????????", path2)
	}
	return nil
}
