package excel2json

import (
	"encoding/json"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-lokas/util/slice"
	"github.com/nomos/go-lokas/util/stringutil"
	"io/ioutil"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type gameDataField struct {
	fieldIndex int
	fieldName string
	fieldDesc string
	fieldType fieldType
}

func (this *gameDataField) string(data gameData)string {

	return ""
}

type gameData map[string]interface{}

//对应一个文件
type gameFileSource struct {
	fields []*gameDataField
	enums map[string]int
	data []gameData
}

func (this *gameFileSource) fieldString()string {
	ret:=""
	for _,field:=range this.fields {
		if field.fieldName == "#" {
			continue
		}
		ret+="\t"+field.fieldName+":"+field.fieldType.TsString()+"\n"
	}
	ret = strings.TrimRight(ret,"\n")
	return ret
}

func (this *gameFileSource) isEmptyLine(data []string)bool {
	for _,s:=range data {
		if s=="" {
			continue
		}
		return false
	}
	return true
}

func (this *gameFileSource) parseData(in []string)error{
	if this.isEmptyLine(in) {
		return nil
	}
	data:=make(gameData)
	hasEnum:=false
	var enumField *gameDataField
	for _,field:=range this.fields {
		if field.fieldName=="#" {
			hasEnum=true
			enumField = field
			break
		}
	}
	ignored:=false
	for _,field:=range this.fields {
		if field.fieldName=="#" {
			continue
		}
		str:=""
		if len(in)<=field.fieldIndex {
			log.Warn("out of range")
		} else {
			str=in[field.fieldIndex]
		}
		t:=field.fieldType
		_,err:=t.check(str)
		if err != nil {
			colName,_:=excelize.ColumnNumberToName(field.fieldIndex)
			return errors.New("列:"+colName+" 格式检查错误:"+err.Error())
		}
		d,err:=t.decode(str)
		if field.fieldName=="id"&&err==errIgnore {
			ignored = true
			continue
		}
		if err != nil {
			colName,_:=excelize.ColumnNumberToName(field.fieldIndex)
			return errors.New("列:"+colName+" 反序列化错误:"+err.Error())
		}
		if field.fieldName == "id"&&hasEnum {
			if len(in)>field.fieldIndex {
				enum:=in[enumField.fieldIndex]
				if enum!="" {
					this.enums[enum] = d.(int)
				}
			}
		}
		data[field.fieldName] = d
	}
	if ignored {
		return nil
	}
	this.data = append(this.data, data)
	return nil
}

//对应一个目录下所有文件
type gameDataSource struct {
	files map[string]*gameFileSource
}

type excel2JsonMiniGame struct {
	embedJson bool
	logger log.ILogger
}

func (this *excel2JsonMiniGame)generateData(source *gameDataSource,p string)error{
	for key,data:=range source.files{
		log.Warnf("key",key)
		err:=this.generateTsSchema(data,key,p)
		if err != nil {
			return err
		}
		if !this.embedJson {
			err=this.generateJsonData(data,key,p)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func Excel2JsonMiniGame(excelPath,distPath string,logger log.ILogger,embed bool)(err error){
	ret := &excel2JsonMiniGame{}
	ret.embedJson = embed
	if logger!=nil {
		ret.logger = logger
	} else {
		ret.logger = log.DefaultLogger()
	}
	defer func() {
		if r := recover(); r != nil {
			err = log.Error(r.(error).Error())
		}
	}()
	err=ret.generate(excelPath,distPath)
	if err != nil {
		return log.Error(err.Error())
	}
	return
}

func  (this *excel2JsonMiniGame)generate(excelPath,distPath string)error{
	paths,err := util.WalkDirFiles(excelPath,false)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	paths = util.FilterFileWithExt(paths,".xlsx",".xlsm")
	paths = util.FilterFileWithFunc(paths, func(s string) bool {
		return !regexp.MustCompile(`[~][$]`).MatchString(s)
	})
	this.logger.Infof("当前路径",paths)
	var source = &gameDataSource{files: map[string]*gameFileSource{}}
	for _,p:=range paths {
		this.logger.Warn("开始读取"+p)
		err=this.fetchGameDataSource(source,p)
		if err != nil {
			return err
		}
	}
	this.logger.Warn("读取目录成功:"+excelPath)
	err=this.generateData(source,distPath)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	this.logger.Warn("-------ALL DONE-------")
	return nil
}

func (this *excel2JsonMiniGame)checkClassName(s string)bool{
	if regexp.MustCompile(`Sheet[0-9]*`).FindString(s) == s {
		return false
	}
	return IsCapitalWord(s)
}

func (this *excel2JsonMiniGame) parseGameFields (key string,data [][]string) (map[int]*gameDataField,error) {
	ret:=make(map[int]*gameDataField)
	hasId:=false
	for index,col:=range data {
		if col[0] == "" {
			continue
		}
		type_str:=col[0]
		type_desc:=col[1]
		type_name:=col[2]
		type_str = strings.TrimSpace(type_str)
		type_desc = strings.TrimSpace(type_desc)
		type_name = strings.TrimSpace(type_name)
		t,err:=getFieldType(type_str)
		colName,_:=excelize.ColumnNumberToName(index-1)
		if err != nil {
			return nil,errors.New("行:1 "+"列:"+colName+" 解析错误:"+err.Error())
		}
		if type_name== "id" {
			hasId = true
		}
		field:= &gameDataField{
			fieldIndex: index,
			fieldName:  type_name,
			fieldDesc:  type_desc,
			fieldType:  t,
		}
		if !hasId {
			return nil,errors.New("class "+key+" with out id")
		}
		ret[index] = field
	}
	return ret,nil
}

func (this *excel2JsonMiniGame)parseGameFile(key string,file *gameFileSource,f *excelize.File)error{
	cols,err:=f.Cols(key)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	rows,err:=f.Rows(key)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	results := make([][]string, 0)
	for cols.Next() {
		col, _ := cols.Rows()
		results = append(results, col)
	}
	resultRows := make([][]string, 0)
	for rows.Next() {
		col, _ := rows.Columns()
		resultRows = append(resultRows, col)
	}
	fields,err:=this.parseGameFields(key,results)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	fields1:=make([]*gameDataField,0)
	for _,field:=range fields {
		fields1 = append(fields1, field)
	}
	file.fields = fields1
	sort.Slice(fields1, func(i, j int) bool {
		return fields1[i].fieldIndex<fields1[j].fieldIndex
	})

	for index,row:=range resultRows {
		if index<3 {
			continue
		}
		err:=file.parseData(row)
		if err != nil {
			this.logger.Error("格式化 "+key+" 行:"+strconv.Itoa(index+1)+err.Error())
			return errors.New("格式化 "+key+" 行:"+strconv.Itoa(index+1)+err.Error())
		}
	}
	return nil
}

func (this *excel2JsonMiniGame)fetchGameFile(key string,source *gameDataSource,f *excelize.File)error {
	file:=&gameFileSource{data: make([]gameData,0),enums: map[string]int{}}
	source.files[key] = file
	err := this.parseGameFile(key,file,f)
	if err != nil {
		return err
	}
	return nil
}

func (this *excel2JsonMiniGame)fetchGameDataSource(source *gameDataSource,p string)error {
	f, err := excelize.OpenFile(p)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	for _,key:=range f.GetSheetList() {
		if !this.checkClassName(key) {
			continue
		}
		this.logger.Warnf("开始读取分表:"+key)
		err = this.fetchGameFile(key,source,f)
		if err != nil {
			return err
		}
		this.logger.Warnf("读取分表完成:"+key)
	}
	return nil
}

func (this *excel2JsonMiniGame)generateEnums(name string,data *gameFileSource)string {
	out:="export enum "+name+"Enum {\n"
	enumsArr:=make([]slice.KVIntString,0)

	for k,v:=range data.enums {
		descStr:=""
		for _,data:=range data.data {
			if data["id"] == v {
				if desc,ok:=data["desc"];ok{
					descStr = desc.(string)
				}
			}
		}
		value:="\t"+k+" = "+strconv.Itoa(v)+","
		if descStr!="" {
			value+=" //"+descStr
		}
		value+="\n"
		enumsArr = append(enumsArr, slice.KVIntString{
			K: v,
			V: value,
		})
	}
	sort.Slice(enumsArr, func(i, j int) bool {
		return enumsArr[i].K<enumsArr[j].K
	})
	for _,v:=range enumsArr {
		out+=v.V
	}
	out+="}\n"
	return out
}


func (this *excel2JsonMiniGame)generateEnumGetter(name string,data *gameFileSource)string {
	out:=""
	enumsArr:=make([]slice.KVIntString,0)
	for k,v:=range data.enums {
		enumsArr = append(enumsArr, slice.KVIntString{
			K: v,
			V: "\t"+`get `+k+`():I`+name+`Data{
		return this.getById(`+strconv.Itoa(v)+`)
	}
`,
		})
	}
	sort.Slice(enumsArr, func(i, j int) bool {
		return enumsArr[i].K<enumsArr[j].K
	})
	for _,v:=range enumsArr {
		out+=v.V
	}
	return out
}

func (this *excel2JsonMiniGame)generateTsSchema(data *gameFileSource,name string,p string)error{
	lowerName:=stringutil.CamelToSnake(name)
	lowerName+="_data_source"

	this.logger.Warnf("开始生成Ts文件",name,path.Join(p,lowerName+".ts"))
	tsPath := util.FindFile(p, lowerName+".ts", false)
	if tsPath == "" {
		tsPath = path.Join(p, lowerName+".ts")
		util.CreateFile(tsPath,0666)
	}
	var template string
	if this.embedJson {
		str,err:=this.generateJson(data)
		if err != nil {
			return err
		}
		str = strings.ReplaceAll(str,`\`,`\\`)
		template = "const json = `"+str+"`"+tsTemplate
	} else {
		template = `const json = require("./${lowerclass}.json")`+tsTemplate
	}
	output:=strings.Replace(template,`${lowerclass}`,lowerName,-1)
	output=strings.Replace(output,`${class}`,name,-1)
	output=strings.Replace(output,`${fields}`,data.fieldString(),-1)
	output=strings.Replace(output,`${enums}`,this.generateEnums(name,data),-1)
	output=strings.Replace(output,`${enumsGetter}`,this.generateEnumGetter(name,data),-1)
	log.Warnf(tsPath)
	err:=ioutil.WriteFile(tsPath,[]byte(output),0666)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	this.logger.Warnf("生成Ts文件成功",name,path.Join(p,lowerName+".ts"))
	return nil
}

func (this *excel2JsonMiniGame)generateJson(data *gameFileSource)(string,error) {
	strarr,err :=json.Marshal(data.data)
	if err != nil {
		this.logger.Error(err.Error())
		return "",err
	}
	return string(strarr),nil

}

func (this *excel2JsonMiniGame) generateJsonData(data *gameFileSource,name string,p string)error{
	lowerName:=stringutil.CamelToSnake(name)
	lowerName+="_data_source"
	this.logger.Warnf("开始生成Json文件",name,path.Join(p,lowerName+".json"))
	output,err:=this.generateJson(data)
	if err != nil {
		return err
	}
	jsonPath := util.FindFile(p, lowerName+".json", false)
	if jsonPath == "" {
		jsonPath = path.Join(p, lowerName+".json")
		util.CreateFile(jsonPath,0666)
	}
	log.Warnf(jsonPath)
	err=ioutil.WriteFile(jsonPath,[]byte(output),0666)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	this.logger.Warnf("生成Json文件成功",name,path.Join(p,lowerName+".json"))
	return nil
}

const tsTemplate = `
export interface I${class}Data {
${fields}
}

${enums}

class _${class}DataSource {
    protected data:Map<number, I${class}Data> = new Map<number, I${class}Data>()
    protected strMap:Map<string, I${class}Data> = new Map<string, I${class}Data>()
    getById(id:number):I${class}Data {
        return this.data.get(id)
    }
    getByName(name:string):I${class}Data {
        return this.strMap.get(name)
    }
    all():I${class}Data[]{
        let ret:I${class}Data[] = []
        this.data.forEach((iter)=>{
            ret.push(iter)
        })
        ret.sort((a,b)=>{
            return a.id-b.id
        })
        return ret
    }
    constructor(json:string) {
        let objs = JSON.parse(json)
        for (let obj of objs) {
            this.data.set(obj.id,obj)
			if (obj["name"]) {
				this.strMap.set(obj["name"],obj)
			}
        }
    }
${enumsGetter}
}

export const ${class}DataSource = new _${class}DataSource(json)
`