package ui

import (
	"github.com/nomos/go-events"
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/ying32/govcl/vcl"
	"sync"
)

type ITreeSchema interface {
	String()string
	SetKey(s string)
	Node()*vcl.TTreeNode
	SetNode(node *vcl.TTreeNode)
	Image()string
	Key()string
	Value()string
	GetRootTree()[]int
	Idx()int
	ToObj()interface{}
	Collapse()bool
	SetCollapse(bool)
	SetIdx(id int)
	InnerIdx()int
	SetInnerIdx(id int)
	Insert(s ITreeSchema)ITreeSchema
	Detach(s ITreeSchema)ITreeSchema
	Root()ITreeSchema
	Parent()ITreeSchema
	AddChild(schema ITreeSchema)ITreeSchema
	Children()[]ITreeSchema
}

type ITree interface {
	sync.Locker
	UpdateTree(schema ITreeSchema)
}

type ITreeData interface {
	Key()string
	ValueString()string
	String()string
}


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
	SetEventEmitter(listener events.EventEmmiter)
	SetLogger(logger log.ILogger)
	SetListener(listener events.EventEmmiter)
	Free()
	SetIndex(num int)
	GetIndex() int
}

type IPageContainer interface {
	IsFrameSelected(frame IFrame) bool
	On(evt events.EventName, listener ...events.Listener)
	Emit(evt events.EventName, args ...interface{})
}