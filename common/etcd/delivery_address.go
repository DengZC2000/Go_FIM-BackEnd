package etcd

import (
	"FIM/core"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/netx"
	"strings"
)

// DeliveryAddress 上送服务地址
func DeliveryAddress(etcdAddr, serviceName, addr string) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		logx.Errorf("地址错误 %s", addr)
		return
	}
	if list[0] == "0.0.0.0" {
		ip := netx.InternalIp()
		addr = strings.ReplaceAll(addr, "0.0.0.0", ip)
	}
	client := core.InitEtcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, addr)
	if err != nil {
		logx.Errorf("地址上送失败 %s", err.Error())
		return
	}
	logx.Infof("地址上送成功 %s %s", serviceName, addr)
}

// GetServiceAddr 获取服务地址
func GetServiceAddr(etcdAddr, serviceName string) (addr string) {
	client := core.InitEtcd(etcdAddr)
	res, err := client.Get(context.Background(), serviceName)
	if err == nil && len(res.Kvs) >= 1 {
		return string(res.Kvs[0].Value)
	}
	return ""
}
