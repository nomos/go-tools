package erpc

import (
	"github.com/nomos/go-lokas/protocol"
	"reflect"
)

const (
	TAG_ConsoleEvent = 221
)

func init(){
	protocol.GetTypeRegistry().RegistryType(TAG_ConsoleEvent,reflect.TypeOf((*ConsoleEvent)(nil)).Elem())
}
