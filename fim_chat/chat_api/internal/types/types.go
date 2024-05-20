// Code generated by goctl. DO NOT EDIT.
package types

type ChatHistoryRequest struct {
	UserID uint `header:"User-ID"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type ChatHistoryResponse struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nick_name"`
	CreateAt string `json:"create_at"`
}

type ChatSession struct {
	UserID     uint   `json:"user_id"`
	Avatar     string `json:"avatar"`
	Nickname   string `json:"nick_name"`
	CreateAt   string `json:"create_at"`   //消息时间
	MsgPreview string `json:"msg_preview"` //消息预览
	IsTop      bool   `json:"is_top"`      //是否置顶
}

type ChatSessionRequest struct {
	UserID uint `header:"User-ID"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
	Key    int  `form:"key,optional"`
}

type ChatSessionResponse struct {
	List  []ChatSession `json:"list"`
	Count int64         `json:"count"`
}

type UserTopRequest struct {
	UserID   uint `header:"User-ID"`
	FriendID uint `json:"friend_id"`
}

type UserTopResponse struct {
}
