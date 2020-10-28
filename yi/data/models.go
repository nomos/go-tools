package data

import (
	"github.com/nomos/go-lokas/protocol"
	"github.com/shopspring/decimal"
	"time"
)

type Market struct {
	Symbol string
}

type Record struct {
	protocol.Serializable
	Timestamp time.Time
	Amount    decimal.Decimal
	Vol       decimal.Decimal
	Open      decimal.Decimal
	Close     decimal.Decimal
	Low       decimal.Decimal
	High      decimal.Decimal
}

type DayRecord struct {
	protocol.Serializable
	Timestamp     time.Time
	Amount        decimal.Decimal
	Vol           decimal.Decimal
	Open          decimal.Decimal
	Close         decimal.Decimal
	Low           decimal.Decimal
	High          decimal.Decimal
	MinuteRecords []*Record
}
