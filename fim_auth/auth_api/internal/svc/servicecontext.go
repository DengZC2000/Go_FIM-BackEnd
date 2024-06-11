package svc

import (
	"FIM/common/log_stash"
	"FIM/common/zprc_interceptor"
	"FIM/core"
	"FIM/fim_auth/auth_api/internal/config"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/fim_user/user_rpc/users"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"log"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	Redis          *redis.Client
	UserRpc        user_rpc.UsersClient
	KqPusherClient *kq.Pusher
	ActionPusher   *log_stash.Pusher
	RuntimePusher  *log_stash.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	redisDb, err := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB, c.Redis.PoolSize)
	if err != nil {
		log.Println("redis连接失败")
	}
	Kq := kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic)
	return &ServiceContext{
		Config:         c,
		DB:             mysqlDb,
		Redis:          redisDb,
		UserRpc:        users.NewUsers(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zprc_interceptor.ClientInfoInterceptor))),
		KqPusherClient: Kq,
		ActionPusher:   log_stash.NewActionPusher(Kq, c.Name),
		RuntimePusher:  log_stash.NewRuntimePusher(Kq, c.Name),
	}
}
