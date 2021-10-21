package erpc

import (
	"errors"
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-tools/tools/dialog"
	"strings"
)

func init(){
	RegisterAdminFunc(OPEN_DIALOG_DIR, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		log.Info("OPEN_DIALOG_DIR")
		dir,err:=dialog.Directory().SetStartDir(params.String()).Title(params.String()).Browse()
		if err != nil {
			log.Error(err.Error())
			return nil,errors.New("cancelled")
		}
		return []byte(dir),nil

	})
	
	RegisterAdminFunc(OPEN_DIALOG_FILE, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		log.Info("OPEN_DIALOG_FILE")
		opStr:=params.String()
		var file string
		var err error
		startDir:=params.String()
		desc:=params.String()
		filterExtensionStr:=params.String()
		extensions:=strings.Split(filterExtensionStr,",")
		title:=params.String()

		builder:= dialog.File().SetStartDir(startDir).Filter(desc, extensions...).Title(title)
		if opStr == FILE_SAVE {
			file,err = builder.Save()
		} else if opStr == FILE_LOAD {
			file,err = builder.Load()
		} else {
			return nil,log.Error("op type not found:"+opStr)
		}
		if err != nil {
			log.Error(err.Error())
			return nil,errors.New("cancelled")
		}
		return []byte(file),nil
	})

	RegisterAdminFunc(OPEN_DIALOG_MSG, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		opstr:=params.String()
		builder:= dialog.Message("%s",params.String()).Title(params.String())
		var ok bool = true
		if opstr == MSG_ERROR {
			builder.Error()
		} else if opstr == MSG_INFO {
			builder.Info()
		} else if opstr == MSG_YESNO {
			ok = builder.YesNo()
		} else {
			return nil,log.Error("op type not found:"+opstr)
		}
		if !ok {
			return nil,errors.New("cancelled")
		}
		return nil,nil
	})
}