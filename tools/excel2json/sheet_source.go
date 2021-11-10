package excel2json

import (
	"encoding/json"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/slice"
	"github.com/nomos/go-lokas/util/stringutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type SheetSource struct {
	Name       string
	DescName   string
	enumField  *DataField
	enums map[string]int32
	DataFields []*DataField
	rows       []*RowSource
	Data       []*DataLine
	file       *excelize.File
}

func NewSheetSource(file *excelize.File,name string,descName string)*SheetSource {
	ret:=&SheetSource{
		Name:       strings.Join(stringutil.SplitCamelCaseCapitalize(name),""),
		DescName:   descName,
		DataFields: []*DataField{},
		enums: map[string]int32{},
		rows:       make([]*RowSource, 0),
		Data:[]*DataLine{},
		file:       file,
	}
	return ret
}

func (this *SheetSource) GetRow(index int)*RowSource{
	return this.rows[index]
}

func (this *SheetSource) GetRowByName(name string)*RowSource{
	index,err:=strconv.Atoi(name)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return this.GetRow(index-1)
}

func (this *SheetSource) GetCell(row int,col int)(*CellSource,error) {
	r:=this.GetRow(row)
	if r==nil {
		return nil,NewExcellError(row,col,"row is empty")
	}
	cell,err:=r.GetCell(col)
	if err != nil {
		return nil,err
	}
	return cell,nil
}

func (this *SheetSource) GetCellByName(name string)(*CellSource,error){
	col,row,err:=excelize.CellNameToCoordinates(name)
	if err != nil {
		return nil,NewExcellError(row,col,err.Error())
	}
	return this.GetCell(row-1,col-1)
}

func (this *SheetSource) Load()error {
	log.Infof("SheetSource:Load:",this.Name,this.DescName)
	rows,err:=this.file.Rows(this.DescName)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	this.rows = make([]*RowSource, 0)
	r:=0
	for rows.Next() {
		col, _ := rows.Columns()
		row:=NewRowSource(r,col)
		this.rows = append(this.rows, row)
		r++
	}
	if len(this.rows)<3 {
		return errors.New("wrong excel format")
	}
	err=this.LoadDataFields()
	if err != nil {
		return err
	}
	err=this.LoadData()
	if err != nil {
		return err
	}
	return nil
}

func (this *SheetSource) LoadDataFields()error{
	startIndex:=0
	for _,cell:=range this.rows[0].Cells {
		if cell.String()=="" {
			continue
		}
		field:=NewDataField(this,cell.Col,startIndex)
		err:=field.Load()
		if err != nil {
			log.Error(err.Error())
			return err
		}
		startIndex++
		if field.Typ==type_tag {
			this.enumField = field
			continue
		}
		this.DataFields = append(this.DataFields, field)
	}
	cols,desc,typ,len:=this.fieldString()
	lenOffset = len
	log.Infof("行号 "+cols)
	log.Infof("     "+desc)
	log.Infof("     "+typ)
	return nil
}

func getDescLen(s string)int{
	spaceLen:=strings.Count(s," ")
	re:=regexp.MustCompile("[!-/]|[:-@]|[\\[-`]|[A-z]|[0-9]|[|]")
	charLen:=len(re.FindAllString(s,-1))
	hanLen:=len(s)-spaceLen-charLen
	return hanLen*2/3+spaceLen+charLen
}

func (this *SheetSource) fieldString()(string,string,string,[]int) {
	cols:=""
	desc:=""
	ret:=""
	loc:=[]int{}
	for _,field:=range this.DataFields {
		if len(ret)<getDescLen(desc) {
			for len(ret)<getDescLen(desc) {
				ret+=" "
			}
			loc = append(loc, getDescLen(desc))
		} else {
			for getDescLen(desc)<len(ret) {
				desc+=" "
			}
			loc = append(loc, len(ret))
		}
		for len(cols)<=len(ret) {
			cols+=" "
		}
		cols+=field.ColName()
		cols+=" "
		desc+=field.Desc
		desc+=" "
		ret+="["
		ret+=field.Name
		ret+="]"
		ret+=field.Typ.GoString()
		ret+=" "
	}
	desc = strings.TrimRight(desc," ")
	ret = strings.TrimRight(ret," ")
	return cols,desc,ret,loc
}

var lineIgnores = []string{}
var lenOffset []int

func (this *SheetSource) LoadData()error{
	lineIgnores = []string{}
	for i:=3;i<len(this.rows);i++ {
		err:=this.readLine(this.rows[i])
		if err != nil {
			return err
		}
	}
	for _,v:=range this.Data {
		log.Infof(stringutil.AddStringGap("["+v.Row.RowName()+"]",7,1)+v.LogString(lenOffset))
	}
	if len(lineIgnores)>0 {
		log.Warnf("跳过空行:"+strings.Join(lineIgnores,"|"))
	}
	log.Infof("read "+strconv.Itoa(len(this.Data))+" lines")
	return nil
}

func (this *SheetSource) readLine(row *RowSource)error{
	line:=NewDataLine(this,row)
	var idx int32
	for i,field:=range this.DataFields {
		if i==0 {
			cell,err:=row.GetCell(field.ColIndex)
			if err != nil {
				lineIgnores = append(lineIgnores, row.RowName())
				return nil
			}
			if cell.String()=="" {
				lineIgnores = append(lineIgnores, row.RowName())
				return nil
			}
			var id int
			id,err = strconv.Atoi(cell.String())
			if err != nil {
				log.Error(err.Error())
				return err
			}
			idx = int32(id)

		}
		data,err:=field.ReadIn(row)
		if err != nil {
			return err
		}
		line.Append(data)
	}
	cell,err:=row.GetCell(this.enumField.ColIndex)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if cell.String()!="" {
		this.enums[cell.String()] = idx
	}

	this.Data = append(this.Data, line)
	return nil
}

const __gostr = `package gamedata

type {ClassName} struct {
{ClassFields}
}
`

func (this *SheetSource) GenerateGoString()string{
	ret:=__gostr
	ret = strings.Replace(ret,"{ClassName}",strings.Join(stringutil.SplitCamelCaseCapitalize(this.Name),""),-1)
	ret = strings.Replace(ret,"{ClassFields}",this.generateGoFields(),-1)
	log.Warnf("generateGoFields",ret)
	return ret
}


const __tsstr = `
export interface I${class}Data {
${fields}
}

${enums}

export class ${class}DataSource {
    protected data:Map<number, I${class}Data> = new Map<number, I${class}Data>()
    protected strMap:Map<string, I${class}Data> = new Map<string, I${class}Data>()
	constructor() {
    }
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
    load(objs:NumericDict<any>) {
        for (let id in objs) {
            let obj = objs[id]
            this.data.set(obj.id,obj)
			if (obj["name"]) {
				this.strMap.set(obj["name"],obj)
			}
        }
    }
${enumsGetter}
}
`

func (this *SheetSource)generateTsEnums()string {
	out:="export enum "+this.Name+"Enum {\n"
	enumsArr:=make([]slice.KVIntString,0)

	for k,v:=range this.enums {
		descStr:=""
		value:="\t"+k+" = "+strconv.Itoa(int(v))+","
		if descStr!="" {
			value+=" //"+descStr
		}
		value+="\n"
		enumsArr = append(enumsArr, slice.KVIntString{
			K: int(v),
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

func (this *SheetSource) generateTsEnumGetter()string {
	out:=""
	enumsArr:=make([]slice.KVIntString,0)
	for k,v:=range this.enums {
		enumsArr = append(enumsArr, slice.KVIntString{
			K: int(v),
			V: "\t"+`get `+k+`():I`+this.Name+`Data{
		return this.getById(`+strconv.Itoa(int(v))+`)
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

func (this *SheetSource) GenerateTsString()string{
	ret:=__tsstr
	ret=strings.Replace(ret,`${class}`,this.Name,-1)
	ret=strings.Replace(ret,`${fields}`,this.generateTsFields(),-1)
	ret=strings.Replace(ret,`${enums}`,this.generateTsEnums(),-1)
	ret=strings.Replace(ret,`${enumsGetter}`,this.generateTsEnumGetter(),-1)
	return ret
}


func (this *SheetSource) GetMainFieldType()string{
	f:=this.DataFields[0]
	return f.Typ.GoString()
}


func (this *SheetSource) GetTsImportFieldString()string {
	importName:=this.Name+"DataSource"
	pathName:=stringutil.CamelToSnake(this.Name)+"_source"
	return `import {`+importName+`} from "./`+pathName+`";`
}

func (this *SheetSource) GetTsLoadFieldString()string{
	return "\t\t"+`this.`+this.Name+`.load(objs["`+this.Name+`"])`
}

func (this *SheetSource) GetTsSourceFieldString()string {
	return "\t"+this.Name+`:`+this.Name+"DataSource = new "+this.Name+"DataSource()"
}

func (this *SheetSource) GetGoDataFieldString()string{
	return "\t"+this.Name+" map["+this.GetMainFieldType()+"]*"+this.Name
}

func (this *SheetSource) GetGoInitFieldString()string{
	return "\tthis."+this.Name+"=make(map["+this.GetMainFieldType()+"]*"+this.Name+")"
}

func (this *SheetSource) generateTsFields()string{
	ret:=""
	for _,f:=range this.DataFields {
		if f.Name == "#" {
			continue
		}
		ret+="\t"+stringutil.FirstToLower(f.Name)+":"+f.Typ.TsString()+"\n"
	}
	ret = strings.TrimRight(ret,"\n")
	return ret
}

func (this *SheetSource) generateGoFields()string{
	ret:=""
	for _,f:=range this.DataFields {
		ret+="\t"
		ret+=f.Name
		ret+=" "
		ret+=f.Typ.GoString()
		ret+=" //"
		ret+=f.Desc
		ret+="\n"
	}
	ret = strings.TrimRight(ret,"\n")
	return ret
}

func (this *SheetSource) GetJsonMap()map[string]*orderedmap.OrderedMap{
	m:=map[string]*orderedmap.OrderedMap{}
	for _,l:=range this.Data {
		str,_:=l.line[0].String()
		m[str] = l.Map
	}
	return m
}

func (this *SheetSource) GenerateJson()string{
	m:=[]*orderedmap.OrderedMap{}
	for _,l:=range this.Data {
		m = append(m, l.Map)
	}
	data,_:=json.Marshal(m)
	ret:=string(data)
	log.Warnf("GenerateJson",ret)
	return ret
}