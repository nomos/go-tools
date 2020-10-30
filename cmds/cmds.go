package cmds

import (
	"github.com/nomos/go-log/log"
	"strconv"
	"strings"
)

type CmdType string

const (
	Cmd_Shell = "shell"
	Cmd_Expect = "expect"
)

func (this CmdType) GetCmdPrefix()string {
	switch this {
	case Cmd_Shell:
		return "/bin/sh"
	case Cmd_Expect:
		return "/usr/bin/expect"
	}
	return ""
}

type CmdResult struct {
	Outputs []string
	Success bool
	Results map[string]interface{}
}

type CmdOutput []string

func (this CmdOutput) LastOutput()string{
	if len(this)==0 {
		return ""
	}
	return this[len(this)-1]
}

type CmdHandler func(CmdOutput)*CmdResult

type WrappedCmd struct {
	CmdString string
	Tips string
	ParamsNum int
	ParamsMap []string
	CmdType CmdType
	CmdHandler CmdHandler
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