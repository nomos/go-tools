package huobicli

import (
	"sync"
)

var _once sync.Once
var _host string
var _proxy string

var _http *HttpClient
var _ws *WsClient

func SetHost(host string){
	_host = host
}

func SetProxy(proxy string){
	_proxy = proxy
}

func CreateHttp(host string,proxy... string)*HttpClient{
	return nil
}

func CreateWs(host string,proxy... string)*WsClient {
	return nil
}