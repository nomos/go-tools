package excel2json

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/util"
	"io/ioutil"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type fieldType int

const (
	type_string = iota+1
	type_int
	type_float
)

func (this fieldType) String()string{
	switch this {
	case type_string:
		return "string"
	case type_int:
		return "number"
	case type_float:
		return "number"
	default:
		return ""
	}
}

func (this fieldType) encode(s interface{}) (string,error){
	switch this {
	case type_string:
		return s.(string),nil
	case type_int:
		return strconv.Itoa(s.(int)),nil
	case type_float:
		return strconv.FormatFloat(s.(float64), 'g', -1, 64),nil
	default:
		return "",errors.New("unrecognized type")
	}
}


func (this fieldType) decode(s string) (interface{},error){
	switch this {
	case type_string:
		return s,nil
	case type_int:
		ret,err:=strconv.Atoi(s)
		if err != nil {
			return nil,err
		}
		return ret,nil
	case type_float:
		ret,err:=strconv.ParseFloat(s, 64)
		if err != nil {
			return nil,err
		}
		return ret,nil
	default:
		return 0,errors.New("unrecognized type:"+s)
	}
}

func getFieldType(s string)(fieldType,error) {
	switch s {
	case "string":
		return type_string,nil
	case "int":
		return type_int,nil
	case "float":
		return type_float,nil
	default:
		return 0,errors.New("unrecognized type:"+s)
	}
}

type gameDataField struct {
	fieldIndex int
	fieldName string
	fieldDesc string
	fieldType fieldType
}

func (this *gameDataField) string(data gameData)string {

	return ""
}

type gameData struct {
	id int
	data []interface{}
}

//对应一个文件
type gameFileSource struct {
	fields []*gameDataField
	data []*gameData
}

func (this *gameFileSource) fieldString()string {
	ret:=""
	for _,field:=range this.fields {
		ret+="\t"+field.fieldName+" "+field.fieldType.String()+"\n"
	}
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
	data:=&gameData{
		id:       0,
		data:     []interface{}{},
	}
	for _,field:=range this.fields {
		str:=in[field.fieldIndex]
		t:=field.fieldType
		d,err:=t.decode(str)
		if err != nil {
			return err
		}
		if field.fieldName == "id" {
			data.id = d.(int)
		}
		data.data = append(data.data, d)
	}
	this.data = append(this.data, data)
	return nil
}

//对应一个目录下所有文件
type gameDataSource struct {
	files map[string]*gameFileSource
}

type excel2JsonMiniGame struct {
	logger log.ILogger
}

func (this *excel2JsonMiniGame)generateData(source *gameDataSource,p string)error{
	for key,data:=range source.files{
		err:=this.generateTsSchema(data,key,p)
		if err != nil {
			return err
		}
		err=this.generatJsonData(data,key,p)
		if err != nil {
			return err
		}
	}
	return nil
}

func Excel2JsonMiniGame(excelPath,distPath string,logger log.ILogger) {
	ret := &excel2JsonMiniGame{}
	if logger!=nil {
		ret.logger = logger
	} else {
		ret.logger = log.DefaultLogger()
	}
	defer func() {
		if r := recover(); r != nil {
			log.Error(r.(error).Error())
		}
	}()
	err:=ret.generate(excelPath,distPath)
	if err != nil {
		log.Error(err.Error())
	}
}

func  (this *excel2JsonMiniGame)generate(excelPath,distPath string)error{
	paths,err := util.WalkDirFiles(excelPath,false)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	paths = util.FilterFileWithExt(".xlsx",paths)
	this.logger.Infof("paths",paths)
	var source = &gameDataSource{files: map[string]*gameFileSource{}}
	for _,p:=range paths {
		this.logger.Warn("开始读取"+p)
		err=this.fetchGameDataSource(source,p)
		if err != nil {
			this.logger.Error(err.Error())
			return err
		}
	}
	this.logger.Warn("读取完成")
	err=this.generateData(source,distPath)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	return nil
}

func IsCapitalWord(s string)bool{
	return regexp.MustCompile(`^[A-Z]\w+`).FindString(s) == s
}

func (this *excel2JsonMiniGame)checkClassName(s string)bool{
	if regexp.MustCompile(`Sheet[0-9]*`).FindString(s) == s {
		return false
	}
	return IsCapitalWord(s)
}


func (this *excel2JsonMiniGame) parseGameFields (data [][]string) (map[int]*gameDataField,error) {
	ret:=make(map[int]*gameDataField)
	for index,col:=range data {
		if col[0] == "" {
			continue
		}
		type_str:=col[0]
		type_desc:=col[1]
		type_name:=col[2]
		t,err:=getFieldType(type_str)
		if err != nil {
			this.logger.Error(err.Error())
			return nil,err
		}
		field:= &gameDataField{
			fieldIndex: index,
			fieldName:  type_name,
			fieldDesc:  type_desc,
			fieldType:  t,
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
	fields,err:=this.parseGameFields(results)
	if err != nil {
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
			this.logger.Error(err.Error())
			return err
		}
	}
	return nil
}


func (this *excel2JsonMiniGame)fetchGameFile(key string,source *gameDataSource,f *excelize.File)error {
	this.logger.Warn("fetchGameFile")

	file:=&gameFileSource{data: []*gameData{}}
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
		err = this.fetchGameFile(key,source,f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *excel2JsonMiniGame)generateTsSchema(data *gameFileSource,name string,p string)error{
	tsPath := util.FindFile(p, name+".ts", false)
	if tsPath == "" {
		tsPath = path.Join(p, name+".ts")
		util.CreateFile(tsPath)
	}
	lowerName:=strings.ToLower(name)
	output:=strings.Replace(tsTemplate,`${lowerclass}`,lowerName,-1)
	output=strings.Replace(output,`${class}`,name,-1)
	output=strings.Replace(output,`${fields}`,data.fieldString(),-1)
	log.Warnf(tsPath)
	err:=ioutil.WriteFile(tsPath,[]byte(output),0)
	if err != nil {
		this.logger.Error(err.Error())
		return err
	}
	this.logger.Warnf("生成Ts文件",name,path.Join(p,name,".ts"))
	//this.logger.Info(output)
	return nil
}

func (this *excel2JsonMiniGame)generatJsonData(data *gameFileSource,name string,p string)error{
	this.logger.Warnf("生成Json文件",name,path.Join(p,name,".json"))
	return nil
}

const tsTemplate = `const json = require("./${lowerclass}.json")

export interface I${class}Data {
${fields}
}

class _${class}DataSource {
    protected data:Map<number, I${class}Data> = new Map<number, I${class}Data>()
    getById(id:number):IDragonData {
        return this.data.get(id)
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
        for (let obj of objs.data) {
            this.data.set(obj.id,obj)
        }
    }
}

export const ${class}DataSource = new _${class}DataSource(json)
`