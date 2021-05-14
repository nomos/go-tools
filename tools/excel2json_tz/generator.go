package excel2json_tz

import (
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/util/stringutil"
	"path"
	"runtime/debug"
)

func New(source string)*Generator {
	ret:=&Generator{
		dirSource:newDirSource(source),
	}
	return ret
}

type Generator struct {
	*dirSource
}

func (this *Generator) Load()error {
	defer func() {
		if err,ok:=recover().(error);ok {
			log.Error(err.Error())
			debug.PrintStack()
		}
	}()
	err:=this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (this *Generator) Generate(p string)error{
	for _,f:=range this.dirSource.files {
		for _,s:=range f.sheets {
			goFilePath:=path.Join(p,stringutil.SplitCamelCaseLowerSlash(s.Name))+".go"
			log.Warnf(goFilePath)
			s.GenerateGoString()
			s.GenerateJson()
		}
	}
	return nil
}


