package excel2json

import (
	"encoding/json"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/stringutil"
	"io/ioutil"
	"path"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
)

func New(source string) *Generator {
	ret := &Generator{
		dirSource: newDirSource(source),
	}
	return ret
}

type Generator struct {
	*dirSource
	CsPackage string
}

func (this *Generator) Load() error {
	defer func() {
		if err, ok := recover().(error); ok {
			log.Error(err.Error())
			debug.PrintStack()
		}
	}()
	err := this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (this *Generator) GenerateJsonString() (string, error) {
	this.dirSource = newDirSource(this.dirSource.dir)
	err := this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	jsonMap := map[string]map[string]*orderedmap.OrderedMap{}
	for _, f := range this.dirSource.files {
		for _, s := range f.sheets {
			jsonMap[s.Name] = s.GetJsonMap()
		}
	}
	jsonStr, err := json.Marshal(jsonMap)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	return string(jsonStr), nil
}

func (this *Generator) GenerateJson(jsonPath string) error {
	jsonStr, err := this.GenerateJsonString()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	jsonPath = path.Join(jsonPath, "data.json")
	log.Warnf(jsonPath)
	err = ioutil.WriteFile(jsonPath, []byte(jsonStr), 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func (this *Generator) Generate(gopath string, jsonPath string) error {
	err := this.GenerateJson(jsonPath)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err = this.GenerateGo(gopath)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (this *Generator) GenerateTs(tsPath string, embed bool) error {
	this.dirSource = newDirSource(this.dirSource.dir)
	err := this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	sheetArr := []*SheetSource{}
	importFieldStr := ""
	sourceFieldStr := ""
	loadFieldStr := ""
	dirName := path.Base(tsPath)
	log.Infof("baseName", dirName)
	for _, f := range this.dirSource.files {
		for _, s := range f.sheets {
			sheetArr = append(sheetArr, s)
			tsFilePath := path.Join(tsPath, dirName+"_"+stringutil.SplitCamelCaseLowerSnake(s.Name)) + "_source.ts"
			log.Info(tsFilePath)
			err := ioutil.WriteFile(tsFilePath, []byte(s.GenerateTsString()), 0644)
			if err != nil {
				log.Errorf(err.Error())
			}
		}
	}
	sort.Slice(sheetArr, func(i, j int) bool {
		return sheetArr[i].Name < sheetArr[j].Name
	})
	for _, s := range sheetArr {
		sourceFieldStr += s.GetTsDataFieldString()
		sourceFieldStr += "\n"
		importFieldStr += s.GetTsImportFieldString(dirName)
		importFieldStr += "\n"
		loadFieldStr += s.GetTsLoadFieldString()
		loadFieldStr += "\n"
	}
	importFieldStr = strings.TrimRight(importFieldStr, "\n")
	sourceFieldStr = strings.TrimRight(sourceFieldStr, "\n")
	loadFieldStr = strings.TrimRight(loadFieldStr, "\n")
	dataStr := `{ImportFields}
{DataFields}

class _DataSource {
{SourceFields}
    constructor(objs:Dict<NumericDict<any>>|string) {
		this.load(objs)
    }
    load(objs:Dict<NumericDict<any>>|string){
		if (typeof objs=="string") {
			objs = JSON.parse(objs)
		}
		if (!objs) {
			return;
		}
{LoadFields}
    }
}

export const DataSource = new _DataSource(json)`

	dataStr = strings.Replace(dataStr, "{ImportFields}", importFieldStr, -1)
	dataStr = strings.Replace(dataStr, "{SourceFields}", sourceFieldStr, -1)
	dataStr = strings.Replace(dataStr, "{LoadFields}", loadFieldStr, -1)
	if embed {
		var jsonStr string
		jsonStr, err = this.GenerateJsonString()
		if err != nil {
			log.Error(err.Error())
			return err
		}
		dataStr = strings.Replace(dataStr, "{DataFields}", "\nconst json = "+jsonStr, -1)
	} else {
		dataStr = strings.Replace(dataStr, "{DataFields}", `const json = null`, -1)
	}
	baseName := path.Base(tsPath)
	dataPath := path.Join(tsPath, baseName+"_data_source.ts")
	err = ioutil.WriteFile(dataPath, []byte(dataStr), 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func (this *Generator) GenerateGo(gopath string) error {
	this.dirSource = newDirSource(this.dirSource.dir)
	err := this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	sheetArr := []*SheetSource{}
	initFieldStr := ""
	dataFieldStr := ""
	for _, f := range this.dirSource.files {
		for _, s := range f.sheets {
			sheetArr = append(sheetArr, s)
			goFilePath := path.Join(gopath, stringutil.SplitCamelCaseLowerSnake(s.Name)) + ".go"
			log.Warnf(goFilePath)
			err := ioutil.WriteFile(goFilePath, []byte(s.GenerateGoString(gopath)), 0644)
			if err != nil {
				log.Errorf(err.Error())
			}
			if s.Name == "PropEnum" {
				generatePropEnum(path.Join(gopath, "enum_prop.go"), s.Data)
			}
		}
	}
	sort.Slice(sheetArr, func(i, j int) bool {
		return sheetArr[i].Name < sheetArr[j].Name
	})
	for _, s := range sheetArr {
		dataFieldStr += s.GetGoDataFieldString()
		dataFieldStr += "\n"
		initFieldStr += s.GetGoInitFieldString()
		initFieldStr += "\n"
	}
	initFieldStr = strings.TrimRight(initFieldStr, "\n")
	dataFieldStr = strings.TrimRight(dataFieldStr, "\n")
	dataStr := `package {Package}

type DataMap struct {
{DataFields}
}

func (this *DataMap) Clear() {
{InitFields}
}`
	dataStr = strings.Replace(dataStr, "{Package}", path.Base(gopath), -1)
	dataStr = strings.Replace(dataStr, "{InitFields}", initFieldStr, -1)
	dataStr = strings.Replace(dataStr, "{DataFields}", dataFieldStr, -1)
	dataPath := path.Join(gopath, "data.go")
	err = ioutil.WriteFile(dataPath, []byte(dataStr), 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func (this *Generator) GenerateCs(csPath string, namespace string) error {
	this.dirSource = newDirSource(this.dirSource.dir)
	this.CsPackage = namespace
	err := this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	sheetArr := []*SheetSource{}
	sourceFieldStr := ""
	dirName := path.Base(csPath)
	log.Infof("baseName", dirName)
	for _, f := range this.dirSource.files {
		for _, s := range f.sheets {
			sheetArr = append(sheetArr, s)
			csFilePath := path.Join(csPath, s.Name+"Source.cs")
			log.Info(csFilePath)
			err := ioutil.WriteFile(csFilePath, []byte(s.GenerateCsString(this)), 0644)
			if err != nil {
				log.Errorf(err.Error())
			}
		}
	}
	sort.Slice(sheetArr, func(i, j int) bool {
		return sheetArr[i].Name < sheetArr[j].Name
	})
	for _, s := range sheetArr {
		sourceFieldStr += s.GetCsDataFieldString()
		sourceFieldStr += "\n"
	}
	sourceFieldStr = strings.TrimRight(sourceFieldStr, "\n")
	dataStr := `using System.Collections.Generic;
using Newtonsoft.Json;

namespace {NameSpace} {
	[JsonObject(MemberSerialization.OptIn)]
	public class DataSource {
		protected static DataSource _instance;
{SourceFields}
    	protected DataSource()
        {
	        
        }
		
        public static DataSource Instance(string data = "")
        {
	        if (_instance == null)
	        {
		        _instance = new DataSource();
		        if (data != "")
		        {
			        Load(data);
		        }
	        }
	        return _instance;
        }
    	public static DataSource Load(string data)
    	{
        	var ret = Instance();
            JsonConvert.PopulateObject(data,ret);
        	return ret;
    	}
	}
}

`

	dataStr = strings.Replace(dataStr, "{NameSpace}", namespace, -1)
	dataStr = strings.Replace(dataStr, "{SourceFields}", sourceFieldStr, -1)
	dataPath := path.Join(csPath, "DataSource.cs")
	err = ioutil.WriteFile(dataPath, []byte(dataStr), 0644)
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func generatePropEnum(p string, lines []*DataLine) {
	str := `package {Package}

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
	lineStr := ""
	str2enum := ""
	enum2str := ""
	allenums := ""
	for _, l := range lines {
		id := l.line[0].Value.(int)
		name := l.line[2].Value.(string)
		desc := l.line[1].Value.(string)
		name = strings.ReplaceAll(name, "%", "Percent")
		lineStr += "\t" + "PROP_" + stringutil.SplitCamelCaseUpperSnake(name) + " PROP_ENUM = " + strconv.Itoa(id) + " //" + desc + "\n"
		str2enum += "\t" + "case " + `"` + desc + `":` + "\n" + "\t\t" + "return " + "PROP_" + stringutil.SplitCamelCaseUpperSnake(name) + "\n"
		enum2str += "\t" + "case " + "PROP_" + stringutil.SplitCamelCaseUpperSnake(name) + `:` + "\n" + "\t\t" + "return " + `"` + desc + `"` + "\n"
		allenums += "PROP_" + stringutil.SplitCamelCaseUpperSnake(name) + ","
	}
	lineStr = strings.TrimRight(lineStr, "\n")
	str2enum = strings.TrimRight(str2enum, "\n")
	enum2str = strings.TrimRight(enum2str, "\n")
	allenums = strings.TrimRight(allenums, ",")
	str = strings.ReplaceAll(str, "{Package}", path.Base(p))
	str = strings.ReplaceAll(str, "{enumline}", lineStr)
	str = strings.ReplaceAll(str, "{str2enum}", str2enum)
	str = strings.ReplaceAll(str, "{enum2str}", enum2str)
	str = strings.ReplaceAll(str, "{allenums}", allenums)
	ioutil.WriteFile(p, []byte(str), 0644)
}
