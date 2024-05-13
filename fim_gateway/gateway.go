package main

import (
	"FIM/common/etcd"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"regexp"
	"strings"
)

//	var serviceMap = map[string]string{
//		"auth": "http://127.0.0.1:20021",
//		"user": "http://127.0.0.1:20022",
//	}
type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func FileResponse(msg string, res http.ResponseWriter) {
	response := BaseResponse{Code: 10086, Msg: msg}
	byteData, _ := json.Marshal(response)
	res.Write(byteData)
}
func auth(authAddr string, res http.ResponseWriter, req *http.Request) bool {
	authRequest, _ := http.NewRequest("POST", authAddr, nil)
	authRequest.Header = req.Header
	authRequest.Header.Set("valid_path", req.URL.Path)
	authRes, err := http.DefaultClient.Do(authRequest)
	if err != nil {
		FileResponse("认证服务错误", res)
		return false
	}

	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data *struct {
			UserID uint `json:"user_id"`
			Role   int  `json:"role"`
		} `json:"data"`
	}
	var authResponse Response
	byteData, _ := io.ReadAll(authRes.Body)
	authErr := json.Unmarshal(byteData, &authResponse)
	if authErr != nil {
		logx.Error(authErr)
		FileResponse("认证服务错误", res)
		return false
	}
	//认证不通过
	if authResponse.Code != 0 {
		res.Write(byteData)
		return false
	}
	if authResponse.Data != nil {
		req.Header.Set("User-ID", fmt.Sprintf("%d", authResponse.Data.UserID))
		req.Header.Set("User-Role", fmt.Sprintf("%d", authResponse.Data.Role))
	}

	return true
}
func proxy(proxyAddr string, res http.ResponseWriter, req *http.Request) {
	byteData, _ := io.ReadAll(req.Body)
	proxyReq, _ := http.NewRequest(req.Method, proxyAddr, bytes.NewBuffer(byteData))

	proxyReq.Header = req.Header
	proxyReq.Header.Del("valid_path")
	response, ProxyErr := http.DefaultClient.Do(proxyReq)
	if ProxyErr != nil {
		logx.Error(ProxyErr.Error())
		FileResponse("服务异常", res)
		return
	}
	io.Copy(res, response.Body)
}
func gateway(res http.ResponseWriter, req *http.Request) {
	// 匹配请求前缀 /api/auth/xx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		FileResponse("请求前缀错误", res)
		return
	}

	service := addrList[1]
	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		logx.Errorf("不匹配的服务 %s", service)
		FileResponse("不匹配的服务", res)
		return
	}
	remoteAddr := strings.Split(req.RemoteAddr, ":")

	// 请求认证服务地址
	authAddr := etcd.GetServiceAddr(config.Etcd, "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)
	proxyUrl := fmt.Sprintf("http://%s%s", addr, req.URL.String())
	//打印日志
	logx.Infof("%s %s", remoteAddr[0], proxyUrl)

	if !auth(authUrl, res, req) {
		return
	}

	proxy(proxyUrl, res, req)

	//fmt.Println(req.URL.Path)
}

var configFile = flag.String("f", "settings.yaml", "the config file")

type Config struct {
	Addr string
	Etcd string
	Log  logx.LogConf
}

var config Config

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &config)
	logx.SetUp(config.Log)
	//回调函数
	http.HandleFunc("/", gateway)

	fmt.Printf("gateway running %s\n", config.Addr)

	//绑定服务
	http.ListenAndServe(config.Addr, nil)
}
