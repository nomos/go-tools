package excel2json_tz

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/nomos/go-log/log"
)

func newFileSource(p string)*fileSource{
	ret:=&fileSource{
		path:    p,
		sheets:  map[string]*SheetSource{},
	}
	return ret
}

type fileSource struct {
	path string
	sheets map[string]*SheetSource
	mapping *mappingSource
	file *excelize.File
}

func (this *fileSource) Load()error {
	log.Infof("fileSource:Load",this.path)
	var err error
	this.file, err = excelize.OpenFile(this.path)
	this.mapping = newMappingSource(this.file)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err=this.readMappingSource()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	err=this.readSheetSource()
	if err != nil {
		return err
	}
	return nil
}

func (this *fileSource) readMappingSource()error {
	var err error
	for _,key:=range this.file.GetSheetList() {
		if key == MappingTag {
			err=this.mapping.Load()
			if err != nil {
				log.Error(err.Error())
				return err
			}
			return nil
		}
	}
	return errors.New("mapping source not found")
}

func (this *fileSource) readSheetSource()error {
	for k,v:=range this.mapping.Mapping {
		this.sheets[k] = NewSheetSource(this.file,k,v)
	}
	for _,s:=range this.sheets {
		err:=s.Load()
		if err != nil {
			return err
		}
	}
	return nil
}