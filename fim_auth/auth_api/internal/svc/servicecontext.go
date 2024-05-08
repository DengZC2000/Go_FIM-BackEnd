package svc

import (
	"FIM/core"
	"FIM/fim_auth/auth_api/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisDb, err := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB, c.Redis.PoolSize)
	if err != nil {
		log.Println("redis连接失败")
	}
	//mysqlDb.AutoMigrate(&auth_models.UserModel{})
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
		Redis:  redisDb,
	}
}
