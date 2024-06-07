package main

import (
	"FIM/common/zprc_interceptor"
	"flag"
	"fmt"

	"FIM/fim_group/group_rpc/internal/config"
	"FIM/fim_group/group_rpc/internal/server"
	"FIM/fim_group/group_rpc/internal/svc"
	"FIM/fim_group/group_rpc/types/group_rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/grouprpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		group_rpc.RegisterGroupsServer(grpcServer, server.NewGroupsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// rpc透传
	s.AddUnaryInterceptors(zprc_interceptor.ServerUnaryInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
