package excel2json_tz

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/nomos/go-log/log"
)

func newMappingSource(file *excelize.File) *mappingSource {
	return &mappingSource{
		file:    file,
		Mapping: map[string]string{},
	}
}

type mappingSource struct {
	file    *excelize.File
	Mapping map[string]string
}

func (this *mappingSource) Load() error {
	rows,err:=this.file.Rows("Mapping")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	results := make([][]string, 0)
	for rows.Next() {
		col, _ := rows.Columns()
		results = append(results, col)
	}
	if len(results)<2 {
		return errors.New("wrong excel format")
	}
	for i,col:=range results {
		if i>0 {
			log.Infof(i,col)
			if len(col)<2 {
				this.Mapping = map[string]string{}
				return errors.New("wrong excel format")
			}
			this.Mapping[col[0]] = col[1]
		}
	}
	return nil
}
