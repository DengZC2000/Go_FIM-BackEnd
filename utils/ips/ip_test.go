package ips

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ip)
}
