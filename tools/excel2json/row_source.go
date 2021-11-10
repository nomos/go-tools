package excel2json

import (
	"strconv"
)

type RowSource struct {
	Row int
	Cells []*CellSource
}

func NewRowSource(row int,source []string)*RowSource{
	ret:= &RowSource{
		Row:row,
		Cells: []*CellSource{},
	}
	ret.loadCells(source)
	return ret
}

func (this *RowSource) RowName()string{
	return strconv.Itoa(this.Row+1)
}


func (this *RowSource) loadCells(cells []string){
	for colIndex,col:=range cells {
		cell:=NewCellSource(this.Row,colIndex,col)
		this.Cells = append(this.Cells, cell)
	}
}

func (this *RowSource) Length()int {
	return len(this.Cells)
}

func (this *RowSource) GetCell(col int)(*CellSource,error){
	if col>=len(this.Cells) {
		return nil,NewExcellError(this.Row,col,ERR_CELL_NOT_EXIST)
	}
	return this.Cells[col],nil
}