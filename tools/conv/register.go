package conv

import (
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-lokas/util/convert"
	"strings"
)

type convertFunc func(string) (string,error)

var convertFuncs = map[string]convertFunc{}

func GetConvertTypePairString(t1,t2 convert.TYPE) string{
	return t1.ToString()+":"+t2.ToString()
}

func RegisterConvertFunc(t1,t2 convert.TYPE,f convertFunc){
	convertFuncs[GetConvertTypePairString(t1,t2)] = f
}

func GetConvertFunc(t1,t2 convert.TYPE)convertFunc{
	return convertFuncs[GetConvertTypePairString(t1,t2)]
}

func GetConvertAble(t convert.TYPE)protocol.IEnumCollection{
	var ret protocol.IEnumCollection = []protocol.IEnum{}
	for k,_:=range convertFuncs {
		split:=strings.Split(k,":")
		if split[0] == t.ToString() {
			ret  = append(ret, convert.ALL_ENC_TYPES.GetEnumByString(split[1]))
		}
	}
	return ret
}