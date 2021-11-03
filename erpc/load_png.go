package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/tools/pics/img_png"
	"os"
)

func loadPngFile(path string,logger log.ILogger)(*PngFile,error) {
	file,err:=os.Stat(path)
	if err != nil {
		logger.Error(err.Error())
		return nil,err
	}

	w,h,data,err:=img_png.Png2ImageMap(path)
	if err != nil {
		logger.Error(err.Error())
		return nil,err
	}
	return NewPngFile(path,file.ModTime(),int32(w),int32(h),data),nil
}

func init(){
	RegisterAdminFunc(LOAD_PNG_FOLDER, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		ret:=NewPngFolderData()
		path:=params.String()
		util.WalkDirFilesWithFunc(path, func(filePath string, file os.FileInfo) bool {
			if util.IsFileWithExt(filePath,"png") {
				ret.AddFile(NewPngFile(filePath,file.ModTime(),0,0,nil))
			}
			file.ModTime()
			return false
		},true)
		return protocol.MarshalBinary(ret)
	})
	RegisterAdminFunc(LOAD_PNG_DATA, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		path:=params.String()
		ret,err:=loadPngFile(path,logger)
		if err != nil {
			logger.Error(err.Error())
			return nil,err
		}
		return protocol.MarshalBinary(ret)
	})

	RegisterAdminFunc(LOAD_PNG_FOLDER_DATA, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		log.Info("LOAD_PNG_FOLDER_DATA")
		ret:=NewPngFolderData()
		path:=params.String()
		var err error
		util.WalkDirFilesWithFunc(path, func(filePath string, file os.FileInfo) bool {
			if util.IsFileWithExt(filePath,".png") {
				logger.Infof("png",filePath)
				var f *PngFile
				f,err =loadPngFile(filePath,logger)
				if err != nil {
					log.Error(err.Error())
					return true
				}
				ret.AddFile(f)
			}
			file.ModTime()
			return false
		},true)
		if err!=nil {
			return nil,err
		}
		return protocol.MarshalBinary(ret)
	})
}