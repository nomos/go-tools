package huobicli

type WsClient struct {
	accessKey string
	secretKey string
	host      string
	proxy string
}

func (this *WsClient) Init(accessKey string,secretKey string,host string,proxy... string){
	this.accessKey = accessKey
	this.secretKey = secretKey
	this.host = host
	if len(proxy)>0 {
		this.proxy = proxy[0]
	}
}

func (this *WsClient) reset() {
}

func (this *WsClient) SetProxy(proxy string) {
	this.proxy = proxy
	this.reset()
}

func (this *WsClient) GetProxy()string{
	return this.proxy
}

func (this *WsClient) SetHost(host string) {
	this.host = host
	this.reset()
}

func (this *WsClient) GetHost()string{
	return this.host
}

func (this *WsClient) SetAccessKey(accessKey string) {
	this.accessKey = accessKey
	this.reset()
}

func (this *WsClient) GetAccessKey()string{
	return this.accessKey
}

func (this *WsClient) SetSecretKey(secretKey string) {
	this.secretKey = secretKey
	this.reset()
}

func (this *WsClient) GetSecretKey()string{
	return this.secretKey
}

//蜡烛

