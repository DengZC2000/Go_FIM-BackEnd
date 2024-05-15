// Code generated by goctl. DO NOT EDIT.
package types

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

type SearchInfo struct {
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nick_name"`
	Profile  string `json:"profile"`   //���˼���
	Avatar   string `json:"avatar"`    //ͷ��
	IsFriend bool   `json:"is_friend"` //是否已经是好友
}

type SearchRequest struct {
	UserID uint   `header:"User-ID"`
	Key    string `form:"key"`    //用户id或昵称
	Online bool   `form:"online"` //搜索在线的用户
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

type VerificationQuestion struct {
	Problem1 *string `json:"problem1,optional" user_conf:"problem1"`
	Problem2 *string `json:"problem2,optional" user_conf:"problem2"`
	Problem3 *string `json:"problem3,optional" user_conf:"problem3"`
	Answer1  *string `json:"answer1,optional" user_conf:"answer1"`
	Answer2  *string `json:"answer2,optional" user_conf:"answer2"`
	Answer3  *string `json:"answer3,optional" user_conf:"answer3"`
}
