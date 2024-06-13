package svc

import (
	"FIM/common/zprc_interceptor"
	"FIM/core"
	"FIM/fim_logs/log_api/internal/config"
	"FIM/fim_logs/log_api/internal/middleware"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/fim_user/user_rpc/users"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	UserRpc         user_rpc.UsersClient
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		DB:              mysqlDb,
		UserRpc:         users.NewUsers(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zprc_interceptor.ClientInfoInterceptor))),
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
