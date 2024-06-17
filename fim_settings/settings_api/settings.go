package main

import (
	"FIM/common/etcd"
	"FIM/common/middleware"
	"FIM/fim_settings/settings_api/internal/service"
	"flag"
	"fmt"

	"FIM/fim_settings/settings_api/internal/config"
	"FIM/fim_settings/settings_api/internal/handler"
	"FIM/fim_settings/settings_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/settings.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	// 查一下网站配置是不是有且仅有一条记录
	service.InitSettings(ctx)

	handler.RegisterHandlers(server, ctx)
	server.Use(middleware.LogMiddleware)
	etcd.DeliveryAddress(c.Etcd, c.Name+"_api", fmt.Sprintf("%s:%d", c.Host, c.Port))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
