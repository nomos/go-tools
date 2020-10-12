package charts

import "time"

type IData interface {
	Date()time.Time
}

type ChartData struct {
	Data []IData
	DataOffset int32
	Period int32
	Right int32
	Symbol string
}