package exchanges

import (
	"github.com/nomos/go-log/log"
	"github.com/nomos/go-lokas/util"
	"github.com/nomos/go-tools/yi/data"
	"github.com/nomos/go-tools/yi/quant"
	"github.com/nomos/huobi/pkg/client/marketwebsocketclient"
	"github.com/nomos/huobi/pkg/model/market"
	"time"
)

type HuobiExchange struct {
	accessKey    string
	secretKey    string
	host         string
	proxy        string
	candleClient *marketwebsocketclient.CandlestickWebSocketClient
}

var _ quant.ISpotExchange = (*HuobiExchange)(nil)

func (this *HuobiExchange) Init(accessKey string, secretKey string, host string, proxy ...string) *HuobiExchange {
	this.accessKey = accessKey
	this.secretKey = secretKey
	this.host = host
	if len(proxy) > 0 {
		this.proxy = proxy[0]
	}
	this.candleClient = new(marketwebsocketclient.CandlestickWebSocketClient).Init(this.host, this.proxy)
	this.candleClient.SetHandler(func() {
		this.candleClient.Request("btcusdt", "1min", 1569361140, 1569366420, "2305")
		this.candleClient.Subscribe("btcusdt", "1min", "2305")
	}, func(response interface{}) {
		resp, ok := response.(market.SubscribeCandlestickResponse)
		if ok {
			if &resp != nil {
				if resp.Tick != nil {
					t := resp.Tick
					log.Infof("time: %s,count: %d,amount: %v,  vol: %v [%v-%v-%v-%v]",
						util.FormatTimeToString(time.Unix(t.Id, 0)), t.Count, t.Amount,
						t.Vol, t.Open, t.Close, t.Low, t.High)
				}

				if resp.Data != nil {
					log.Infof("WebSocket returned data, count=%d", len(resp.Data))
					for _, t := range resp.Data {
						log.Infof("time: %s, count: %d,amount: %v, vol: %v [%v-%v-%v-%v]",
							util.FormatTimeToString(time.Unix(t.Id, 0)), t.Count, t.Amount, t.Vol, t.Open, t.Close, t.Low, t.High)
					}
				}
			}
		} else {
			log.Errorf("Unknown response: %v", resp)
		}
	})
	this.candleClient.Connect(true)
	return this
}

func (this *HuobiExchange) GetName() (name string) {
	return "Huobi"
}

func (this *HuobiExchange) GetTime() (tm int64, err error) {
	return time.Now().Unix(), nil
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
