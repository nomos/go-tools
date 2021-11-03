//this is a generate file,do not modify it!
package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/ecs"
	"github.com/nomos/go-lokas/protocol"
	"reflect"
	"time"
)

var _ lokas.IComponent = (*PngFile)(nil)

type PngFile struct {
	ecs.Component `json:"-" bson:"-"`
	Path string 
	ModTime time.Time 
	Width int32 
	Height int32 
	Data []byte 
}

func (this *PngFile) GetId()(protocol.BINARY_TAG,error){
	return protocol.GetTypeRegistry().GetTagByType(reflect.TypeOf(this).Elem())
}

func (this *PngFile) Serializable()protocol.ISerializable {
	return this
}
