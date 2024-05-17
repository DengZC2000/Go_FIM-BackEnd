package user_models

import (
	"FIM/common/models"
	"FIM/common/models/ctype"
)

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Model
	SendUserID           uint                        `json:"send_user_id"` //发起验证方
	SendUserModel        UserModel                   `gorm:"foreignKey:SendUserID" json:"-"`
	RevUserID            uint                        `json:"rev_user_id"` //接收验证方
	RevUserModel         UserModel                   `gorm:"foreignKey:RevUserID" json:"-"`
	Notice               string                      `json:"notice"`                              //备注
	Status               int8                        `json:"status"`                              //0 未操作 1 同意 2 拒绝 3 忽略 4删除
	AdditionalMessages   string                      `gorm:"size:128" json:"additional_messages"` //附加消息
	VerificationQuestion *ctype.VerificationQuestion `json:"friend_question"`                     //验证问题,不一定需要，验证方式为3和4的时候需要

}
