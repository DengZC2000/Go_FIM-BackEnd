package main

import (
	"FIM/common/etcd"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"net/http"
	"regexp"
	"strings"
)

//var serviceMap = map[string]string{
//	"auth": "http://127.0.0.1:20021",
//	"user": "http://127.0.0.1:20022",
//}

func gateway(res http.ResponseWriter, req *http.Request) {
	// 匹配请求前缀 /api/auth/xx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		res.Write([]byte("err"))
		return
	}
	service := addrList[1]

	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		fmt.Printf("不匹配的服务%s", service)
		res.Write([]byte("err"))
		return
	}

	remoteAddr := strings.Split(req.RemoteAddr, ":")
	url := fmt.Sprintf("http://%s%s", addr, req.URL.String())
	proxyReq, _ := http.NewRequest(req.Method, url, req.Body)
	proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	response, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		fmt.Println(err.Error())
		res.Write([]byte("服务异常"))
		return
	}
	io.Copy(res, response.Body)

	//fmt.Println(req.URL.Path)
}

var configFile = flag.String("f", "settings.yaml", "the config file")

type Config struct {
	Addr string
	Etcd string
}

var config Config

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &config)

	//回调函数
	http.HandleFunc("/", gateway)

	fmt.Printf("gateway running %s\n", config.Addr)

	//绑定服务
	http.ListenAndServe(config.Addr, nil)
}
