package group_models

import (
	"FIM/common/models"
	"FIM/common/models/ctype"
)

// GroupVerifyModel 群验证消息
type GroupVerifyModel struct {
	models.Model
	GroupID              uint                        `json:"group_id"` //群id
	GroupModel           GroupModel                  `gorm:"foreignKey:GroupID" json:"-"`
	UserID               uint                        `json:"user_id"`                            //需要加群或者是退群的用户id
	Status               int8                        `json:"status"`                             //状态 0 未操作 1 同意 2 拒绝 3 忽略
	AdditionalMessages   string                      `gorm:"size:32" json:"additional_messages"` // 附加消息
	VerificationQuestion *ctype.VerificationQuestion `json:"verification_question"`              // 验证问题 为3和4的时候需要
	Type                 int8                        `json:"type"`                               // 类型 1 加群 2 退群
}
