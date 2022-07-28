package img_png

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/tools/pics/img_util"
	"image/png"
	"os"
	"path"
	"strings"
)

func SplitImage(imgPath string, width, height int, logger log.ILogger) error {
	if logger == nil {
		logger = log.DefaultLogger()
	}
	img, err := ReadImageFile(imgPath)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	imgs := img_util.SplitImage(img, width, height)
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
		logger.Warnf("输出图片", path2)
	}
	return nil
}
