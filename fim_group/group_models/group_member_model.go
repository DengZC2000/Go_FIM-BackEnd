package group_models

import "FIM/common/models"

type GroupMemberModel struct {
	models.Model
	GroupID         uint       `json:"group_id"` //群id
	GroupModel      GroupModel `gorm:"foreignKey:GroupID" json:"-"`
	UserID          uint       `json:"user_id"`                        //用户id
	MemberNickname  string     `gorm:"size:32" json:"member_nickname"` //群昵称
	Role            int        `json:"role"`                           //1 群主 2 管理员 3 普通用户
	ProhibitionTime *int       `json:"prohibition_time"`               //禁言时间，单位 min
}
