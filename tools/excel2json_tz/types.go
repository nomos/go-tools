package excel2json_tz

import "errors"

const MappingTag = "Mapping"

type FieldType int

const (
	ERR_CELL_NOT_EXIST = "cell not exist"
)

const (
	EXCEL_INDEX_LINE = 0
	EXCEL_EXPORT_LINE = 1
	EXCEL_NAME_LINE  = 2
	EXCEL_TYPE_LINE  = 3
	EXCEL_DESC_LINE  = 4
)

const (
	TypeString FieldType= iota + 1
	TypeInt
	TypeFloat

)

type ExportType string

const (
	TypeAll    ExportType = "All"
	TypeClient ExportType = "Client"
	TypeServer ExportType = "Server"
	TypeIgnore ExportType = "Ignore"
)

func (this ExportType) String()string {
	return string(this)
}

func GetExportType(s string)(ExportType,error){
	if s=="" {
		return TypeAll,nil
	}
	switch s {
	case TypeAll.String():
		return TypeAll,nil
	case TypeClient.String():
		return TypeClient,nil
	case TypeServer.String():
		return TypeServer,nil
	case TypeIgnore.String():
		return TypeIgnore,nil
	default:
		return "",errors.New("type not exist:"+s)
	}
}

func (this FieldType) String() string {
	switch this {
	case TypeString:
		return "text"
	case TypeInt:
		return "int"
	case TypeFloat:
		return "float"
	default:
		return ""
	}
}

func (this FieldType) GoString() string {
	switch this {
	case TypeString:
		return "string"
	case TypeInt:
		return "int32"
	case TypeFloat:
		return "float64"
	default:
		return ""
	}
}

func GetFieldType(s string)(FieldType,error){
	switch s {
	case TypeInt.String():
		return TypeInt,nil
	case TypeFloat.String():
		return TypeFloat,nil
	case TypeString.String():
		return TypeString,nil
	default:
		return -1,errors.New("type not exist:"+s)
	}
}