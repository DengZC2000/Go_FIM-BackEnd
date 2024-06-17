package svc

import (
	"FIM/core"
	"FIM/fim_settings/settings_api/internal/config"
	"FIM/fim_settings/settings_api/internal/middleware"
	"gorm.io/gorm"
	"net/http"
)

type ServiceContext struct {
	Config          config.Config
	DB              *gorm.DB
	AdminMiddleware func(next http.HandlerFunc) http.HandlerFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		DB:              mysqlDb,
		AdminMiddleware: middleware.NewAdminMiddleware().Handle,
	}
}
