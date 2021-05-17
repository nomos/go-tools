package excel2json_tz

import (
	"encoding/json"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/util/stringutil"
	"io/ioutil"
	"path"
	"runtime/debug"
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



func (this *Generator) Generate(p string)error{
	jsonMap:=map[string][]*orderedmap.OrderedMap{}
	sheetArr:=[]*SheetSource{}
	initFieldStr:=""
	dataFieldStr:=""

	for _,f:=range this.dirSource.files {
		for _,s:=range f.sheets {
			sheetArr = append(sheetArr, s)
			initFieldStr+=s.GetInitFieldString()
			initFieldStr+="\n"
			dataFieldStr+=s.GetDataFieldString()
			dataFieldStr+="\n"
			goFilePath:=path.Join(p,stringutil.SplitCamelCaseLowerSlash(s.Name))+".go"
			log.Warnf(goFilePath)
			err := ioutil.WriteFile(goFilePath, []byte(s.GenerateGoString()), 0644)
			if err != nil {
				log.Errorf(err.Error())
			}
			jsonMap[s.Name] = s.GetJsonMap()
		}
	}
	initFieldStr = strings.TrimRight(initFieldStr,"\n")
	dataFieldStr = strings.TrimRight(dataFieldStr,"\n")
	dataStr := `package gamedata

import (
	"encoding/json"
	"github.com/nomos/go-log/log"
)

var DataMap = &dataMap{
{InitFields}
}

type dataMap struct {
{DataFields}
}

func (this *dataMap) LoadJsonData(data string)error{
	err:=json.Unmarshal([]byte(data),this)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}`
	dataStr = strings.Replace(dataStr,"{InitFields}",initFieldStr,-1)
	dataStr = strings.Replace(dataStr,"{DataFields}",dataFieldStr,-1)
	jsonStr,err:=json.Marshal(jsonMap)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	jsonPath:=path.Join(p,"data.json")
	err = ioutil.WriteFile(jsonPath, jsonStr, 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	dataPath:=path.Join(p,"data.go")
	err = ioutil.WriteFile(dataPath, []byte(dataStr), 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}


