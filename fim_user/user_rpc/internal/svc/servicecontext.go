package svc

import (
	"FIM/core"
	"FIM/fim_user/user_rpc/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	MyRedis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisDb, err := core.InitRedis(c.MyRedis.Addr, c.MyRedis.Password, c.MyRedis.DB, c.MyRedis.PoolSize)
	if err != nil {
		log.Println("redis连接失败")
	}
	return &ServiceContext{
		Config:  c,
		DB:      mysqlDb,
		MyRedis: redisDb,
	}
}
