package logic

import (
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"FIM/fim_group/group_models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_deleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_deleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_deleteLogic {
	return &Group_deleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_deleteLogic) Group_delete(req *types.GroupDeleteRequest) (resp *types.GroupDeleteResponse, err error) {
	//只能是群主才能调用
	var groupMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&groupMember, "group_id = ? and user_id = ?", req.ID, req.UserID).Error
	if err != nil {
		return nil, errors.New("群不存在或者此用户不是该群成员")
	}
	if groupMember.Role != 1 {
		return nil, errors.New("只有群主才能解散该群哦")
	}
	//关联的这个群的信息也要删掉
	var msgList []group_models.GroupMsgModel
	l.svcCtx.DB.Find(&msgList, "group_id = ?", req.ID).Delete(&msgList)
	//群成员也要删掉
	var memberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&memberList, "group_id = ?", req.ID).Delete(&memberList)
	//群验证消息
	var vList []group_models.GroupVerifyModel
	l.svcCtx.DB.Find(&vList, "group_id = ?", req.ID).Delete(&vList)
	//群解散
	var group group_models.GroupModel
	l.svcCtx.DB.Take(&group, req.ID).Delete(&group)
	return
}
