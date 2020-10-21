package cmds

import (
	"strconv"
	"strings"
)

type WrappedCmd string

func (this WrappedCmd) FillParams(param... string)string{
	ret :=string(this)
	for i,s:=range param{
		ret = strings.Replace(ret,"$"+strconv.Itoa(i+1),s,1)
	}
	return ret
}