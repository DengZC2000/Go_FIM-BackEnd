// Code generated by goctl. DO NOT EDIT.
package types

type ChatAdminHistoryRemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ChatAdminHistoryRemoveResponse struct {
}

type ChatAdminHistoryRequest struct {
	SendUserID uint `form:"send_user_id"`
	RevUserID  uint `form:"rev_user_id"`
	Page       int  `form:"page,optional"`
	Limit      int  `form:"limit,optional"`
}

type ChatAdminHistoryResponse struct {
}

type ChatAdminSessionRequest struct {
	RevUserId uint `form:"rev_user_id"`
}

type ChatAdminSessionResponse struct {
	List  []UserInfo `json:"list"`
	Count int        `json:"count"`
}

type ChatDeleteRequest struct {
	UserID uint   `header:"User-ID"`
	IdList []uint `json:"id_list"`
}

type ChatDeleteResponse struct {
}

type ChatHistoryRequest struct {
	UserID   uint `header:"User-ID"`
	Page     int  `form:"page,optional"`
	Limit    int  `form:"limit,optional"`
	FriendID uint `form:"friend_id"` //好友id
}

type ChatHistoryResponse struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nick_name"`
	CreateAt string `json:"create_at"`
}

type ChatRequest struct {
	UserID uint `header:"User-ID"`
}

type ChatResponse struct {
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

type UserInfo struct {
	UserID   uint   `json:"user_id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type UserTopRequest struct {
	UserID   uint `header:"User-ID"`
	FriendID uint `json:"friend_id"`
}

type UserTopResponse struct {
}
