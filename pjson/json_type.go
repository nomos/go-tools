package pjson

import (
	"github.com/nomos/go-lokas/util/log"
	"github.com/ying32/govcl/vcl"
	"strconv"
	"strings"
)

type Type int

const (
	Null Type = iota
	Object
	Array
	String
	Number
	Boolean
)

var default_map = make(map[Type]string)
var name_map = make(map[Type]string)

func init(){
	default_map[Null] = "null"
	default_map[Object] = ""
	default_map[Array] = ""
	default_map[String] = ""
	default_map[Number] = "NaN"
	default_map[Boolean] = "false"
	name_map[Null] = "Null"
	name_map[Object] = "Object"
	name_map[Array] = "Array"
	name_map[String] = "String"
	name_map[Number] = "Number"
	name_map[Boolean] = "Boolean"
}

func (this Type) String()string{
	ret,ok:=name_map[this]
	if ok {
		return ret
	}
	return "Unknown"
}

func (this Type) Default()string {
	if this == Object||this == Array {
		log.Panic("call default string with "+this.String())
	}
	return default_map[this]
}

func(this Type) CheckValue(s string)(string,bool) {
	if this == Object||this == Array {
		log.Panic("check value with "+this.String())
	}vcl.NewPopupMenu(mainForm)
	switch this {
	case Null:
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		if s== "null" {
			return "null",true
		}
		return "",false
	case String:
		return s,true
	case Number:
		s = strings.TrimSpace(s)
		if s == "NaN"||s == "nan" {
			return "NaN",true
		}
		_,err:=strconv.Atoi(s)
		if err!= nil {
			return "",false
		}
		return s,true
	case Boolean:
		s = strings.TrimSpace(s)
		if s=="True"||s=="TRUE"||s=="true" {
			return "true",true
		}
		if s=="False"||s=="FALSE"||s=="false" {
			return "false",true
		}
		return "",false
	default:
		log.Panic("unrecognized type")
	}
	return "",false
}