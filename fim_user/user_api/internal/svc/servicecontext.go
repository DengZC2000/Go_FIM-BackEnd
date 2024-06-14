package svc

import (
	"FIM/core"
	"FIM/fim_chat/chat_rpc/chat"
	"FIM/fim_chat/chat_rpc/types/chat_rpc"
	"FIM/fim_group/group_rpc/groups"
	"FIM/fim_group/group_rpc/types/group_rpc"
	"FIM/fim_user/user_api/internal/config"
	"FIM/fim_user/user_api/internal/middleware"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/fim_user/user_rpc/users"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	Redis           *redis.Client
	UserRpc         user_rpc.UsersClient
	ChatRpc         chat_rpc.ChatClient
	GroupRpc        group_rpc.GroupsClient
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisDb, err := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB, c.Redis.PoolSize)
	if err != nil {
		log.Println("redis连接失败")
	}
	return &ServiceContext{
		Config:          c,
		DB:              mysqlDb,
		Redis:           redisDb,
		UserRpc:         users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
		ChatRpc:         chat.NewChat(zrpc.MustNewClient(c.ChatRpc)),
		GroupRpc:        groups.NewGroups(zrpc.MustNewClient(c.GroupRpc)),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
