package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
	"strings"
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
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func GetFilePrefix(fileName string) string {
	nameList := strings.Split(fileName, ".")
	var prefix string
	for i := 0; i < len(nameList)-1; i++ {
		if i == len(nameList)-2 {
			prefix += nameList[i]
			continue
		} else {
			prefix += nameList[i] + "."
		}
	}
	return prefix
}

// DeduplicationList 去重
func DeduplicationList[T string | int | uint | uint32](req []T) (response []T) {
	Map := make(map[T]bool)
	for _, val := range req {
		if !Map[val] {
			Map[val] = true
		}
	}
	response = make([]T, 0)
	for key, _ := range Map {
		response = append(response, key)
	}
	return
}
