package excel2json_tz

import "github.com/nomos/go-log/log"

func New(source string,dist string)*generator{
	ret:=&generator{
		dirSource:newDirSource(source),
	}
	return ret
}

type generator struct {
	*dirSource
}

func (this *generator) generate()error{
	err:=this.dirSource.Load()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}


