package pwd

import (
	"fmt"
	"testing"
)

func TestPwdHash(t *testing.T) {
	hash := HashPwd("123456")
	fmt.Println(hash)
}
