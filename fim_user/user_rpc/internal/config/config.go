package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	MyRedis struct {
		Addr     string
		Password string
		DB       int
		PoolSize int
	}
}
