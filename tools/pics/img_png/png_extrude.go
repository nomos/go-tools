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

func UnExtrudeImage(imgPath string, width, height int, logger log.ILogger) error {
	if logger == nil {
		logger = log.DefaultLogger()
	}
	img, err := ReadImageFile(imgPath)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	out, err := img_util.ExtrudeImage(img, width, height, false)
	ext := path.Ext(imgPath)
	path1 := strings.Replace(imgPath, ext, "", 1)

	path2 := path.Join(path1 + "_unbleed.png")

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
	err = png.Encode(f, out)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func ExtrudeImage(imgPath string, width, height int, logger log.ILogger) error {
	if logger == nil {
		logger = log.DefaultLogger()
	}
	img, err := ReadImageFile(imgPath)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	out, err := img_util.ExtrudeImage(img, width, height, true)
	ext := path.Ext(imgPath)
	path1 := strings.Replace(imgPath, ext, "", 1)

	path2 := path.Join(path1 + "_bleed.png")

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
	err = png.Encode(f, out)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
