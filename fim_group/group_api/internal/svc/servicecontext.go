package svc

import (
	"FIM/common/zprc_interceptor"
	"FIM/core"
	"FIM/fim_group/group_api/internal/config"
	"FIM/fim_group/group_api/internal/middleware"
	"FIM/fim_group/group_rpc/groups"
	"FIM/fim_group/group_rpc/types/group_rpc"
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
		UserRpc:         users.NewUsers(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zprc_interceptor.ClientInfoInterceptor))),
		GroupRpc:        groups.NewGroups(zrpc.MustNewClient(c.GroupRpc, zrpc.WithUnaryClientInterceptor(zprc_interceptor.ClientInfoInterceptor))),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
