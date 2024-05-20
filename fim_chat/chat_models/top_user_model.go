package chat_models

// TopUserModel 置顶用户表
type TopUserModel struct {
	UserID    uint `json:"user_id"`
	TopUserID uint `json:"top_user_id"`
}
