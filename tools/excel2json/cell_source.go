package excel2json

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/nomos/go-lokas/log"
	"regexp"
	"strconv"
)

type CellSource struct {
	Row int
	Col int
	Value string
}

func NewCellSource(row int,col int,value string)*CellSource {
	ret:=&CellSource{
		Row:   row,
		Col:   col,
		Value: value,
	}
	return ret
}

func (this *CellSource) RowName()string{
	return strconv.Itoa(this.Row+1)
}

func (this *CellSource) CellName()string{
	return this.ColName()+this.RowName()
}

func (this *CellSource) ColName()string {
	ret,_:=excelize.ColumnNumberToName(this.Col+1)
	return ret
}

func (this *CellSource) NewError(msg string)*ExcelError{
	return NewExcellError(this.Row,this.Col,msg)
}

func (this *CellSource) Float()(float64,error) {
	if this.Value == "" {
		return 0,nil
	}
	r:=regexp.MustCompile(`^(([1-9]\d*)|0)*[.]*\d*`)
	if r.FindString(this.Value)==this.Value {
		ret,err:= strconv.ParseFloat(this.Value,10)
		if err != nil {
			log.Error(err.Error())
			return 0,this.NewError(err.Error())
		}
		return ret,nil
	}
	return 0,this.NewError("not a float")
}

func (this *CellSource) Int()(int,error) {
	if this.Value == "" {
		return 0,nil
	}
	r:=regexp.MustCompile(`^([-]*)[0-9]+`)
	if r.FindString(this.Value)==this.Value {
		ret,err:= strconv.Atoi(this.Value)
		if err != nil {
			log.Error(err.Error())
			return 0,this.NewError(err.Error())
		}
		return ret,nil
	}
	return 0,this.NewError("not a int")
}

func (this *CellSource) String()string {
	return this.Value
}
