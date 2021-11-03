//this is a generate file,edit implement on this file only!
package erpc

import (
	"github.com/nomos/go-lokas"
)

func NewPngFolderData()*PngFolderData{
	ret:=&PngFolderData{
		Files: []*PngFile{},
	}
	return ret
}

func (this *PngFolderData) AddFile(file *PngFile)*PngFile{
	this.Files = append(this.Files, file)
	return file
}

func (this *PngFolderData) OnAdd(e lokas.IEntity, r lokas.IRuntime) {
	
}

func (this *PngFolderData) OnRemove(e lokas.IEntity, r lokas.IRuntime) {
	
}

func (this *PngFolderData) OnCreate(r lokas.IRuntime) {
	
}

func (this *PngFolderData) OnDestroy(r lokas.IRuntime) {
	
}