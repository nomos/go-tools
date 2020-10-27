package exchanges

import (
	"github.com/nomos/go-tools/yi/data"
	"github.com/nomos/go-tools/yi/quant"
	"time"
)

type HuobiExchange struct {

}

var _ quant.ISpotExchange = (*HuobiExchange)(nil)

func (this *HuobiExchange) GetName() (name string) {
	return "Huobi"
}

func (this *HuobiExchange) GetTime() (tm int64, err error) {
	return time.Now().Unix(),nil
}

func (this *HuobiExchange) GetBalance(currency string) (result *quant.SpotBalance, err error) {
	panic("implement me")
}

func (this *HuobiExchange) GetOrderBook(symbol string, depth int) (result *quant.OrderBook, err error) {
	panic("implement me")
}

func (this *HuobiExchange) GetRecords(symbol string, period data.PERIOD, from int64, end int64, limit int) (records []*data.Record, err error) {
	panic("implement me")
}

func (this *HuobiExchange) Buy(symbol string, orderType quant.ORDER_TYPE, price float64, size float64) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HuobiExchange) Sell(symbol string, orderType quant.ORDER_TYPE, price float64, size float64) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HuobiExchange) PlaceOrder(symbol string, direction quant.DIRECTION, orderType quant.ORDER_TYPE, price float64, size float64, opts ...quant.PlaceOrderOption) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HuobiExchange) GetOpenOrders(symbol string, opts ...quant.OrderOption) (result []*quant.Order, err error) {
	panic("implement me")
}

func (this *HuobiExchange) GetHistoryOrders(symbol string, opts ...quant.OrderOption) (result []*quant.Order, err error) {
	panic("implement me")
}

func (this *HuobiExchange) GetOrder(symbol string, id string, opts ...quant.OrderOption) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HuobiExchange) CancelAllOrders(symbol string, opts ...quant.OrderOption) (err error) {
	panic("implement me")
}

func (this *HuobiExchange) CancelOrder(symbol string, id string, opts ...quant.OrderOption) (result *quant.Order, err error) {
	panic("implement me")
}

func (this *HuobiExchange) IO(name string, params string) (string, error) {
	panic("implement me")
}