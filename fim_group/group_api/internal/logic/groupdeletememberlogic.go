package logic

import (
	"FIM/fim_group/group_models"
	"context"
	"errors"
	"gorm.io/gorm"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_delete_memberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_delete_memberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_delete_memberLogic {
	return &Group_delete_memberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_delete_memberLogic) Group_delete_member(req *types.GroupRemoveMemberRequest) (resp *types.GroupRemoveMemberResponse, err error) {

	var CurrentMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&CurrentMember, "user_id = ? and group_id = ?", req.UserID, req.ID).Error
	if err != nil {
		return nil, errors.New("该群不存在或者你不在群中")
	}

	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "user_id = ? and group_id = ?", req.MemberID, req.ID).Error
	if err != nil {
		return nil, errors.New("操作的用户不在该群")
	}
	if CurrentMember.Role == 3 {
		//普通用户
		return nil, errors.New("你没有权限操作")
	}
	if req.UserID == req.MemberID {
		return nil, errors.New("自己不能踢自己")
	}
	//是群主
	if CurrentMember.Role == 1 {
		//群主可以肆无忌惮的踢人
		err = DeleteMember(l.svcCtx.DB, req.MemberID)
		if err != nil {
			return nil, err
		}
	} else {
		//否则就是管理员了
		if member.Role == 1 || member.Role == 2 {
			return nil, errors.New("你没有权限操作群主或其他管理员")
		}
		err = DeleteMember(l.svcCtx.DB, req.MemberID)
		if err != nil {
			return nil, err
		}

	}
	return
}
func DeleteMember(db *gorm.DB, ID uint) (err error) {
	err = db.Delete(&group_models.GroupMemberModel{}, "user_id = ?", ID).Error
	return
}
