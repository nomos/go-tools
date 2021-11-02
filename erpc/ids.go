//this is a generated file,do not modify it!!!
package erpc

import (
	"github.com/nomos/go-lokas/protocol"
	"reflect"
)

const (
	TAG_CONSOLE_EVENT  protocol.BINARY_TAG = 221
)

func init() {
	protocol.GetTypeRegistry().RegistryType(TAG_CONSOLE_EVENT,reflect.TypeOf((*ConsoleEvent)(nil)).Elem())
}


