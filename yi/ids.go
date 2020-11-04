//this is a generate file,do not edit it
package yi


import (
	"github.com/nomos/go-lokas/protocol"
	"github.com/nomos/go-tools/yi/charts"
	"github.com/nomos/go-tools/yi/data"
	"github.com/nomos/go-tools/yi/quant"
	"reflect"
)

const (
	YI_Item protocol.BINARY_TAG = 36001
	YI_OrderBook protocol.BINARY_TAG = 36002
	YI_Order protocol.BINARY_TAG = 36003
	YI_ChartData protocol.BINARY_TAG = 36004
	YI_Record protocol.BINARY_TAG = 36005
	YI_DayRecord protocol.BINARY_TAG = 36006
	YI_SpotAsset protocol.BINARY_TAG = 36007
)

func init() {
	protocol.GetTypeRegistry().RegistryType(YI_Item,reflect.TypeOf((*quant.OrderItem)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(YI_OrderBook,reflect.TypeOf((*quant.OrderBook)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(YI_Order,reflect.TypeOf((*quant.Order)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(YI_ChartData,reflect.TypeOf((*charts.ChartData)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(YI_Record,reflect.TypeOf((*data.Record)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(YI_DayRecord,reflect.TypeOf((*data.DayRecord)(nil)).Elem())
	protocol.GetTypeRegistry().RegistryType(YI_SpotAsset,reflect.TypeOf((*quant.SpotAsset)(nil)).Elem())
}
