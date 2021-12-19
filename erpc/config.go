package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/rpc"
	"github.com/nomos/go-lokas/util"
)

func init(){
	rpc.RegisterAdminFunc(LOAD_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		path:=params.String()
		err:=Instance().LoadConfig(path)
		if err != nil {
			logger.Error(err.Error())
		}
		return nil,nil
	})
	rpc.RegisterAdminFunc(GET_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		defer func() {
			if r := recover(); r != nil {
				util.Recover(r,false)
			}
		}()
		key:=params.String()
		subs:=[]string{}
		for {
			if sub:=params.StringOpt();sub!="" {
				log.Infof("subs",sub)
				subs = append(subs, sub)
			} else {
				break
			}
		}
		v:=Instance().GetConfigValue(key,subs...)
		return []byte(v),nil
	})
	rpc.RegisterAdminFunc(SET_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		key:=params.String()
		value:=params.String()
		subs:=[]string{}
		for {
			if sub:=params.StringOpt();sub!="" {
				subs = append(subs, sub)
			} else {
				break
			}
		}
		Instance().SetConfigValue(key,value,subs...)
		return nil,nil
	})
}

func (this *App) LoadConfig(name string)error{
	this.config = lox.NewAppConfig(name)
	return this.config.Load()
}

func (this *App) SubConfig(subNames ... string)lokas.IConfig{
	if util.IsNil(this.config) {
		err:=this.LoadConfig("config")
		if err != nil {
			log.Error(err.Error())
			return nil
		}
	}
	var conf lokas.IConfig=this.config
	for _,v:=range subNames {
		conf=conf.Sub(v)
	}
	return conf
}

func (this *App) GetConfigValue(key string,subNames... string)string{
	conf:=this.SubConfig(subNames...)
	return conf.GetString(key)
}

func (this *App) GetConfig()lokas.IConfig{
	return this.config
}


func (this *App) SetConfigValue(key string,value interface{},subNames... string){
	conf:=this.SubConfig(subNames...)
	conf.Set(key,value)
}