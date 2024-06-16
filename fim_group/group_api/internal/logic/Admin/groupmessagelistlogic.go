package Admin

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/common/models/ctype"
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils"
	"context"
	"time"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_message_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_message_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_message_listLogic {
	return &Group_message_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HistoryResponse struct {
	UserID         uint       `json:"user_id"`
	UserNickname   string     `json:"user_nickname"`
	UserAvatar     string     `json:"user_avatar"`
	Msg            *ctype.Msg `json:"msg"`
	MsgPreview     string     `json:"msg_preview"`
	ID             uint       `json:"id"`
	MsgType        int8       `json:"msg_type"`
	CreatedAt      time.Time  `json:"created_at"`
	IsMe           bool       `json:"is_me"`
	MemberNickname string     `json:"member_nickname"` //群昵称，是不是有备注的该显示备注？
}
type HistoryListResponse struct {
	List  []HistoryResponse `json:"list"`
	Count int               `json:"count"`
}

func (l *Group_message_listLogic) Group_message_list(req *types.GroupMessageListRequest) (resp *HistoryListResponse, err error) {

	groupMsgList, count, _ := list_query.ListQuery(l.svcCtx.DB, group_models.GroupMsgModel{GroupID: req.ID}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
		Preloads: []string{"GroupMemberModel"},
	})

	var userIDList []uint32
	for _, model := range groupMsgList {
		userIDList = append(userIDList, uint32(model.SendUserID))
	}
	//去重,主要是减轻rpc复杂度
	userIDList = utils.DeduplicationList(userIDList)
	userListResponse, err1 := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	var list = make([]HistoryResponse, 0)
	for _, model := range groupMsgList {
		info := HistoryResponse{
			ID:         model.ID,
			UserID:     model.SendUserID,
			Msg:        model.Msg,
			MsgType:    model.MsgType,
			MsgPreview: model.MsgPreview,
			CreatedAt:  model.CreatedAt,
		}
		if model.GroupMemberModel != nil {
			info.MemberNickname = model.GroupMemberModel.MemberNickname
		}
		if err1 == nil {
			info.UserNickname = userListResponse.UserInfo[uint32(info.UserID)].NickName
			info.UserAvatar = userListResponse.UserInfo[uint32(info.UserID)].Avatar
		}

		list = append(list, info)
	}
	resp = &HistoryListResponse{}
	resp.List = list
	resp.Count = int(count)

	return
}
