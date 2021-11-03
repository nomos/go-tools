//this is a generated file,do not modify it!!!
package erpc

import (
	"github.com/nomos/go-lokas/protocol"
	"reflect"
)

const (
	TAG_CONSOLE_EVENT  protocol.BINARY_TAG = 221
	TAG_PNG_FILE  protocol.BINARY_TAG = 222
	TAG_PNG_FOLDER_DATA  protocol.BINARY_TAG = 223
)

func init() {
	protocol.GetTypeRegistry().RegistryType(TAG_CONSOLE_EVENT,reflect.TypeOf((*ConsoleEvent)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(TAG_PNG_FILE,reflect.TypeOf((*PngFile)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(TAG_PNG_FOLDER_DATA,reflect.TypeOf((*PngFolderData)(nil)).Elem())
}


