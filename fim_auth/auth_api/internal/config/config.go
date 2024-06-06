package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
		PoolSize int
	}
	OpenLoginList []struct {
		Name string
		Icon string
		Href string
	}
	QQ struct {
		AppID    string
		AppKey   string
		Redirect string
	}
	Etcd         string
	UserRpc      zrpc.RpcClientConf
	WhiteList    []string //白名单
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
