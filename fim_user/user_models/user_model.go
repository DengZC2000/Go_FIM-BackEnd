package user_models

import "FIM/common/models"

// UserModel 用户表
type UserModel struct {
	models.Model
	NickName       string         `gorm:"size:64" json:"nick_name"`
	Password       string         `gorm:"size:128" json:"-"`
	Profile        string         `gorm:"size:128" json:"profile"` //个人简介
	Avatar         string         `gorm:"size:256" json:"avatar"`  //头像
	IP             string         `gorm:"size:32" json:"ips"`
	Address        string         `gorm:"size:64" json:"address"`
	Role           int8           `json:"role"`                           //1 管理员 2 普通用户
	OpenID         string         `gorm:"size:128" json:"-"`              //第三方平台登陆的凭证，唯一
	RegisterSource string         `gorm:"size:16" json:"register_source"` //注册来源 qq 。。
	UserConfModel  *UserConfModel `gorm:"foreignKey:UserID" json:"UserConfModel"`
}
