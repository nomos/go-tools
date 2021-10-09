package ui

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/util/events"
	"github.com/ying32/govcl/vcl"
)

type FrameOption func(frame IFrame)

func WithConfig(conf lokas.IConfig) FrameOption {
	return func(frame IFrame) {
		frame.SetConfig(conf)
	}
}

func WithContent(s string, data interface{}) FrameOption {
	return func(frame IFrame) {
		frame.SetContent(s, data)
	}
}

func WithListener(listener events.EventEmmiter) FrameOption{
	return func(frame IFrame) {
		frame.SetListener(listener)
	}
}

func WithParent(parent vcl.IWinControl)FrameOption{
	return func(frame IFrame) {
		frame.SetParent(parent)
	}
}

func WithLogger(logger log.ILogger)FrameOption{
	return func(frame IFrame) {
		frame.SetLogger(logger)
	}
}