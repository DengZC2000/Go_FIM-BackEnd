package user_models

import (
	"FIM/common/models"
	"gorm.io/gorm"
)

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserID     uint      `json:"send_user_id"` //发起验证方
	SendUserModel  UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	RevUserID      uint      `json:"rev_user_id"` //接收验证方
	RevUserModel   UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	SendUserNotice string    `gorm:"size:128" json:"send_user_notice"` //备注
	RevUserNotice  string    `gorm:"size:128" json:"rev_user_notice"`  //备注
	/*
	   A --> B
	   SendUserNotice 是A对B的备注
	   RevUserNotice  是B对A的备注
	*/
}

func (f *FriendModel) IsFriend(db *gorm.DB, A, B uint) bool {
	err := db.Take(&f, "(send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)", A, B, B, A).Error
	if err != nil {
		return false
	}
	return true
}
func (f *FriendModel) Friends(db *gorm.DB, MyUserID uint) (list []FriendModel) {
	db.Find(&list, "(send_user_id = ? or rev_user_id = ?) ", MyUserID, MyUserID)
	return
}
func (f *FriendModel) FriendNotice(UserID uint) string {
	if UserID == f.SendUserID {
		return f.SendUserNotice
	}
	if UserID == f.RevUserID {
		return f.RevUserNotice
	}
	return ""
}
