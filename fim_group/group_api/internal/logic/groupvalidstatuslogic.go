package logic

import (
	"FIM/fim_group/group_models"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_valid_statusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_valid_statusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_valid_statusLogic {
	return &Group_valid_statusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_valid_statusLogic) Group_valid_status(req *types.GroupUpdateValidStatusRequest) (resp *types.GroupUpdateValidStatusResponse, err error) {
	var groupValidModel group_models.GroupVerifyModel
	err = l.svcCtx.DB.Take(&groupValidModel, req.ValidID).Error
	if err != nil {
		return nil, errors.New("不存在的验证记录")
	}
	if groupValidModel.Status != 0 {
		return nil, errors.New("已经处理过该验证请求了")
	}
	//判断我有没有权限处理这个验证请求
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "user_id = ? and group_id = ?", req.UserID, groupValidModel.GroupID).Error
	if err != nil {
		return nil, errors.New("没有处理该操作的权限")
	}
	if !(member.Role == 1 || member.Role == 2) {
		return nil, errors.New("没有处理该操作的权限")
	}
	groupValidModel.Status = req.Status
	switch req.Status {
	case 0: //未操作
		return
	case 1: //同意
		//将用户加到群里面去
		var member1 = group_models.GroupMemberModel{
			GroupID: groupValidModel.GroupID,
			UserID:  groupValidModel.UserID,
			Role:    3,
		}
		l.svcCtx.DB.Create(&member1)
	case 2: //拒绝
	case 3: //忽略
	}
	l.svcCtx.DB.Model(&groupValidModel).UpdateColumn("status", req.Status)
	return
}
