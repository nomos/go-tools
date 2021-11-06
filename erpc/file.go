package erpc

import (
	"encoding/json"
	"errors"
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/rpc"
	"github.com/nomos/go-lokas/util"
	"io/ioutil"
)

func init(){
	rpc.RegisterAdminFunc(READ_FILE, func(cmd *lox.AdminCommand,params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		path:=cmd.ParamsValue().String()
		data,err:=ioutil.ReadFile(path)
		log.Warnf("path",path,string(data))
		if err != nil {
			log.Error(err.Error())
			return nil,err
		}
		return data,nil
	})
	rpc.RegisterAdminFunc(PATH_EXIST, func(cmd *lox.AdminCommand,params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		path:=cmd.ParamsValue().String()
		exist,_:=util.PathExists(path)
		if exist {
			return nil,nil
		} else {
			return nil,errors.New("not exist")
		}
	})
	rpc.RegisterAdminFunc(CREATE_FILE, func(cmd *lox.AdminCommand,params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		path:=params.String()
		log.Warnf("path",path)
		perm:=params.Int()
		log.Warnf("perm",perm)
		err:=util.CreateFile(path,perm)
		if err != nil {
			log.Error(err.Error())
			return nil,errors.New("create file failed")
		}
		return nil,nil
	})
	rpc.RegisterAdminFunc(WALK_DIR, func(cmd *lox.AdminCommand,params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		path:=params.String()
		log.Warnf("path",path)
		recursive:=params.Bool()
		log.Warnf("recursive",recursive)
		arr,err:=util.WalkDir(path,recursive)
		if err != nil {
			log.Error(err.Error())
			return nil,errors.New("walk dir failed")
		}
		data,_:=json.Marshal(arr)
		log.Warnf("data",string(data))
		return data,nil
	})
	rpc.RegisterAdminFunc(EXEC_PATH, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		p,_:=util.ExecPath()
		return []byte(p),nil
	})

}
