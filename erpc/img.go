package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/rpc"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/tools/pics/img_dds"
	"github.com/nomos/go-tools/tools/pics/img_png"
	"image"
	"image/png"
	"os"
	"strings"
)

func convertDDS(filePath string, flip bool,logger log.ILogger) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Panicf(err)
	}
	defer file.Close()
	tex, err := img_dds.Decode(file, flip)
	if err != nil {
		logger.Errorf("decode errror", err)
		return
	}
	logger.Infof("decoded", tex.Width, tex.Height, len(tex.Data), filePath)
	image := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{tex.Width, tex.Height}})
	image.Pix = tex.Data
	file1 := strings.Replace(filePath, ".dds", ".png", 1)
	out, _ := os.Create(file1)
	png.Encode(out, image)
}

func convertDDSFolder(filePath string, flip bool,logger log.ILogger) error {
	files, err := util.WalkDirFiles(filePath, true)
	if err != nil {
		return err
	}
	file1 := util.FilterFileWithExt(files, ".dds")
	for _, file := range file1 {
		convertDDS(file, flip,logger)
	}
	return nil
}

func init() {
	rpc.RegisterAdminFunc(CUT_PNG, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		p := params.String()
		width := params.Int()
		height := params.Int()
		err := img_png.SplitImage(p, width, height, logger)
		if err != nil {
			log.Error(err.Error())
		}
		return nil, nil
	})
	rpc.RegisterAdminFunc(CONVERT_DDS, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		p := params.String()
		flip := params.Bool()
		convertDDSFolder(p,flip,logger)
		return nil,nil
	})
}
