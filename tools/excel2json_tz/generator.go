package excel2json_tz

import (
	"encoding/json"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/util"
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

func (this *Generator) Generate(p string,jsonPath string)error{
	this.dirSource = newDirSource(this.dirSource.dir)
	err:=this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	jsonMap:=map[string]map[string]*orderedmap.OrderedMap{}
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

type DataMap struct {
{DataFields}
	*DataMapImpl
}

func NewDataMap()*DataMap{
	ret:=&DataMap{
		DataMapImpl:&DataMapImpl{},
	}
	ret.Clear()
	ret.OnCreate()
	return ret
}

func (this *DataMap) Clear() {
	this.OnClear()
{InitFields}
}

func (this *DataMap) LoadJsonData(data []byte) error {
	err := json.Unmarshal(data, this)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.OnLoad()
	return nil
}

func (this *DataMap) LoadFromDb() error {
	return nil
}

func (this *DataMap) SaveToDb() error {
	return nil
}
`

	dataImplStr:=`package gamedata

type DataMapImpl struct {

}

func (this *DataMapImpl) OnCreate(){

}

func (this *DataMapImpl) OnClear(){

}

func (this *DataMapImpl) OnLoad(){
	
}`

	dataStr = strings.Replace(dataStr,"{InitFields}",initFieldStr,-1)
	dataStr = strings.Replace(dataStr,"{DataFields}",dataFieldStr,-1)
	jsonStr,err:=json.Marshal(jsonMap)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	jsonPath=path.Join(jsonPath,"data.json")
	err = ioutil.WriteFile(jsonPath, jsonStr, 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	dataPath:=path.Join(p,"data.go")
	err = ioutil.WriteFile(dataPath, []byte(dataStr), 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	implPath:=path.Join(p,"data_impl.go")
	log.Warnf("implPath",implPath,util.IsFileExist(implPath))
	if !util.IsFileExist(implPath) {
		err = ioutil.WriteFile(implPath, []byte(dataImplStr), 0644)
	}
	return nil
}

