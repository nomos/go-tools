package pjson

import (
	"github.com/nomos/go-lokas/util/log"
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
	default_map[Number] = "0"
	default_map[Boolean] = "false"
	name_map[Null] = "Null"
	name_map[Object] = "Object"
	name_map[Array] = "Array"
	name_map[String] = "String"
	name_map[Number] = "Number"
	name_map[Boolean] = "Boolean"
}
func (this Type) IsValue()bool{
	return this == Number||this==String||this==Boolean||this==Null
}

func GetTypeByString(s string)Type {
	for k,v:=range name_map {
		if v == s {
			return k
		}
	}
	log.Panic("unknown type")
	return -1
}

func (this Type) CreateDefaultSchema()*Schema {
	ret:=NewSchema()
	ret.Type = this
	ret.Value = this.Default()
	return ret
}

func (this Type) String()string{
	ret,ok:=name_map[this]
	if ok {
		return ret
	}
	return "Unknown"
}

func (this Type) Default()string {
	return default_map[this]
}

func(this Type) CheckValue(s string)(string,bool) {
	if this == Object||this == Array {
		log.Panic("check value with "+this.String())
	}
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
		_,err:=strconv.ParseFloat(s,64)
		if err!= nil {
			return "",false
		}
		return s,true
	case Boolean:
		s = strings.TrimSpace(s)
		log.Warnf("check bool",s)
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