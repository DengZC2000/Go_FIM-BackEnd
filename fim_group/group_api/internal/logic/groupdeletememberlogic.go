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
	if req.UserID == req.MemberID {
		//自己退出群聊
		if CurrentMember.Role == 1 {
			// 群主不能直接退出群聊，只能解散群聊
			return nil, errors.New("您是群主，只能解散该群")
		}
		// 把member中的与这个用户的记录删除
		err = DeleteMember(l.svcCtx.DB, req.MemberID)
		if err != nil {
			return nil, err
		}
		// 给群验证表里面加条记录
		l.svcCtx.DB.Create(&group_models.GroupVerifyModel{
			GroupID: req.ID,
			UserID:  req.UserID,
			Type:    2, // 2 代表退群
		})
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
			return nil, errors.New("你没有权限踢出群主或其他管理员")
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
