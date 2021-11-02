//this is a generate file,edit implement on this file only!
package erpc

import (
	"github.com/nomos/go-lokas"
)

func NewConsoleEvent(text string)*ConsoleEvent{
	ret:=&ConsoleEvent{
		Text: text,
	}
	return ret
}

func (this *ConsoleEvent) OnAdd(e lokas.IEntity, r lokas.IRuntime) {
	
}

func (this *ConsoleEvent) OnRemove(e lokas.IEntity, r lokas.IRuntime) {
	
}

func (this *ConsoleEvent) OnCreate(r lokas.IRuntime) {
	
}

func (this *ConsoleEvent) OnDestroy(r lokas.IRuntime) {
	
}