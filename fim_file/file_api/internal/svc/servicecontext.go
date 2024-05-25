package svc

import (
	"FIM/core"
	"FIM/fim_file/file_api/internal/config"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/fim_user/user_rpc/users"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	UserRpc user_rpc.UsersClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		DB:      mysqlDb,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}
