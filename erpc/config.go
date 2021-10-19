package erpc

import (
	"github.com/nomos/go-lokas"
	"github.com/nomos/go-lokas/cmds"
	"github.com/nomos/go-lokas/lox"
)

func init(){
	registerFunc(LOAD_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		path:=cmd.ParamsValue().String()
		Instance().LoadConfig(path)
		return nil,nil
	})
	registerFunc(GET_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		key:=cmd.ParamsValue().String()
		subs:=[]string{}
		for {
			if sub:=cmd.ParamsValue().StringOpt();sub!="" {
				subs = append(subs, sub)
			} else {
				break
			}
		}
		v:=Instance().GetConfig(key,subs...)
		return []byte(v),nil
	})
	registerFunc(SET_CONF, func(cmd *lox.AdminCommand, params *cmds.ParamsValue) ([]byte, error) {
		key:=cmd.ParamsValue().String()
		value:=cmd.ParamsValue().String()
		subs:=[]string{}
		for {
			if sub:=cmd.ParamsValue().StringOpt();sub!="" {
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
	var conf lokas.IConfig=this.config
	for _,v:=range subNames {
		conf=conf.Sub(v)
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