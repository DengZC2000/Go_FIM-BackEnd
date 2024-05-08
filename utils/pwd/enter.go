package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd 生成hash加密密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err.Error())
	}
	return string(hash)
}

// CheckPwd 验证密码 hash之后的密码，输入的密码，方法会比较这两者
func CheckPwd(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
