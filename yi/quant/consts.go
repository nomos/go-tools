package quant

import "github.com/nomos/go-lokas/protocol"

// Direction 委托/持仓方向
type DIRECTION protocol.Enum

const (
	DIRECTION_BUY        DIRECTION = iota // 做多
	DIRECTION_SELL                        // 做空
	DIRECTION_CLOSE_BUY                   // 平多
	DIRECTION_CLOSE_SELL                  // 平空
)

func (this DIRECTION) String() string {
	switch this {
	case DIRECTION_BUY:
		return "Buy"
	case DIRECTION_SELL:
		return "Sell"
	default:
		return "None"
	}
}

// TradeMode 策略模式
type TRADE_MODE protocol.Enum

const (
	TRADE_MODE_BACK_TEST     TRADE_MODE = iota //回测交易
	TRADE_MODE_PAPER_TRADING                   //测试交易
	TRADE_MODE_LIVE_TRADING                    //线上交易
)

func (this TRADE_MODE) String() string {
	switch this {
	case TRADE_MODE_BACK_TEST:
		return "Backtest"
	case TRADE_MODE_PAPER_TRADING:
		return "PaperTrading"
	case TRADE_MODE_LIVE_TRADING:
		return "LiveTrading"
	default:
		return "None"
	}
}
