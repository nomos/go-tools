package excel2json_tz

const MappingTag = "Mapping"

type fieldType int

const (
	type_string = iota + 1
	type_int
	type_float
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

func GetExportType(s string)ExportType{
	if s=="" {
		return TypeAll
	}
	switch s {
	case TypeAll.String():
		return TypeAll
	case TypeClient.String():
		return TypeClient
	case TypeServer.String():
		return TypeServer
	default:
		panic("type not exist:"+s)
	}
}

func (this fieldType) String() string {
	switch this {
	case type_string:
		return "string"
	case type_int:
		return "int"
	case type_float:
		return "float"
	default:
		return ""
	}
}