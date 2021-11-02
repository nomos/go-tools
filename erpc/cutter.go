package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-tools/tools/pics/img_png"
)

func init()  {
	RegisterAdminFunc(CUT_PNG, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		p:=params.String()
		width:=params.Int()
		height:=params.Int()
		err:=img_png.SubImage(p,width,height,logger)
		if err != nil {
			log.Error(err.Error())
		}
		return nil,nil
	})
}