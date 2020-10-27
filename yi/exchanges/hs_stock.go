package exchanges

import (
	"github.com/nomos/go-tools/yi/data"
	"github.com/nomos/go-tools/yi/quant"
	"time"
)

type HSStockExchange struct {

}

var _ quant.ISpotExchange = (*HSStockExchange)(nil)

func (this *HSStockExchange) GetName() (name string) {
	return "SH-SZ Stock"
}

func (this *HSStockExchange) GetTime() (tm int64, err error) {
	return time.Now().Unix(),nil
}

func (this *HSStockExchange) GetBalance(currency string) (result *quant.SpotBalance, err error) {
	panic("implement me")
}

func (this *HSStockExchange) GetOrderBook(symbol string, depth int) (result *quant.OrderBook, err error) {
	panic("implement me")
}

func (this *HSStockExchange) GetRecords(symbol string, period data.PERIOD, from int64, end int64, limit int) (records []*data.Record, err error) {
	panic("implement me")
}

func (this *HSStockExchange) Buy(symbol string, orderType quant.ORDER_TYPE, price float64, size float64) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HSStockExchange) Sell(symbol string, orderType quant.ORDER_TYPE, price float64, size float64) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HSStockExchange) PlaceOrder(symbol string, direction quant.DIRECTION, orderType quant.ORDER_TYPE, price float64, size float64, opts ...quant.PlaceOrderOption) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HSStockExchange) GetOpenOrders(symbol string, opts ...quant.OrderOption) (result []*quant.Order, err error) {
	panic("implement me")
}

func (this *HSStockExchange) GetHistoryOrders(symbol string, opts ...quant.OrderOption) (result []*quant.Order, err error) {
	panic("implement me")
}

func (this *HSStockExchange) GetOrder(symbol string, id string, opts ...quant.OrderOption) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HSStockExchange) CancelAllOrders(symbol string, opts ...quant.OrderOption) (err error) {
	panic("implement me")
}

func (this *HSStockExchange) CancelOrder(symbol string, id string, opts ...quant.OrderOption) (result *quant.Order, err error) {
	panic("implement me")
}
