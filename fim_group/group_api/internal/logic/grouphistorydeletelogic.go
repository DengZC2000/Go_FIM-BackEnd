package logic

import (
	"FIM/fim_group/group_models"
	"FIM/utils/set"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_history_deleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_history_deleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_history_deleteLogic {
	return &Group_history_deleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_history_deleteLogic) Group_history_delete(req *types.GroupHistoryDeleteRequest) (resp *types.GroupHistoryDeleteResponse, err error) {
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
	// 做差集,之前删过的，就别再次出现去删了,保证这次删的都是没删过的
	addMsgIDList := set.Difference(req.MsgIDList, msgIDList)
	if len(addMsgIDList) == 0 {
		return
	}
	//用户传过来的消息id，消息不一定能够存在，保证入库的删除消息id合法
	var msgIDFindList []uint
	l.svcCtx.DB.Model(group_models.GroupMsgModel{}).
		Where("id in ?", addMsgIDList).Select("id").Scan(&msgIDFindList)
	if len(addMsgIDList) != len(msgIDFindList) {
		//瞎调
		return nil, errors.New("消息一致性错误")
	}

	var list []group_models.GroupUserMsgDeleteModel
	for _, id := range addMsgIDList {
		list = append(list, group_models.GroupUserMsgDeleteModel{
			GroupID: req.ID,
			UserID:  req.UserID,
			MsgID:   id,
		})
	}
	err = l.svcCtx.DB.Create(&list).Error
	return
}
