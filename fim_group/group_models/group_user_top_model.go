package group_models

import "FIM/common/models"

// GroupUserTopModel 用户置顶群聊表
type GroupUserTopModel struct {
	models.Model
	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
}
