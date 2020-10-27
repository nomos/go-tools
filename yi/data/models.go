package data

import (
	"github.com/nomos/go-lokas/protocol"
	"time"
)

type Market struct {
	Symbol string
}

type Record struct {
	protocol.Serializable
	Symbol    string
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}