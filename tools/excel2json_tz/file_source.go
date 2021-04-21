package excel2json_tz

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/nomos/go-log/log"
)

func newFileSource(p string)*fileSource{
	ret:=&fileSource{
		path:    p,
		sheets:  []*sheetSource{},
		mapping: newMappingSource(),
	}
	return ret
}

type fileSource struct {
	path string
	sheets []*sheetSource
	mapping *mappingSource
	file *excelize.File
}

func (this *fileSource) Load()error {
	var err error
	this.file, err = excelize.OpenFile(this.path)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err=this.readMappingSource()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (this *fileSource) readMappingSource()error {
	var err error
	for _,key:=range this.file.GetSheetList() {
		if key == MappingTag {
			this.mapping.ReadIn()
		}
	}
	return errors.New("mapping source not found")
}