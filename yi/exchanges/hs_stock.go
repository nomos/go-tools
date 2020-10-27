package exchanges

import "github.com/nomos/go-tools/yi/quant"

type HSStockExchange struct {

}


var _ quant.ISpotExchange = (*HSStockExchange)(nil)
