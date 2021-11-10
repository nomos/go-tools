package excel2json

import (
	"encoding/json"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/stringutil"
	"io/ioutil"
	"path"
	"regexp"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
)

func New(source string)*Generator {
	ret:=&Generator{
		dirSource:newDirSource(source),
	}
	return ret
}

type Generator struct {
	*dirSource
}

func (this *Generator) Load()error {
	defer func() {
		if err,ok:=recover().(error);ok {
			log.Error(err.Error())
			debug.PrintStack()
		}
	}()
	err:=this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (this *Generator) GenerateJson(jsonPath string)error{
	this.dirSource = newDirSource(this.dirSource.dir)
	err:=this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	jsonMap:=map[string]map[string]*orderedmap.OrderedMap{}
	for _,f:=range this.dirSource.files {
		for _,s:=range f.sheets {
			jsonMap[s.Name] = s.GetJsonMap()
		}
	}
	jsonStr,err:=json.Marshal(jsonMap)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	jsonPath=path.Join(jsonPath,"data.json")
	log.Warnf(jsonPath)
	err = ioutil.WriteFile(jsonPath, jsonStr, 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func (this *Generator) Generate(gopath string,jsonPath string)error{
	err:=this.GenerateJson(jsonPath)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err=this.GenerateGo(gopath)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (this *Generator) GenerateTs(tsPath string) error{
	return nil
}

func (this *Generator) GenerateGo(gopath string)error{
	this.dirSource = newDirSource(this.dirSource.dir)
	err:=this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	sheetArr:=[]*SheetSource{}
	initFieldStr:=""
	dataFieldStr:=""
	for _,f:=range this.dirSource.files {
		for _,s:=range f.sheets {
			sheetArr = append(sheetArr, s)
			goFilePath:=path.Join(gopath,stringutil.SplitCamelCaseLowerSnake(s.Name))+".go"
			log.Warnf(goFilePath)
			err := ioutil.WriteFile(goFilePath, []byte(s.GenerateGoString()), 0644)
			if err != nil {
				log.Errorf(err.Error())
			}
			if s.Name == "PropEnum"{
				generatePropEnum(path.Join(gopath,"enum_prop.go"),s.Data)
			}
		}
	}
	sort.Slice(sheetArr, func(i, j int) bool {
		return sheetArr[i].Name<sheetArr[j].Name
	})
	for _,s:=range sheetArr {
		dataFieldStr+=s.GetDataFieldString()
		dataFieldStr+="\n"
		initFieldStr+=s.GetInitFieldString()
		initFieldStr+="\n"
	}
	initFieldStr = strings.TrimRight(initFieldStr,"\n")
	dataFieldStr = strings.TrimRight(dataFieldStr,"\n")
	dataStr := `package gamedata

type DataMap struct {
{DataFields}
}

func (this *DataMap) Clear() {
{InitFields}
}`

	dataStr = strings.Replace(dataStr,"{InitFields}",initFieldStr,-1)
	dataStr = strings.Replace(dataStr,"{DataFields}",dataFieldStr,-1)
	dataStr = regexp.MustCompile(`[*]Chap.+Normal`).ReplaceAllString(dataStr,"*ChapterNormal")
	dataPath:=path.Join(gopath,"data.go")
	err = ioutil.WriteFile(dataPath, []byte(dataStr), 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func generatePropEnum(p string,lines []*DataLine){
	str:=`package gamedata

import "github.com/nomos/go-lokas/protocol"

type PROP_ENUM protocol.Enum

const (
{enumline}
)

func (this PROP_ENUM) Int32()int32{
	return int32(this)
}

var ALL_PROP_ENUM = []protocol.IEnum{{allenums}}

func TO_PROP_ENUM(s string)PROP_ENUM{
	switch s {
{str2enum}
	}
	return -1
}

func (this PROP_ENUM) Enum()protocol.Enum{
	return protocol.Enum(this)
}


func (this PROP_ENUM) ToString()string{
	switch this {
{enum2str}
	}
	return ""
}

`
	lineStr:=""
	str2enum:=""
	enum2str:=""
	allenums:=""
	for _,l:=range lines {
		id:=l.line[0].Value.(int)
		name:=l.line[2].Value.(string)
		desc:=l.line[1].Value.(string)
		name = strings.ReplaceAll(name,"%","Percent")
		lineStr+="\t"+"PROP_"+stringutil.SplitCamelCaseUpperSnake(name)+" PROP_ENUM = "+strconv.Itoa(id)+" //"+desc+"\n"
		str2enum+="\t"+"case "+`"`+desc+`":`+"\n"+"\t\t"+"return "+"PROP_"+stringutil.SplitCamelCaseUpperSnake(name)+"\n"
		enum2str+="\t"+"case "+"PROP_"+stringutil.SplitCamelCaseUpperSnake(name)+`:`+"\n"+"\t\t"+"return "+`"`+desc+`"`+"\n"
		allenums+="PROP_"+stringutil.SplitCamelCaseUpperSnake(name)+","
	}
	lineStr = strings.TrimRight(lineStr,"\n")
	str2enum = strings.TrimRight(str2enum,"\n")
	enum2str = strings.TrimRight(enum2str,"\n")
	allenums = strings.TrimRight(allenums,",")
	str = strings.ReplaceAll(str,"{enumline}",lineStr)
	str = strings.ReplaceAll(str,"{str2enum}",str2enum)
	str = strings.ReplaceAll(str,"{enum2str}",enum2str)
	str = strings.ReplaceAll(str,"{allenums}",allenums)
	ioutil.WriteFile(p,[]byte(str),0644)
}
