package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/common/models/ctype"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_historyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_historyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_historyLogic {
	return &Group_historyLogic{
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

func (l *Group_historyLogic) Group_history(req *types.GroupHistoryRequest) (resp *HistoryListResponse, err error) {
	var CurrentMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&CurrentMember, "group_id = ? and user_id = ?", req.ID, req.UserID).Error
	if err != nil {
		return nil, errors.New("该群不存在或者你不是群成员")
	}
	// 去查我删除了哪些聊天记录,id
	msgIDList := make([]uint, 0)
	l.svcCtx.DB.Model(group_models.GroupUserMsgDeleteModel{}).
		Where("group_id = ? and user_id = ?", req.ID, req.UserID).
		Select("msg_id").
		Scan(&msgIDList)
	msgIDList = append(msgIDList, 0) //这一句其实很重要，要不然就是id not in null ，然后结果就是什么都查不出来

	groupMsgList, count, _ := list_query.ListQuery(l.svcCtx.DB, group_models.GroupMsgModel{GroupID: req.ID}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			//Sort:  "created_at desc",
		},

		Where: l.svcCtx.DB.Where("id not in ?", msgIDList),
	})
	var memberMap = map[uint]string{}
	var memberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&memberList, "group_id = ?", req.ID)
	for _, info := range memberList {
		memberMap[info.UserID] = info.MemberNickname
	}

	var userIDList []uint32
	for _, model := range groupMsgList {
		userIDList = append(userIDList, uint32(model.SendUserID))
	}
	//去重,主要是减轻rpc复杂度
	userIDList = utils.DeduplicationList(userIDList)
	userListResponse, err1 := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	var list = make([]HistoryResponse, 0)
	for _, model := range groupMsgList {
		info := HistoryResponse{
			ID:             model.ID,
			UserID:         model.SendUserID,
			Msg:            model.Msg,
			MsgType:        model.MsgType,
			MsgPreview:     model.MsgPreview,
			CreatedAt:      model.CreatedAt,
			MemberNickname: memberMap[model.SendUserID],
		}

		if err1 == nil {
			info.UserNickname = userListResponse.UserInfo[uint32(info.UserID)].NickName
			info.UserAvatar = userListResponse.UserInfo[uint32(info.UserID)].Avatar
		}
		if req.UserID == info.UserID {
			info.IsMe = true
		}
		list = append(list, info)
	}
	resp = &HistoryListResponse{}
	resp.List = list
	resp.Count = int(count)
	return
}
