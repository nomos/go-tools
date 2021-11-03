//this is a generate file,edit implement on this file only!
package erpc

import (
	"github.com/nomos/go-lokas"
	"time"
)

func NewPngFile(path string,modTime time.Time,width,height int32,data []byte)*PngFile{
	ret:=&PngFile{
		Path: path,
		ModTime: modTime,
		Data: data,
	}
	if data==nil {
		ret.Data = []byte{}
	}
	return ret
}

func (this *PngFile) OnAdd(e lokas.IEntity, r lokas.IRuntime) {
	
}

func (this *PngFile) OnRemove(e lokas.IEntity, r lokas.IRuntime) {
	
}

func (this *PngFile) OnCreate(r lokas.IRuntime) {
	
}

func (this *PngFile) OnDestroy(r lokas.IRuntime) {
	
}