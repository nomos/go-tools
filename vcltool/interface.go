package vcltool

import (
	"github.com/nomos/go-events"
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/ying32/govcl/vcl"
)

type IConfigAble interface {
	Config() lokas.IConfig
	SheetName() string
	SetConfig(config lokas.IConfig)
	SetContent(s string, data interface{})
	GetContent(s string) interface{}
	setSheetName(s string)
	setContainer(IPageContainer)
}

type IFrame interface {
	IConfigAble
	Name() string
	OnCreate()
	OnDestroy()
	SetParent(vcl.IWinControl)
	SetEventEmitter(emmiter events.EventEmmiter)
	SetLogger(logger log.ILogger)
	Free()
	SetIndex(num int)
	GetIndex() int
}

type IPageContainer interface {
	IsFrameSelected(frame IFrame) bool
	On(evt events.EventName, listener ...events.Listener)
	Emit(evt events.EventName, args ...interface{})
}