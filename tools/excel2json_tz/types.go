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