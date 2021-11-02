//this is a generate file,do not modify it!
package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/ecs"
	"github.com/nomos/go-lokas/protocol"
	"reflect"
)

var _ lokas.IComponent = (*PngFolderData)(nil)

type PngFolderData struct {
	ecs.Component `json:"-" bson:"-"`
	Files []*PngFile 
}

func (this *PngFolderData) GetId()(protocol.BINARY_TAG,error){
	return protocol.GetTypeRegistry().GetTagByType(reflect.TypeOf(this).Elem())
}

func (this *PngFolderData) Serializable()protocol.ISerializable {
	return this
}
