package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"FIM/fim_group/group_models"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type Group_sessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_sessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_sessionLogic {
	return &Group_sessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type SessionData struct {
	GroupID       uint   `gorm:"column:group_id"`
	NewMsgDate    string `gorm:"column:newMsgDate"`    //最新的消息时间
	NewMsgPreview string `gorm:"column:newMsgPreview"` //最新的消息内容
	IsTop         bool   `gorm:"column:isTop"`
}

func (l *Group_sessionLogic) Group_session(req *types.GroupSessionRequest) (resp *types.GroupSessionListResponse, err error) {
	// 先查我有哪些群
	var userGroupIDList []uint
	l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
		Where("user_id = ?", req.UserID).
		Select("group_id").
		Scan(&userGroupIDList)
	column := fmt.Sprintf("(if((select 1 from group_user_top_models where user_id = %d and group_user_top_models.group_id = group_msg_models.group_id),1,0)) as isTop", req.UserID)

	//查哪些聊天记录是被删掉的
	var msgDeleteIDList []uint
	l.svcCtx.DB.Model(group_models.GroupUserMsgDeleteModel{}).Where("group_id in ?", userGroupIDList).Select("msg_id").Scan(&msgDeleteIDList)
	sessionList, count, _ := list_query.ListQuery(l.svcCtx.DB, SessionData{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "isTop desc, newMsgDate desc",
		},
		Table: func() (string, any) {
			return "(?) as u", l.svcCtx.DB.Model(&group_models.GroupMsgModel{}).
				Select("group_id",
					"max(created_at) as newMsgDate",
					column,
					"(select msg_preview from group_msg_models as g where g.group_id = group_msg_models.group_id order by g.created_at desc limit 1) as newMsgPreview").
				Where("group_id in (?) and id not in (?)", userGroupIDList, msgDeleteIDList).
				Group("group_id")
		},
	})
	resp = &types.GroupSessionListResponse{}
	var groupIDList []uint
	for _, data := range sessionList {
		groupIDList = append(groupIDList, data.GroupID)
	}
	var groupListModel []group_models.GroupModel
	l.svcCtx.DB.Find(&groupListModel, groupIDList)
	var groupMap = map[uint]group_models.GroupModel{}
	for _, model := range groupListModel {
		groupMap[model.ID] = model
	}
	for _, data := range sessionList {
		resp.List = append(resp.List, types.GroupSessionResponse{
			GroupID:       data.GroupID,
			Title:         groupMap[data.GroupID].Title,
			Avatar:        groupMap[data.GroupID].Avatar,
			NewMsgDate:    data.NewMsgDate,
			NewMsgPreview: data.NewMsgPreview,
			IsTop:         data.IsTop,
		})
	}
	resp.Count = int(count)
	return
}
