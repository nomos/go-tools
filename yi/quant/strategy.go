package quant

type Strategy interface {
	Name() string
	SetName(name string)
	SetSelf(self Strategy) error
	//Setup(mode TradeMode, exchanges ...Exchange) error
	Setup(mode TRADE_MODE, exchanges ...interface{}) error
	IsStopped() bool
	StopNow()
	TradeMode() TRADE_MODE
	SetOptions(options map[string]interface{}) error
	Run() error
	OnInit() error
	OnTick() error
	OnExit() error
}