package cmds

import (
	"github.com/nomos/go-log/log"
	"strconv"
	"strings"
)

type WrappedCmd struct {
	CmdString string
	Tips string
	ParamsNum int
	ParamsMap []string
}

func (this *WrappedCmd) FillParams(param... string)string{
	ret :=this.CmdString
	for i,s:=range param{
		ret = strings.Replace(ret,"$"+strconv.Itoa(i+1),s,1)
	}
	return ret
}

var wrappedCmds = make(map[string]*WrappedCmd)

func GetAllCmds()map[string]*WrappedCmd {
	return wrappedCmds
}

func RegisterCmd(s string,cmd *WrappedCmd){
	if _,ok:=wrappedCmds[s];ok{
		log.Warnf("duplicated cmd:%s,overwrite",s)
	}
	wrappedCmds[s] = cmd
}

func GetCmdByName(s string)*WrappedCmd {
	return wrappedCmds[s]
}