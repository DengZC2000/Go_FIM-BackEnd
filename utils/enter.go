package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
)

func InList(list []string, key string) bool {
	for _, val := range list {
		if val == key {
			return true
		}
	}
	return false
}
func InListRegex(list []string, key string) bool {

	for _, val := range list {
		regex, err := regexp.Compile(val)
		if err != nil {
			logx.Error(err)
			return false
		}
		if regex.MatchString(key) {
			return true
		}
	}
	return false
}
