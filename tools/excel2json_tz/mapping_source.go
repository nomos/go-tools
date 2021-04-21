package excel2json_tz

import "github.com/360EntSecGroup-Skylar/excelize"

func newMappingSource(file *excelize.File) *mappingSource {
	return &mappingSource{
		file: file,
		mapping: map[string]string{},
	}
}

type mappingSource struct {
	file    *excelize.File
	mapping map[string]string
}

func (this *mappingSource) Load() error {
	return nil
}
