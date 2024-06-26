// Code generated by goctl. DO NOT EDIT.
package types

type AddGroupRequest struct {
	UserID               uint                  `header:"User-ID"`
	GroupID              uint                  `json:"group_id"`
	AdditionalMessages   string                `json:"additional_messages,optional"`   //验证消息
	VerificationQuestion *VerificationQuestion `json:"verification_question,optional"` //问题和答案
}

type AddGroupResponse struct {
}

type GroupAddMemberRequest struct {
	UserID       uint   `header:"User-ID"`      //自己的id
	ID           uint   `json:"id"`             //群id
	MemberIDList []uint `json:"member_id_list"` //被操作的用户id
}

type GroupAddMemberResponse struct {
}

type GroupChatRequest struct {
	UserID uint `header:"User-ID"`
}

type GroupChatResponse struct {
}

type GroupCreateRequest struct {
	UserID     uint   `header:"User-ID"`
	Mode       int8   `json:"mode,optional"`
	Title      string `json:"title,optional"` //群聊名字
	IsSearch   bool   `json:"is_search,optional"`
	Size       int    `json:"size,optional"`
	UserIDList []uint `json:"user_id_list,optional"`
}

type GroupCreateResponse struct {
}

type GroupDeleteRequest struct {
	UserID uint `header:"User-ID"`
	ID     uint `path:"id"` //群id
}

type GroupDeleteResponse struct {
}

type GroupFriendsListRequest struct {
	UserID uint `header:"User-ID"` //自己的id
	ID     uint `form:"id"`        //群id
}

type GroupFriendsListResponse struct {
	List  []GroupFriendsResponse `json:"list"`
	Count int                    `json:"count"`
}

type GroupFriendsResponse struct {
	UserID    uint   `json:"user_id"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	IsInGroup bool   `json:"is_in_group"` //是否在群里面
}

type GroupHistoryDeleteRequest struct {
	UserID    uint   `header:"User-ID"`
	ID        uint   `path:"id"`
	Page      int    `form:"page,optional"`
	Limit     int    `form:"limit,optional"`
	MsgIDList []uint `json:"msg_id_list"`
}

type GroupHistoryDeleteResponse struct {
}

type GroupHistoryListResponse struct {
}

type GroupHistoryRequest struct {
	UserID uint `header:"User-ID"`
	ID     uint `path:"id"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type GroupInfoRequest struct {
	UserID uint `header:"User-ID"`
	ID     uint `path:"id"` //群id
}

type GroupInfoResponse struct {
	GroupID           uint       `json:"group_id"`            //群id
	Title             string     `json:"title"`               //群名称
	Abstract          string     `json:"abstract"`            //群简介
	Avatar            string     `json:"avatar"`              //群头像
	Creator           UserInfo   `json:"creator"`             //群主
	MemberCount       int        `json:"member_count"`        //群聊用户总数
	MemberOnlineCount int        `json:"member_online_count"` //在线用户数量
	AdminList         []UserInfo `json:"admin_list"`          //管理员列表
	Role              int8       `json:"role"`                //角色
	IsProhibition     bool       `json:"is_prohibition"`      //是否开启全员禁言
	ProhibitionTime   *int       `json:"prohibition_time"`    //自己的禁言时间，单位 min
}

type GroupListInfoResponse struct {
	ID                uint       `json:"id"`
	CreatedAt         string     `json:"created_at"`
	Title             string     `json:"title"`
	Abstract          string     `json:"abstract"`
	Avatar            string     `json:"avatar"`
	Creater           UserInfo   `json:"creater"`            // 群主
	AdminList         []UserInfo `json:"admin_list"`         // 管理员列表
	MessageCount      int        `json:"message_count"`      // 消息数量
	MemberCount       int        `json:"member_count"`       // 成员数量
	MemberOnlineCount int        `json:"member_online_ount"` // 成员在线数量
}

type GroupListRemoveRequest struct {
	IdList []uint `json:"id_list"`
}

type GroupListRemoveResponse struct {
}

type GroupListRequest struct {
	Page  int    `form:"page,optional"`
	Limit int    `form:"limit,optional"`
	Key   string `form:"key,optional"`
}

type GroupListResponse struct {
	List  []GroupListInfoResponse `json:"list"`
	Count int                     `json:"count"`
}

type GroupMemberInfo struct {
	UserID         uint   `json:"user_id"`
	UserNickname   string `json:"user_nickname"`
	Avatar         string `json:"avatar"`
	IsOnline       bool   `json:"is_online"`
	Role           int8   `json:"role"`
	MemberNickname string `json:"member_nickname"`
	CreatedAt      string `json:"created_at"`
	NewMsgDate     string `json:"new_msg_date"`
}

type GroupMemberRequest struct {
	UserID uint   `header:"User-ID"`
	ID     uint   `form:"id"`
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
	Sort   string `form:"sort,optional"`
}

type GroupMemberResponse struct {
	List  []GroupMemberInfo `json:"list"`
	Count int               `json:"count"`
}

type GroupMessageListRequest struct {
	ID    uint `path:"id"`
	Page  int  `form:"page,optional"`
	Limit int  `form:"limit,optional"`
}

type GroupMessageListResponse struct {
}

type GroupMessageRemoveRequest struct {
	IdList []uint `json:"id_list"`
}

type GroupMessageRemoveResponse struct {
}

type GroupMineListResponse struct {
	List  []GroupMineResponse `json:"list"`
	Count int                 `json:"count"`
}

type GroupMineRequest struct {
	UserID uint `header:"User-ID"`
	Mode   int8 `json:"mode,optional"` //1 表示我创建的 2 表示我加入的
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type GroupMineResponse struct {
	GroupID          uint   `json:"group_id"`
	GroupTitle       string `json:"group_title"`
	GroupAvatar      string `json:"group_avatar"`
	GroupMemberCount int    `json:"group_member_count"`
	Role             int8   `json:"role"` //角色
	Mode             int8   `json:"mode"` //1 表示我创建的 0(不传)或者其他表示我加入的
}

type GroupRemoveMemberRequest struct {
	UserID   uint `header:"User-ID"` //自己的id
	ID       uint `form:"id"`        //群id
	MemberID uint `form:"member_id"` //被操作的用户id
}

type GroupRemoveMemberResponse struct {
}

type GroupSearchListResponse struct {
	List  []GroupSearchResponse `json:"list"`
	Count int                   `json:"count"` //群的个数，因为可能群昵称搜索之类的
}

type GroupSearchRequest struct {
	UserID uint   `header:"User-ID"` //自己的id
	Key    string `form:"key,optional"`
	Page   int    `form:"page,optional"`
	Limit  int    `form:"limit,optional"`
}

type GroupSearchResponse struct {
	GroupID         uint   `json:"group_id"`
	Title           string `json:"title"`
	Abstract        string `json:"abstract"`
	Avatar          string `json:"avatar"`
	IsInGroup       bool   `json:"is_in_group"`       //我是否在群里
	UserCount       int    `json:"user_count"`        //群用户总数
	UserOnlineCount int    `json:"user_online_count"` //群用户在线总数
}

type GroupSessionListResponse struct {
	List  []GroupSessionResponse `json:"list"`
	Count int                    `json:"count"`
}

type GroupSessionRequest struct {
	UserID uint `header:"User-ID"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type GroupSessionResponse struct {
	GroupID       uint   `json:"group_id"`
	Title         string `json:"title"`
	Avatar        string `json:"avatar"`
	NewMsgDate    string `json:"new_msg_date"`    //最新的消息时间
	NewMsgPreview string `json:"new_msg_preview"` //最新的消息内容
	IsTop         bool   `json:"is_top"`          //是否置顶
}

type GroupTopRequest struct {
	UserID  uint `header:"User-ID"`
	GroupID uint `json:"group_id"`
	IsTop   bool `json:"is_top"` // true表示置顶 false表示不置顶
}

type GroupTopResponse struct {
}

type GroupUpdateMemberNicknameRequest struct {
	UserID   uint   `header:"User-ID"` //自己的id
	ID       uint   `json:"id"`        //群id
	MemberID uint   `json:"member_id"`
	Nickname string `json:"nickname"`
}

type GroupUpdateMemberNicknameResponse struct {
}

type GroupUpdateRequest struct {
	UserID               uint                  `header:"User-ID"`
	ID                   uint                  `json:"id"` //群id
	IsSearch             *bool                 `json:"is_search,optional" conf:"is_search"`
	Avatar               *string               `json:"avatar,optional" conf:"avatar"`
	Abstract             *string               `json:"abstract,optional" conf:"abstract"`
	Title                *string               `json:"title,optional" conf:"title"`
	Verification         *int8                 `json:"verification,optional" conf:"verification"`
	IsInvite             *bool                 `json:"is_invite,optional" conf:"is_invite"`
	IsTemporarySession   *bool                 `json:"is_temporary_session,optional" conf:"is_temporary_session"`
	IsProhibition        *bool                 `json:"is_prohibition,optional" conf:"is_prohibition"`
	VerificationQuestion *VerificationQuestion `json:"verification_question,optional" conf:"verification_question"`
}

type GroupUpdateResponse struct {
}

type GroupUpdateRoleRequest struct {
	UserID   uint `header:"User-ID"` //自己的id
	ID       uint `json:"id"`        //群id
	MemberID uint `json:"member_id"`
	Role     int8 `json:"role"`
}

type GroupUpdateRoleResponse struct {
}

type GroupUpdateUserProhibitionRequest struct {
	UserID          uint `header:"User-ID"`
	GroupID         uint `json:"group_id"`
	MemberID        uint `json:"member_id"`
	ProhibitionTime *int `json:"prohibition_time,optional"`
}

type GroupUpdateUserProhibitionResponse struct {
}

type GroupUpdateValidStatusRequest struct {
	UserID  uint `header:"User-ID"`
	ValidID uint `json:"valid_id"`
	Status  int8 `json:"status"`
}

type GroupUpdateValidStatusResponse struct {
}

type GroupValidInfoResponse struct {
	ID                   uint                  `json:"id"`
	GroupID              uint                  `json:"group_id"`
	UserID               uint                  `json:"user_id"`
	UserNickname         string                `json:"user_nickname"`
	UserAvatar           string                `json:"userAvatar"`
	Status               int8                  `json:"status"`
	AdditionalMessages   string                `json:"additional_messages"`
	VerificationQuestion *VerificationQuestion `json:"verification_question"`
	Title                string                `json:"title"`
	CreatedAt            string                `json:"created_at"`
	Type                 int8                  `json:"type"` //1 加群 2退群
}

type GroupValidListRequest struct {
	UserID uint `header:"User-ID"`
	Page   int  `form:"page,optional"`
	Limit  int  `form:"limit,optional"`
}

type GroupValidListResponse struct {
	List  []GroupValidInfoResponse `json:"list"`
	Count int                      `json:"count"`
}

type GroupValidRequest struct {
	UserID  uint `header:"User-ID"`
	GroupID uint `form:"group_id"`
}

type GroupValidResponse struct {
	Verification         int8                 `json:"verification"`          //好友验证,0 不允许任何人 1允许任何人 2 需要验证消息 3 需要回答问题 4 需要正确回答设置的问题
	VerificationQuestion VerificationQuestion `json:"verification_question"` //这是问题和答案，记得答案就别返回了
}

type UserInfo struct {
	UserID   uint   `json:"user_id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type VerificationQuestion struct {
	Problem1 *string `json:"problem1,optional" conf:"problem1"`
	Problem2 *string `json:"problem2,optional" conf:"problem2"`
	Problem3 *string `json:"problem3,optional" conf:"problem3"`
	Answer1  *string `json:"answer1,optional" conf:"answer1"`
	Answer2  *string `json:"answer2,optional" conf:"answer2"`
	Answer3  *string `json:"answer3,optional" conf:"answer3"`
}
