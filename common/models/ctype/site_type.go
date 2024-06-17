package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

type SiteType struct {
	CreatedAt   string `json:"created_at"`
	BeiAn       string `json:"bei_an"`
	Version     string `json:"version"`
	QQImage     string `json:"qq_image"`
	WechatImage string `json:"wechat_image"`
	BiliBiliUrl string `json:"bili_bili_url"`
	GiteeUrl    string `json:"gitee_url"`
	GithubUrl   string `json:"github_url"`
}

// Scan 取出来的时候的数据
func (c *SiteType) Scan(val interface{}) error {
	err := json.Unmarshal(val.([]byte), c)
	if err != nil {
		return err
	}

	return nil
}

// Value 入库的数据
func (c *SiteType) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
