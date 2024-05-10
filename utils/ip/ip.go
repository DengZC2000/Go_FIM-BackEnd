package ip

import (
	"fmt"
	"net"
)

func GetIP() (addr string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					return ipNet.IP.String() //返回第一个就行，后面就不循环了
				}
			}
		}
	}
	return
}
