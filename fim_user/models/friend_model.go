package models

import (
	"FIM/common/models"
)

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserID    uint      `json:"send_user_id"` //发起验证方
	SendUserModel UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	RevUserID     uint      `json:"rev_user_id"` //接收验证方
	RevUserModel  UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	Notice        string    `gorm:"size:128" json:"notice"` //备注

}
