package excel2json

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
)

type ExcelError struct {
	Row int
	Col int
	Message string
}

func NewExcellError(row int,col int,msg string)*ExcelError{
	ret:=&ExcelError{
		Row:     row,
		Col:     col,
		Message: msg,
	}
	return ret
}

func (this *ExcelError) IsEmptyError()bool{
	return this.Message == ERR_CELL_NOT_EXIST
}

func (this *ExcelError) Error()string{
	return this.CellName()+":"+this.Message
}

func (this *ExcelError) RowName()string{
	return strconv.Itoa(this.Row+1)
}

func (this *ExcelError) CellName()string{
	return this.ColName()+this.RowName()
}

func (this *ExcelError) ColName()string {
	ret,_:=excelize.ColumnNumberToName(this.Col+1)
	return ret
}