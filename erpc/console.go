package erpc

import (
	"github.com/nomos/go-lokas/protocol"
)

type ConsoleEvent struct {
	Text string
}

func (this *ConsoleEvent) GetId() (protocol.BINARY_TAG, error) {
	return TAG_ConsoleEvent,nil
}

func (this *ConsoleEvent) Serializable() protocol.ISerializable {
	return this
}

func newConsoleEvent(text string)*ConsoleEvent{
	return &ConsoleEvent{Text: text}
}
