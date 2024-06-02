package logic

import (
	"FIM/fim_group/group_models"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_add_memberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_add_memberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_add_memberLogic {
	return &Group_add_memberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_add_memberLogic) Group_add_member(req *types.GroupAddMemberRequest) (resp *types.GroupAddMemberResponse, err error) {
	//先查该群是否存在，自己是否在群，
	var CurrentMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Preload("GroupModel").Take(&CurrentMember, "group_id = ? and user_id = ?", req.ID, req.UserID).Error
	if err != nil {
		return nil, errors.New("该群不存在或者你不是群成员")
	}
	//有没有设置群成员不能邀请新用户进群
	if CurrentMember.Role == 3 && CurrentMember.GroupModel.IsInvite == false {
		return nil, errors.New("群设置限制，现在只有群主和管理员可以邀请新用户进群")
	}
	var MemberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&MemberList, "group_id = ? and user_id in ?", req.ID, req.MemberIDList)
	if len(MemberList) > 0 {
		return nil, errors.New("已经有好友在这群里了")
	}
	var AddMemberList []group_models.GroupMemberModel
	for _, memberID := range req.MemberIDList {
		AddMemberList = append(AddMemberList, group_models.GroupMemberModel{
			Role:    3,
			GroupID: req.ID,
			UserID:  memberID,
		})
	}
	err = l.svcCtx.DB.Create(&AddMemberList).Error
	if err != nil {
		return nil, errors.New("邀请好友失败")
	}

	return
}
