package logic

import (
	"FIM/fim_group/group_models"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_update_roleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_update_roleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_update_roleLogic {
	return &Group_update_roleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_update_roleLogic) Group_update_role(req *types.GroupUpdateRoleRequest) (resp *types.GroupUpdateRoleResponse, err error) {
	//先查该群是否存在，自己是否在群，
	var CurrentMember group_models.GroupMemberModel
	var UpdateMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&CurrentMember, "group_id = ? and user_id = ?", req.ID, req.UserID).Error
	if err != nil {
		return nil, errors.New("该群不存在或者你不是群成员")
	}
	err = l.svcCtx.DB.Take(&UpdateMember, "group_id = ? and user_id = ?", req.ID, req.MemberID).Error
	if err != nil {
		return nil, errors.New("该用户不是群成员")
	}

	if CurrentMember.Role != 1 {
		return nil, errors.New("你不是群主，无法修改他人角色")
	}
	if req.Role == 1 {
		return nil, errors.New("无法赋予群主权限")
	}
	if UpdateMember.Role == int(req.Role) {
		//已经是就别改了
		return
	}
	l.svcCtx.DB.Model(&UpdateMember).Update("role", req.Role)
	return
}
