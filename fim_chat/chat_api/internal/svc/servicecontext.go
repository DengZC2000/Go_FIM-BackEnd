package svc

import (
	"FIM/core"
	"FIM/fim_chat/chat_api/internal/config"
	"FIM/fim_file/file_rpc/files"
	"FIM/fim_file/file_rpc/types/file_rpc"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/fim_user/user_rpc/users"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"log"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	UserRpc user_rpc.UsersClient
	FileRpc file_rpc.FilesClient
	Redis   *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisDb, err := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB, c.Redis.PoolSize)
	if err != nil {
		log.Println("redis连接失败")
	}
	return &ServiceContext{
		Config:  c,
		DB:      mysqlDb,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		FileRpc: files.NewFiles(zrpc.MustNewClient(c.FileRpc)),
		Redis:   redisDb,
	}
}
