package excel2json_tz

import (
	"github.com/iancoleman/orderedmap"
	"strconv"
	"strings"
)

type Data struct {
	Field *DataField
	Cell *CellSource
	Value interface{}
}

func NewData(field *DataField,cell*CellSource,value interface{})*Data{
	ret:=&Data{
		Field: field,
		Value: value,
		Cell:cell,
	}
	return ret
}

func (this *Data) Int()(int,error){

	s,ok:=this.Value.(int)
	if !ok {
		return 0,this.Cell.NewError("value is not a int")
	}
	return s,nil
}

func (this *Data) Float()(float64,error){
	s,ok:=this.Value.(float64)
	if !ok {
		return 0,this.Cell.NewError("value is not a float")
	}
	return s,nil
}

func (this *Data) String()(string,error){
	switch this.Field.Typ {
	case TypeInt:
		s,err:=this.Int()
		if err != nil {
			return "",err
		}
		return strconv.Itoa(s),nil
	case TypeFloat:
		s,err:=this.Float()
		if err != nil {
			return "",err
		}
		return strconv.FormatFloat(s, 'f', -1, 32),nil
	case TypeString:
		s,ok:=this.Value.(string)
		if !ok {
			return "",this.Cell.NewError("value is not a string")
		}
		return s,nil
	default:
		return "",this.Cell.NewError("field error")
	}
}

type DataLine struct {
	sheet *SheetSource

	line  []*Data
	Map *orderedmap.OrderedMap
	Row *RowSource
}

func NewDataLine(sheet *SheetSource,row *RowSource)*DataLine{
	ret:=&DataLine{
		sheet: sheet,
		line:  []*Data{},
		Map:orderedmap.New(),
		Row:row,
	}
	return ret
}

func (this *DataLine) Append(d *Data){
	this.line = append(this.line, d)
	this.Map.Set(d.Field.Name,d.Value)
}

func (this *DataLine) LogString(lenOffset []int)string {
	strArr:=[]string{}
	for _,v:=range this.line {
		s,_:=v.String()
		strArr = append(strArr, s)
	}
	ret:=""
	for i,v:=range strArr {
		for getDescLen(ret)<lenOffset[i] {
			ret+=" "
		}
		ret+=v
		ret+=" "
	}
	return strings.TrimRight(ret," ")
}