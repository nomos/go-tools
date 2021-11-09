package excel2json

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util"
	"regexp"
)

func newDirSource(dir string)*dirSource{
	ret:=&dirSource{
		dir:   dir,
		files: map[string]*fileSource{},
	}
	return ret
}

type dirSource struct {
	dir string
	files map[string]*fileSource
}

func (this *dirSource) Load()error {
	log.Infof("dirSource:Load",this.dir)
	paths,err := util.WalkDirFiles(this.dir,false)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	paths = util.FilterFileWithExt(paths,".xlsx",".xlsm")
	paths = util.FilterFileWithFunc(paths, func(s string) bool {
		return !regexp.MustCompile(`[~][$]`).MatchString(s)
	})
	for _,p:=range paths {
		this.addFileSource(p)
	}
	for _,f:=range this.files {
		err:=f.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *dirSource) addFileSource(p string) {
	log.Infof("dirSource:addFileSource:",p)
	this.files[p] = newFileSource(p)
}