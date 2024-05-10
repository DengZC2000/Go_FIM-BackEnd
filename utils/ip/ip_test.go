package ip

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	ip := GetIP()
	fmt.Println(ip)
}
