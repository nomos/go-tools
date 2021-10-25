package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/log"
	"github.com/nomos/go-lokas/lox"
	"github.com/nomos/go-lokas/util"
)

func init(){
	RegisterAdminFunc(LOAD_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
		path:=params.String()
		Instance().LoadConfig(path)
		return nil,nil
	})
	RegisterAdminFunc(GET_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
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
		v:=Instance().GetConfig(key,subs...)
		return []byte(v),nil
	})
	RegisterAdminFunc(SET_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue,logger log.ILogger) ([]byte, error) {
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
		Instance().SetConfig(key,value,subs...)
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
		log.Infof(conf,v,conf.(*lox.AppConfig).Viper)
	}
	return conf
}

func (this *App) GetConfig(key string,subNames... string)string{
	conf:=this.SubConfig(subNames...)
	return conf.GetString(key)
}

func (this *App) SetConfig(key string,value string,subNames... string){
	conf:=this.SubConfig(subNames...)
	conf.Set(key,value)
}