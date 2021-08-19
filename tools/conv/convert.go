package conv

import (
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/protocol"
)

type ENC_TYPE protocol.Enum

const (
	ENC_STRING ENC_TYPE = iota
	ENC_UNICODE
	ENC_RUNE
	ENC_NUMBER
)

var _ protocol.IEnum = (*ENC_TYPE)(nil)

var ALL_ENC_TYPES = []protocol.IEnum{ENC_STRING,ENC_UNICODE,ENC_RUNE,ENC_NUMBER}

func (this ENC_TYPE) ToString() string {
	switch this {
	case ENC_STRING:
		return "string"
	case ENC_UNICODE:
		return "unicode"
	case ENC_RUNE:
		return "rune"
	case ENC_NUMBER:
		return "number"
	default:
		log.Panic("type not supported")
	}
	return ""
}

func (this ENC_TYPE) Enum() protocol.Enum {
	return protocol.Enum(this)
}

