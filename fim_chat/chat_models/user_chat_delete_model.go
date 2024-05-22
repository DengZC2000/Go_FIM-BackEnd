package chat_models

type UserChatDeleteModel struct {
	UserID uint `json:"user_id"`
	ChatID uint `json:"chat_id"` //聊天记录的id

}
