package erpc

import (
	"errors"
	"github.com/gen2brain/dlgs"
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/util"
	"github.com/sqweek/dialog"
)

func init(){
	registerFunc(OPEN_DIALOG_DIR, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		defer func() {
			if r := recover(); r != nil {
				util.Recover(r,false)
			}
		}()
		dir,success,err:=dlgs.File(params.String(),"",true)
		if err != nil {
			log.Error(err.Error())
			return nil,log.Error("cancelled")
		}
		if !success {
			return nil,log.Error("cancelled")
		}
		return []byte(dir),nil

	})
	
	registerFunc(OPEN_DIALOG_FILE, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		defer func() {
			if r := recover(); r != nil {
				util.Recover(r,false)
			}
		}()
		dir,success,err:=dlgs.File(params.String(),params.String(),false)
		if err != nil {
			log.Error(err.Error())
			return nil,log.Error("cancelled")
		}
		if !success {
			return nil,log.Error("cancelled")
		}
		return []byte(dir),nil
	})

	registerFunc(OPEN_DIALOG_MSG, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		defer func() {
			if r := recover(); r != nil {
				util.Recover(r,false)
			}
		}()
		log.Warnf("OPEN_DIALOG_MSG")
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