package erpc

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"io/ioutil"
)

func init(){
	registerFunc(READ_FILE, func(cmd *lox.AdminCommand) ([]byte, error) {
		path:=cmd.ParamsValue().String()
		data,err:=ioutil.ReadFile(path)
		log.Warnf("path",path,string(data))
		if err != nil {
			log.Error(err.Error())
			return nil,err
		}
		return data,nil
	})
	registerFunc("test", func(cmd *lox.AdminCommand) ([]byte, error) {
		log.Warn("test")
		return nil,nil
	})
}
