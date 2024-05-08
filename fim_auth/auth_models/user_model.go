package auth_models

import "FIM/common/models"

// UserModel 用户表
type UserModel struct {
	models.Model
	NickName string `gorm:"size:64" json:"nick_name"`
	Password string `gorm:"size:32" json:"password"`
	Profile  string `gorm:"size:128" json:"profile"` //个人简介
	Avatar   string `gorm:"size:256" json:"avatar"`  //头像
	IP       string `gorm:"size:32" json:"ip"`
	Address  string `gorm:"size:64" json:"address"`
	Role     int8   `json:"role"` //1 管理员 2 普通用户
}
