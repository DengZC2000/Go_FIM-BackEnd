package group_models

import "FIM/common/models"

type GroupUserMsgDeleteModel struct {
	models.Model
	UserID  uint `json:"user_id"`
	MsgID   uint `json:"msg_id"`
	GroupID uint `json:"group_id"`
}
