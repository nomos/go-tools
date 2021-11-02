//this is a generate file,do not modify it!
package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/ecs"
	"github.com/nomos/go-lokas/protocol"
	"reflect"
)

var _ lokas.IComponent = (*ConsoleEvent)(nil)

type ConsoleEvent struct {
	ecs.Component `json:"-" bson:"-"`
	Text string
}

func (this *ConsoleEvent) GetId()(protocol.BINARY_TAG,error){
	return protocol.GetTypeRegistry().GetTagByType(reflect.TypeOf(this).Elem())
}

func (this *ConsoleEvent) Serializable()protocol.ISerializable {
	return this
}
