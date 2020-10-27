package quant

import (
	"github.com/nomos/go-lokas/protocol"
	"time"
)




// OrderType 委托类型
type ORDER_TYPE protocol.Enum

const (
	ORDER_TYPE_MARKET      ORDER_TYPE = iota // 市价单
	ORDER_TYPE_LIMIT                         // 限价单
	ORDER_TYPE_STOP_MARKET                   // 市价止损单
	ORDER_TYPE_STOP_LIMIT                    // 限价止损单
)

func (t ORDER_TYPE) String() string {
	switch t {
	case ORDER_TYPE_MARKET:
		return "Market"
	case ORDER_TYPE_LIMIT:
		return "Limit"
	case ORDER_TYPE_STOP_MARKET:
		return "StopMarket"
	case ORDER_TYPE_STOP_LIMIT:
		return "StopLimit"
	default:
		return "None"
	}
}

// OrderStatus 委托状态
type ORDER_STATUS protocol.Enum

const (
	ORDER_STATUS_CREATED          ORDER_STATUS = iota // 创建委托
	ORDER_STATUS_REJECTED                             // 委托被拒绝
	ORDER_STATUS_NEW                                  // 委托待成交
	ORDER_STATUS_PARTIALLY_FILLED                     // 委托部分成交
	ORDER_STATUS_FILLED                               // 委托完全成交
	ORDER_STATUS_CANCEL_PENDING                       // 委托取消
	ORDER_STATUS_CANCELLED                            // 委托被取消
	ORDER_STATUS_UNTRIGGERED                          // 等待触发条件委托单
	ORDER_STATUS_TRIGGERED                            // 已触发条件单
)

func (s ORDER_STATUS) String() string {
	switch s {
	case ORDER_STATUS_CREATED:
		return "Created"
	case ORDER_STATUS_REJECTED:
		return "Rejected"
	case ORDER_STATUS_NEW:
		return "New"
	case ORDER_STATUS_PARTIALLY_FILLED:
		return "PartiallyFilled"
	case ORDER_STATUS_FILLED:
		return "Filled"
	case ORDER_STATUS_CANCEL_PENDING:
		return "CancelPending"
	case ORDER_STATUS_CANCELLED:
		return "Cancelled"
	case ORDER_STATUS_UNTRIGGERED:
		return "Untriggered"
	case ORDER_STATUS_TRIGGERED:
		return "Triggered"
	default:
		return "None"
	}
}

type Item struct {
	protocol.Serializable
	Price  float64
	Amount float64
}

type OrderBook struct {
	protocol.Serializable
	Symbol string
	Time   time.Time
	Asks   []Item
	Bids   []Item
}

// Ask 卖一
func (o *OrderBook) Ask() (result Item) {
	if len(o.Asks) > 0 {
		result = o.Asks[0]
	}
	return
}

// Bid 买一
func (o *OrderBook) Bid() (result Item) {
	if len(o.Bids) > 0 {
		result = o.Bids[0]
	}
	return
}

// AskPrice 卖一价
func (o *OrderBook) AskPrice() (result float64) {
	if len(o.Asks) > 0 {
		result = o.Asks[0].Price
	}
	return
}

type Order struct {
	protocol.Serializable
	Id string
	ClientId string	//exchange only
	Symbol string
	Created bool
	Time time.Time
	Price float64
	StopPx float64
	Amount float64
	FilledAmount float64
	Commission   float64  // 支付的佣金
	Direction DIRECTION
}

type Balance struct {
	Equity        float64 // 净值
	Available     float64 // 可用余额
	Margin        float64 // 已用保证金
	RealizedPnl   float64
	UnrealisedPnl float64
}

type OrderParameter struct {
	Stop bool // 是否是触发委托
}

// 订单选项
type OrderOption func(p *OrderParameter)

// 触发委托选项
func OrderStopOption(stop bool) OrderOption {
	return func(p *OrderParameter) {
		p.Stop = stop
	}
}

func ParseOrderParameter(opts ...OrderOption) *OrderParameter {
	p := &OrderParameter{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

type PlaceOrderParameter struct {
	BasePrice  float64
	StopPx     float64
	PostOnly   bool
	ReduceOnly bool
	PriceType  string
	ClientOId  string
}

// 订单选项
type PlaceOrderOption func(p *PlaceOrderParameter)

// 基础价格选项(如: bybit 需要提供此参数)
func OrderBasePriceOption(basePrice float64) PlaceOrderOption {
	return func(p *PlaceOrderParameter) {
		p.BasePrice = basePrice
	}
}

// 触发价格选项
func OrderStopPxOption(stopPx float64) PlaceOrderOption {
	return func(p *PlaceOrderParameter) {
		p.StopPx = stopPx
	}
}

// 被动委托选项
func OrderPostOnlyOption(postOnly bool) PlaceOrderOption {
	return func(p *PlaceOrderParameter) {
		p.PostOnly = postOnly
	}
}

// 只减仓选项
func OrderReduceOnlyOption(reduceOnly bool) PlaceOrderOption {
	return func(p *PlaceOrderParameter) {
		p.ReduceOnly = reduceOnly
	}
}

// OrderPriceType 选项
func OrderPriceTypeOption(priceType string) PlaceOrderOption {
	return func(p *PlaceOrderParameter) {
		p.PriceType = priceType
	}
}

func OrderClientOIdOption(clientOId string) PlaceOrderOption {
	return func(p *PlaceOrderParameter) {
		p.ClientOId = clientOId
	}
}

func ParsePlaceOrderParameter(opts ...PlaceOrderOption) *PlaceOrderParameter {
	p := &PlaceOrderParameter{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
