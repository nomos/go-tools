package crypto

import (
	"github.com/nomos/go-log/log"
	"github.com/nomos/huobi/config"
	"github.com/nomos/huobi/pkg/client"
	"time"
)

func InitHuobiApi() {
	client := new(client.MarketClient).Init(config.Host)
	for {
		time.Sleep(time.Second*5)
		resp, err := client.GetLatestTrade("ethusdt")
		if err != nil {
			log.Error(err.Error())
		} else {
			log.Warnf("price",resp.Data[0].Price)
		}
	}
}