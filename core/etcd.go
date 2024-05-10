package core

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func InitEtcd(addr string) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("[etcd: %s] 连接成功\n", addr)
	return cli
}
