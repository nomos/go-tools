package erpc

import (
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/network/shell"
	"github.com/nomos/go-lokas/network/sshc"
	"github.com/nomos/go-lokas/rpc"
)

func init(){
	rpc.RegisterAdminFunc(EXEC_CMD, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		err:=execCmd(params,logger)
		return nil,err
	})
	rpc.RegisterAdminFunc(SCP, func(cmd *lox.AdminCommand, params *cmds.ParamsValue, logger log.ILogger) ([]byte, error) {
		addr:=params.String()
		pass:=params.String()
		localPath:=params.String()
		remotePath:=params.String()
		s:=sshc.NewSshClient("root",pass,addr+":22",false)
		_,err:=s.Connect().Await()
		if err != nil {
			logger.Error(err.Error())
			return nil,err
		}
		s.SetConsoleWriter(logger)
		err=s.Upload(localPath,remotePath)
		if err != nil {
			logger.Error(err.Error())
			return nil,err
		}
		return nil,nil
	})
}

func execCmd(params *cmds.ParamsValue, logger log.ILogger)error{
	s:=shell.New(true,mergeCmd(params),false)
	err:=s.Start()
	s.SetWriter(log.DefaultLogger())
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = s.Wait()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return err
}

func mergeCmd(params *cmds.ParamsValue)string{
	cmdStr:=""
	for {
		str:=params.StringOpt()
		if str=="" {
			break
		}
		cmdStr+=str
		cmdStr+=" "
	}
	return cmdStr
}