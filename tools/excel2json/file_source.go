package excel2json

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/nomos/go-lokas/log"
	"regexp"
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
	file *excelize.File
}

func (this *fileSource) Load()error {
	log.Infof("fileSource:Load",this.path)
	var err error
	this.file, err = excelize.OpenFile(this.path)
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

func (this *fileSource)checkClassName(s string)bool{
	if regexp.MustCompile(`Sheet[0-9]*`).FindString(s) == s {
		return false
	}
	return IsCapitalWord(s)
}

func (this *fileSource) readSheetSource()error {
	for _,v:=range this.file.GetSheetList() {
		if this.checkClassName(v) {
			rows,err:=this.file.Rows(v)
			if err != nil {
				log.Error(err.Error())
				return err
			}
			r:=0
			for rows.Next() {
				r++
			}
			if r<3 {
				continue
			}
			this.sheets[v] = NewSheetSource(this.file,v,v)
		}
	}

	for _,s:=range this.sheets {
		err:=s.Load()
		if err != nil {
			return err
		}
	}
	return nil
}