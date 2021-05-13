package excel2json_tz

import (
	"encoding/json"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/iancoleman/orderedmap"
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/util/stringutil"
	"regexp"
	"strconv"
	"strings"
)

type SheetSource struct {
	Name       string
	DescName   string
	DataFields []*DataField
	rows       []*RowSource
	Data       []*DataLine
	file       *excelize.File
}

func NewSheetSource(file *excelize.File,name string,descName string)*SheetSource {
	ret:=&SheetSource{
		Name:       name,
		DescName:   descName,
		DataFields: []*DataField{},
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
	rows,err:=this.file.Rows(this.Name)
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
	if len(this.rows)<6 {
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
	this.GenerateGoString()
	this.GenerateJson()
	return nil
}

func (this *SheetSource) LoadDataFields()error{
	for _,cell:=range this.rows[0].Cells {
		index,err:=cell.Int()
		if err != nil {
			log.Error(err.Error())
			return err
		}
		field:=NewDataField(this,cell.Col,index)
		this.DataFields = append(this.DataFields, field)
	}
	for _,field:=range this.DataFields {
		err:=field.Load()
		if err != nil {
			log.Error(err.Error())
			return err
		}
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
		ret+=field.Typ.String()
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
	for i:=5;i<len(this.rows);i++ {
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
		}
		data,err:=field.ReadIn(row)
		if err != nil {
			return err
		}
		line.Append(data)
	}
	this.Data = append(this.Data, line)
	return nil
}

const __gostr = `package data

type {ClassName} struct {
{ClassFields}
}
`


func (this *SheetSource) GenerateGoString()string{
	ret:=__gostr
	ret = strings.Replace(ret,"{ClassName}",this.Name,-1)
	ret = strings.Replace(ret,"ClassFields",this.generateGoFields(),-1)
	log.Warnf("generateGoFields",ret)
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