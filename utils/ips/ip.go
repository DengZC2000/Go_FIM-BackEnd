package ips

import (
	"fmt"
	"net"
	"strings"
)

// GetOutBoundIP 已经用另外一种方式替代获取IP，这个文件夹不删是因为做个参考
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
