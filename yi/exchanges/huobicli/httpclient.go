package huobicli

import "github.com/nomos/huobi/pkg/client"

type HttpClient struct {
	*client.AccountClient
	*client.AlgoOrderClient
	*client.CommonClient
	*client.CrossMarginClient
	*client.ETFClient
	*client.IsolatedMarginClient
	*client.MarketClient
	*client.OrderClient
	*client.StableCoinClient
	*client.SubUserClient
	*client.WalletClient

	accessKey string
	secretKey string
	host      string
	proxy string
}

func (this *HttpClient) Init(accessKey string, secretKey string, host string, proxy ...string)*HttpClient {
	this.accessKey = accessKey
	this.secretKey = secretKey
	this.host = host
	if len(proxy)>0 {
		this.proxy = proxy[0]
	}
	this.reset()
	return this
}

func (this *HttpClient) reset(){
	this.AccountClient = new(client.AccountClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.AlgoOrderClient = new(client.AlgoOrderClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.CrossMarginClient = new(client.CrossMarginClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.ETFClient = new(client.ETFClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.IsolatedMarginClient = new(client.IsolatedMarginClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.OrderClient = new(client.OrderClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.StableCoinClient = new(client.StableCoinClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.SubUserClient = new(client.SubUserClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.WalletClient = new(client.WalletClient).Init(this.accessKey, this.secretKey, this.host, this.proxy)
	this.MarketClient = new(client.MarketClient).Init(this.host, this.proxy)
	this.CommonClient = new(client.CommonClient).Init(this.host, this.proxy)
}

func (this *HttpClient) SetProxy(proxy string) {
	this.proxy = proxy
	this.reset()
}

func (this *HttpClient) GetProxy()string{
	return this.proxy
}

func (this *HttpClient) SetHost(host string) {
	this.host = host
	this.reset()
}

func (this *HttpClient) GetHost()string{
	return this.host
}

func (this *HttpClient) SetAccessKey(accessKey string) {
	this.accessKey = accessKey
	this.reset()
}

func (this *HttpClient) GetAccessKey()string{
	return this.accessKey
}

func (this *HttpClient) SetSecretKey(secretKey string) {
	this.secretKey = secretKey
	this.reset()
}

func (this *HttpClient) GetSecretKey()string{
	return this.secretKey
}