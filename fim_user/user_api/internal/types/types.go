// Code generated by goctl. DO NOT EDIT.
package types

type AddFriendRequest struct {
	UserID               uint                  `header:"User-ID"`
	FriendID             uint                  `json:"friend_id"`
	AdditionalMessages   string                `json:"additional_messages,optional"`
	VerificationQuestion *VerificationQuestion `json:"verification_question,optional"` //问题和答案
}

type AddFriendResponse struct {
}

type DeleteFriendRequest struct {
	UserID   uint `header:"User-ID"`
	FriendID uint `json:"friend_id"`
}

type DeleteFriendResponse struct {
}

type FriendInfoRequest struct {
	UserID   uint `header:"User-ID"`
	Role     int8 `header:"User-Role"`
	FriendID uint `form:"friend_id"`
}

type FriendInfoResponse struct {
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nick_name"`
	Profile  string `json:"profile"` //���˼���
	Avatar   string `json:"avatar"`  //ͷ��
	Notice   string `json:"notice"`
}

type FriendListRequest struct {
	UserID uint `header:"User-ID"`
	Role   int8 `header:"User-Role"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type FriendListResponse struct {
	List  []FriendInfoResponse `json:"list"`
	Count int                  `json:"count"`
}

type FriendNoticeUpdateRequest struct {
	UserID   uint   `header:"User-ID"`
	FriendID uint   `json:"friend_id"`
	Notice   string `json:"notice"`
}

type FriendNoticeUpdateResponse struct {
}

type FriendValidInfo struct {
	Nickname             string                `json:"nick_name"`
	UserID               uint                  `json:"user_id"`
	Avatar               string                `json:"avatar"`
	AdditionalMessages   string                `json:"additional_messages,optional"`
	VerificationQuestion *VerificationQuestion `json:"verification_question,optional"` //问题和答案
	Status               int8                  `json:"status"`                         //0 未操作 1 同意 2 拒绝 3 忽略
	Verification         int8                  `json:"verification"`                   //好友验证,0 不允许任何人 1允许任何人 2 需要验证消息 3 需要回答问题 4 需要正确回答设置的问题
	ID                   uint                  `json:"id"`                             //验证记录的id
	Flag                 string                `json:"flag"`                           //send 我是发送方，rev我是接收方
}

type FriendValidResponse struct {
	List  []FriendValidInfo `json:"list"`
	Count int64             `json:"count"`
}

type FriendValidResquest struct {
	UserID uint `header:"User-ID"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type FriendValidStatusRequest struct {
	UserID   uint `header:"User-ID"`
	VerifyID uint `json:"verify_id"`
	Status   int8 `json:"status"` //状态
}

type FriendValidStatusResponse struct {
}

type SearchInfo struct {
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nick_name"`
	Profile  string `json:"profile"`   //���˼���
	Avatar   string `json:"avatar"`    //ͷ��
	IsFriend bool   `json:"is_friend"` //是否已经是好友
}

type SearchRequest struct {
	UserID uint   `header:"User-ID"`
	Key    string `form:"key,optional"`    //用户id或昵称
	Online bool   `form:"online,optional"` //搜索在线的用户
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
}

type SearchResponse struct {
	List  []SearchInfo `json:"list"`
	Count int64        `json:"count"`
}

type UserInfoRequest struct {
	UserID uint `header:"User-ID"`
	Role   int8 `header:"User-Role"`
}

type UserInfoResponse struct {
	UserID               uint                  `json:"user_id"`
	Nickname             string                `json:"nick_name"`
	Profile              string                `json:"profile"` //���˼���
	Avatar               string                `json:"avatar"`  //ͷ��
	RecallMessage        *string               `json:"recall_message"`
	FriendOnline         bool                  `json:"friend_online"`
	Sound                bool                  `json:"sound"`
	SecureLink           bool                  `json:"secure_link"`
	SavePwd              bool                  `json:"save_pwd"`
	SearchUser           int8                  `json:"search_user"`
	Verification         int8                  `json:"verification"`
	VerificationQuestion *VerificationQuestion `json:"verification_question"`
	Online               bool                  `json:"online"`
}

type UserInfoUpdateRequest struct {
	UserID               uint                  `header:"User-ID"`
	Nickname             *string               `json:"nick_name,optional" user:"nickname"`
	Profile              *string               `json:"profile,optional" user:"profile"` //���˼���
	Avatar               *string               `json:"avatar,optional" user:"avatar"`   //ͷ��
	RecallMessage        *string               `json:"recall_message,optional" user_conf:"recall_message"`
	FriendOnline         *bool                 `json:"friend_online,optional" user_conf:"friend_online"`
	Sound                *bool                 `json:"sound,optional" user_conf:"sound"`
	SecureLink           *bool                 `json:"secure_link,optional" user_conf:"secure_link"`
	SavePwd              *bool                 `json:"save_pwd,optional" user_conf:"save_pwd"`
	SearchUser           *int8                 `json:"search_user,optional" user_conf:"search_user"`
	Verification         *int8                 `json:"verification,optional" user_conf:"verification"`
	VerificationQuestion *VerificationQuestion `json:"verification_question,optional" user_conf:"verification_question"`
}

type UserInfoUpdateResponse struct {
}

type UserValidRequest struct {
	UserID   uint `header:"User-ID"`
	FriendID uint `json:"friend_id"`
}

type UserValidResponse struct {
	Verification         int8                 `json:"verification"`          //好友验证,0 不允许任何人 1允许任何人 2 需要验证消息 3 需要回答问题 4 需要正确回答设置的问题
	VerificationQuestion VerificationQuestion `json:"verification_question"` //这是问题和答案，记得答案就别返回了
}

type VerificationQuestion struct {
	Problem1 *string `json:"problem1,optional" user_conf:"problem1"`
	Problem2 *string `json:"problem2,optional" user_conf:"problem2"`
	Problem3 *string `json:"problem3,optional" user_conf:"problem3"`
	Answer1  *string `json:"answer1,optional" user_conf:"answer1"`
	Answer2  *string `json:"answer2,optional" user_conf:"answer2"`
	Answer3  *string `json:"answer3,optional" user_conf:"answer3"`
}
