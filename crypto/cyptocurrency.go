package crypto

import (
	"github.com/HuobiRDCenter/huobi_Golang/config"
	"github.com/HuobiRDCenter/huobi_Golang/pkg/client"
	"github.com/nomos/go-lokas/util/log"
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