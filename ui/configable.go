package ui

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/events"
)


type ConfigAble struct {
	content   map[string]interface{}
	conf      lokas.IConfig
	log       log.ILogger
	sheetName string
	listener  events.EventEmmiter
	container IPageContainer
	index     int
}

func (this *ConfigAble) Options()[]FrameOption {
	return []FrameOption{WithConfig(this.conf),WithLogger(this.log),WithListener(this.listener)}
}

func (this *ConfigAble) GetListener()events.EventEmmiter{
	return this.listener
}

func (this *ConfigAble) SetListener(listener events.EventEmmiter){
	this.listener = listener
}

func (this *ConfigAble) SetContent(s string, data interface{}) {
	if this.content == nil {
		this.content = map[string]interface{}{}
	}
	this.content[s] = data
}

func (this *ConfigAble) GetContent(s string) interface{} {
	if this.content == nil {
		this.content = map[string]interface{}{}
	}
	return this.content[s]
}

func (this *ConfigAble) setContainer(container IPageContainer) {
	this.container = container
}

func (this *ConfigAble) Container()IPageContainer{
	return this.container
}

func (this *ConfigAble) SetEventEmitter(listener events.EventEmmiter) {
	this.listener = listener
}

func (this *ConfigAble) SetIndex(num int) {
	this.index = num
}

func (this *ConfigAble) GetIndex() int {
	return this.index
}

func (this *ConfigAble) SetLogger(log log.ILogger) {
	this.log = log
}

func (this *ConfigAble) SetConfig(config lokas.IConfig) {
	this.conf = config
}

func (this *ConfigAble) GetLogger() log.ILogger {
	return this.log
}

func (this *ConfigAble) Config() lokas.IConfig {
	return this.conf
}

func (this *ConfigAble) setSheetName(s string) {
	this.sheetName = s
}

func (this *ConfigAble) SheetName() string {
	return this.sheetName
}
