package vcltool

import (
	"errors"
	"fmt"
	"github.com/nomos/go-promise"
	"strconv"
	"strings"
)

type ICommandSender interface {
	SendCmd(string)
	OnSelect()
	OnDeselect()
}

type ICommand interface {
	Name()string
	Exec(param *ParamsValue,console *TConsoleShell)*promise.Promise
	Tips()string
}

func SplitCommand(cmd string)(string,[]string){
	splits := strings.Split(cmd," ")
	if len(splits)==0 {
		return "",[]string{}
	}
	return splits[0],splits[1:]
}

type Command struct {
	name string
	exec func(value *ParamsValue,console *TConsoleShell)*promise.Promise
	tips string
}

func (this *Command) Name()string{
	return this.name
}

func (this *Command) Exec(param *ParamsValue,console *TConsoleShell)*promise.Promise {
	if this.exec!=nil {
		return this.exec(param,console)
	}
	return promise.Reject(errors.New("cant found exec"))
}

func (this *Command) Tips()string{
	return this.tips
}

func NewCommand(name string,tips string,f func(value *ParamsValue,console *TConsoleShell)*promise.Promise)ICommand{
	ret:=&Command{
		name: name,
		exec: f,
		tips:tips,
	}
	return ret
}

type ParamsValue struct {
	cmd string
	value []string
	offset int
}

const (
	CMD_ERROR_PARAM_LEN = iota
	CMD_ERROR_PARAM_TYPE
)

type CmdError struct {
	errorType int
	cmd string
	offset int
	paramType string
}

func (this *CmdError) Error() string {
	switch this.errorType {
	case CMD_ERROR_PARAM_LEN:
		return fmt.Sprintf("cmd %s params length error, type %s help|? ",this.cmd,this.cmd)
	case CMD_ERROR_PARAM_TYPE:
		return fmt.Sprintf("cmd %s params[%d] type error,must be %s,type %s help|?",this.cmd,this.offset,this.paramType,this.cmd)
	default:
		return fmt.Sprintf("CmdError Type Error:%d",this.errorType)
	}
}

func NewCmdError(errType int,cmd string,offset int,typ string)*CmdError {
	return &CmdError{
		errorType: errType,
		cmd:cmd,
		offset:offset,
		paramType: typ,
	}
}

func (this *ParamsValue) StringOpt()string {
	if len(this.value)-1<this.offset {
		return ""
	}
	ret:=this.value[this.offset]
	this.offset++
	return ret
}

func (this *ParamsValue) String()string {
	if len(this.value)-1<this.offset {
		panic(NewCmdError(CMD_ERROR_PARAM_LEN,this.cmd,this.offset,"string"))
	}
	ret:=this.value[this.offset]
	this.offset++
	return ret
}

func (this *ParamsValue) Int()int {
	if len(this.value)-1<this.offset {
		panic(NewCmdError(CMD_ERROR_PARAM_LEN,this.cmd,this.offset,"int"))
	}
	ret,err:=strconv.Atoi(this.value[this.offset])
	if err != nil {
		panic(NewCmdError(CMD_ERROR_PARAM_TYPE,this.cmd,this.offset,"int"))
	}
	this.offset++
	return ret
}


func (this *ParamsValue) IntOpt()int {
	if len(this.value)-1<this.offset {
		return 0
	}
	ret,err:=strconv.Atoi(this.value[this.offset])
	if err != nil {
		panic(NewCmdError(CMD_ERROR_PARAM_TYPE,this.cmd,this.offset,"int"))
	}
	this.offset++
	return ret
}

func (this *ParamsValue) Float()float64 {
	if len(this.value)-1<this.offset {
		panic(NewCmdError(CMD_ERROR_PARAM_LEN,this.cmd,this.offset,"float"))
	}
	ret,err:=strconv.ParseFloat(this.value[this.offset],3)
	if err != nil {
		panic(NewCmdError(CMD_ERROR_PARAM_TYPE,this.cmd,this.offset,"float"))
	}
	this.offset++
	return ret
}

func (this *ParamsValue) FloatOpt()float64 {
	if len(this.value)-1<this.offset {
		return 0
	}
	ret,err:=strconv.ParseFloat(this.value[this.offset],3)
	if err != nil {
		panic(NewCmdError(CMD_ERROR_PARAM_TYPE,this.cmd,this.offset,"float"))
	}
	this.offset++
	return ret
}
