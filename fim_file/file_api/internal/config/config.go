package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Etcd  string
	Mysql struct {
		DataSource string
	}
	FileSize  float64
	WriteList []string //图片文件名白名单
	BlackList []string //文件上传的黑名单
	UpLoadDir string   //文件上传保存的目录
	UserRpc   zrpc.RpcClientConf
}
