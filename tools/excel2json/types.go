package excel2json

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const MappingTag = "Mapping"


const (
	ERR_CELL_NOT_EXIST = "cell not exist"
)

const (
	EXCEL_TYPE_LINE  = 0
	EXCEL_DESC_LINE  = 1
	EXCEL_NAME_LINE  = 2
)

//type ExportType string
//
//const (
//	TypeAll    ExportType = "All"
//	TypeClient ExportType = "Client"
//	TypeServer ExportType = "Server"
//	TypeIgnore ExportType = "Ignore"
//)
//
//func (this ExportType) String()string {
//	return string(this)
//}
//
//func GetExportType(s string)(ExportType,error){
//	if s=="" {
//		return TypeAll,nil
//	}
//	switch s {
//	case TypeAll.String():
//		return TypeAll,nil
//	case TypeClient.String():
//		return TypeClient,nil
//	case TypeServer.String():
//		return TypeServer,nil
//	case TypeIgnore.String():
//		return TypeIgnore,nil
//	case strings.ToLower(TypeAll.String()):
//		return TypeAll,nil
//	case strings.ToLower(TypeClient.String()):
//		return TypeClient,nil
//	case strings.ToLower(TypeServer.String()):
//		return TypeServer,nil
//	case strings.ToLower(TypeIgnore.String()):
//		return TypeIgnore,nil
//	default:
//		return "",errors.New("type not exist:"+s)
//	}
//}



type fieldType int

const (
	type_tag = iota+1
	type_string
	type_bool
	type_int
	type_float
	type_int_arr
	type_float_arr
	type_string_arr
	type_string_int_map
	type_string_float_map
	type_string_string_map
	type_int_int_map
	type_int_float_map
	type_int_string_map
)

func IsCapitalWord(s string)bool{
	return regexp.MustCompile(`^[A-Z]\w+`).FindString(s) == s
}

func (this fieldType) TsString()string{
	switch this {
	case type_string:
		return "string"
	case type_bool:
		return "boolean"
	case type_int:
		return "number"
	case type_float:
		return "number"
	case type_int_arr:
		return "number[]"
	case type_float_arr:
		return "number[]"
	case type_string_arr:
		return "string[]"
	case type_string_int_map:
		return "{[key:string]:number}"
	case type_string_float_map:
		return "{[key:string]:number}"
	case type_string_string_map:
		return "{[key:string]:string}"
	case type_int_int_map:
		return "{[key:number]:number}"
	case type_int_float_map:
		return "{[key:number]:number}"
	case type_int_string_map:
		return "{[key:number]:string}"
	default:
		return ""
	}
}

func (this fieldType) check(s string)(string,error) {
	switch this {
	case type_int_arr:
		reg:=regexp.MustCompile(`(([0-9]+)([,]([0-9]+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check []int error")
		}
		return s,nil
	case type_float_arr:
		reg:=regexp.MustCompile(`((([0-9]|[.])+)([,](([0-9]|[.])+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check []float error")
		}
		return s,nil
	case type_string_arr:
		reg:=regexp.MustCompile(`((\w+)([,](\w+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check []string error")
		}
		return s,nil
	case type_string_string_map:
		reg:=regexp.MustCompile(`((\w+\s*[:]\s*\w+)([,](\w+\s*[:]\s*\w+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check [string]string error")
		}
		return s,nil
	case type_string_int_map:
		reg:=regexp.MustCompile(`((\w+\s*[:]\s*[0-9]+)([,](\w+\s*[:]\s*[0-9]+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check [string]int error")
		}
		return s,nil
	case type_string_float_map:
		reg:=regexp.MustCompile(`((\w+\s*[:]\s*([0-9]|[.])+)([,](\w+\s*[:]\s*([0-9]|[.])+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check [string]float error")
		}
		return s,nil
	case type_int_string_map:
		reg:=regexp.MustCompile(`(([0-9]+\s*[:]\s*\w+)([,]([0-9]+\s*[:]\s*\w+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check [int]string error")
		}
		return s,nil
	case type_int_int_map:
		reg:=regexp.MustCompile(`(([0-9]+\s*[:]\s*[0-9]+)([,]([0-9]+\s*[:]\s*[0-9]+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check [int]int error")
		}
		return s,nil
	case type_int_float_map:
		reg:=regexp.MustCompile(`(([0-9]+\s*[:]\s*([0-9]|[.])+)([,]([0-9]+\s*[:]\s*([0-9]|[.])+))*)?`)
		if reg.FindString(s)!=s {
			return "",errors.New("check [int]float error")
		}
		return s,nil
	default:
		return s,nil
	}
}

func trimLR(s string)[]string{
	s = strings.TrimRight(s,"]")
	s = strings.TrimLeft(s,"[")
	return strings.Split(s,",")
}

var errIgnore = errors.New("ignore")

func (this fieldType) decode(s string) (interface{},error){
	switch this {
	case type_string:
		return s,nil
	case type_bool:
		if s=="1" {
			return true,nil
		} else if s=="0" {
			return false,nil
		} else {
			return nil,errors.New("wrong bool format:"+s)
		}
	case type_int:
		if s=="#" {
			return 0,errIgnore
		}
		if s =="" {
			return 0,nil
		}
		ret,err:=strconv.Atoi(s)
		if err != nil {
			return nil,err
		}
		return ret,nil
	case type_float:
		if s =="" {
			return 0.0,nil
		}
		ret,err:=strconv.ParseFloat(s, 64)
		if err != nil {
			return nil,err
		}
		return ret,nil
	case type_int_arr:
		ret:= make([]float64,0)
		s ="["+s+"]"
		err:=json.Unmarshal([]byte(s),&ret)
		if err != nil {
			return nil,err
		}
		return ret,nil
	case type_float_arr:
		s ="["+s+"]"
		ret:= make([]float64,0)
		err:=json.Unmarshal([]byte(s),&ret)
		if err != nil {
			return nil,err
		}
		return ret,nil
	case type_string_arr:
		ret:=trimLR(s)
		return ret,nil
	case type_string_int_map:
		arr:=trimLR(s)
		ret := make(map[string]int)
		for _,iter:=range arr {
			iterArr:=strings.Split(iter,":")
			if len(iterArr)!=2 {
				return nil,errors.New("unmarshal error")
			}
			key:=iterArr[0]
			value,err:=strconv.Atoi(iterArr[1])
			if err != nil {
				return nil,err
			}
			ret[key] = value
		}
		return ret,nil
	case type_string_float_map:
		arr:=trimLR(s)
		ret := make(map[string]float64)
		for _,iter:=range arr {
			iterArr:=strings.Split(iter,":")
			if len(iterArr)!=2 {
				return nil,errors.New("unmarshal error")
			}
			key:=iterArr[0]
			value,err:=strconv.ParseFloat(iterArr[1], 64)
			if err != nil {
				return nil,err
			}
			ret[key] = value
		}
		return ret,nil
	case type_string_string_map:
		arr:=trimLR(s)
		ret := make(map[string]string)
		for _,iter:=range arr {
			iterArr:=strings.Split(iter,":")
			if len(iterArr)!=2 {
				return nil,errors.New("unmarshal error")
			}
			key:=iterArr[0]
			ret[key] = iterArr[1]
		}
		return ret,nil
	case type_int_int_map:
		arr:=trimLR(s)
		ret := make(map[int]int)
		for _,iter:=range arr {
			iterArr:=strings.Split(iter,":")
			if len(iterArr)!=2 {
				return nil,errors.New("unmarshal error")
			}
			key,err:=strconv.Atoi(iterArr[0])
			if err != nil {
				return nil,err
			}
			value,err:=strconv.Atoi(iterArr[1])
			if err != nil {
				return nil,err
			}
			ret[key] = value
		}
		return ret,nil
	case type_int_float_map:
		arr:=trimLR(s)
		ret := make(map[int]float64)
		for _,iter:=range arr {
			iterArr:=strings.Split(iter,":")
			if len(iterArr)!=2 {
				return nil,errors.New("unmarshal error")
			}
			key,err:=strconv.Atoi(iterArr[0])
			if err != nil {
				return nil,err
			}
			value,err:=strconv.ParseFloat(iterArr[1], 64)
			if err != nil {
				return nil,err
			}
			ret[key] = value
		}
		return ret,nil
	case type_int_string_map:
		arr:=trimLR(s)
		ret := make(map[int]string)
		for _,iter:=range arr {
			iterArr:=strings.Split(iter,":")
			if len(iterArr)!=2 {
				return nil,errors.New("unmarshal error")
			}
			key,err:=strconv.Atoi(iterArr[0])
			if err != nil {
				return nil,err
			}
			ret[key] = iterArr[1]
		}
		return ret,nil
	default:
		return 0,errors.New("unrecognized type:"+s)
	}
	return 0,errors.New("unrecognized type:"+s)
}

func getFieldType(s string)(fieldType,error) {
	switch s {
	case "#":
		return type_tag,nil
	case "string":
		return type_string,nil
	case "bool":
		return type_bool,nil
	case "int":
		return type_int,nil
	case "float":
		return type_float,nil
	case "[]int":
		return type_int_arr,nil
	case "[]float":
		return type_float_arr,nil
	case "[]string":
		return type_string_arr,nil
	case "[string]int":
		return type_string_int_map,nil
	case "[string]string":
		return type_string_string_map,nil
	case "[string]float":
		return type_string_float_map,nil
	default:
		return 0,errors.New("unrecognized type:"+s)
	}
}

func (this fieldType) GoString() string {
	switch this {
	case type_string:
		return "string"
	case type_bool:
		return "bool"
	case type_int:
		return "int32"
	case type_float:
		return "float64"
	case type_int_arr:
		return "[]int32"
	case type_float_arr:
		return "[]float64"
	case type_string_arr:
		return "[]string"
	case type_string_int_map:
		return "map[string]int32"
	case type_string_float_map:
		return "map[string]float64"
	case type_string_string_map:
		return "map[string]string"
	case type_int_int_map:
		return "map[int32]int32"
	case type_int_float_map:
		return "map[int32]float64"
	case type_int_string_map:
		return "map[int32]string"
	default:
		return ""
	}
}