package excel2json_tz

import "github.com/360EntSecGroup-Skylar/excelize"

type DataField struct {
	ColIndex   int
	ExportType ExportType
	Id         int
	Typ        FieldType
	Name       string
	Desc       string
	sheet      *SheetSource
}

func NewDataField(sheet *SheetSource, colIndex int, index int) *DataField {
	ret := &DataField{
		ColIndex:   colIndex,
		ExportType: "",
		Id:         index,
		Typ:        0,
		Name:       "",
		Desc:       "",
		sheet:      sheet,
	}
	return ret
}

func (this *DataField) ColName()string {
	ret,_:=excelize.ColumnNumberToName(this.ColIndex+1)
	return ret
}

func (this *DataField) ReadIn(row *RowSource) (*Data, error) {
	cell ,err:= row.GetCell(this.ColIndex)
	if err != nil {
		if excelErr,ok:=err.(*ExcelError);ok&&excelErr.IsEmptyError() {
			return this.readIn(row.Row,NewCellSource(row.Row,this.ColIndex,""))
		}
		return nil,err
	}
	return this.readIn(row.Row, cell)
}

func (this *DataField) readIn(row int, cell *CellSource) (*Data, error) {
	switch this.Typ {
	case TypeString:
		return NewData(this, cell, cell.String()), nil
	case TypeFloat:
		s, err := cell.Float()
		if err != nil {
			return nil, err
		}
		return NewData(this, cell, s), nil
	case TypeInt:
		s, err := cell.Int()
		if err != nil {
			return nil, err
		}
		return NewData(this, cell, s), nil
	default:
		return nil, NewExcellError(row, this.ColIndex, "type is not "+this.Typ.String())
	}
}

func (this *DataField) Load() error {
	exportCell ,err:= this.sheet.GetCell(EXCEL_EXPORT_LINE, this.ColIndex)
	if err != nil {
		return err
	}
	this.ExportType, err = GetExportType(exportCell.String())
	if err != nil {
		return NewExcellError(EXCEL_EXPORT_LINE, this.ColIndex, err.Error())
	}
	typeCell,err := this.sheet.GetCell(EXCEL_TYPE_LINE, this.ColIndex)
	if err != nil {
		return err
	}
	this.Typ, err = GetFieldType(typeCell.String())
	if err != nil {
		return err
	}
	nameCell ,err:= this.sheet.GetCell(EXCEL_NAME_LINE, this.ColIndex)
	if err != nil {
		return err
	}
	this.Name = nameCell.String()
	descCell,err := this.sheet.GetCell(EXCEL_DESC_LINE, this.ColIndex)
	if err != nil {
		return err
	}
	this.Desc = descCell.String()
	return nil
}
