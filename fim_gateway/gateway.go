package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var serviceMap = map[string]string{
	"auth": "http://127.0.0.1:20021",
	"user": "http://127.0.0.1:20022",
}

func gateway(res http.ResponseWriter, req *http.Request) {
	// 匹配请求前缀 /api/user/xx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		res.Write([]byte("err"))
		return
	}
	service := addrList[1]
	addr, ok := serviceMap[service]
	if !ok {
		fmt.Printf("不匹配的服务%s", service)
		res.Write([]byte("err"))
		return
	}
	remoteAddr := strings.Split(req.RemoteAddr, ":")
	url := fmt.Sprintf("%s%s", addr, req.URL.String())
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
}

func main() {
	flag.Parse()

	var c Config
	conf.MustLoad(*configFile, &c)

	//回调函数
	http.HandleFunc("/", gateway)

	fmt.Printf("gateway running %s\n", c.Addr)

	//绑定服务
	http.ListenAndServe(c.Addr, nil)
}
