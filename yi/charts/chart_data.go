package charts

import (
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-tools/yi/data"
)

type ChartData struct {
	protocol.Serializable
	Data       []data.Record
	DataOffset int32
	Period     data.PERIOD
	Right      int32
	Symbol     string
}
