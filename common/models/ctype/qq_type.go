package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type QQType struct {
	Enable   bool   `json:"enable"` //是否启用
	AppID    string `json:"app_id"`
	Key      string `json:"key"`
	Redirect string `json:"redirect"`
	WebPath  string `json:"web_path"` // 跳转的地址，但是不存,用↓下面这个方法算出来的
}

func (qq *QQType) GetPath() string {
	if qq.Key == "" || qq.AppID == "" || qq.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", qq.AppID, qq.Redirect)
}

// Scan 取出来的时候的数据
func (c *QQType) Scan(val interface{}) error {
	err := json.Unmarshal(val.([]byte), c)
	if err != nil {
		return err
	}

	return nil
}

// Value 入库的数据
func (c *QQType) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
