package logic

import (
	"FIM/fim_group/group_models"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_update_nicknameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_update_nicknameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_update_nicknameLogic {
	return &Group_update_nicknameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_update_nicknameLogic) Group_update_nickname(req *types.GroupUpdateMemberNicknameRequest) (resp *types.GroupUpdateMemberNicknameResponse, err error) {
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
	if req.UserID == req.MemberID {
		//如果是自己修改自己
		l.svcCtx.DB.Model(&CurrentMember).Updates(map[string]any{
			"member_nickname": req.Nickname,
		})
	}
	//剩下的都是修改别的人的
	if CurrentMember.Role == 3 {
		return nil, errors.New("你没有权限修改他的群昵称")
	}
	if CurrentMember.Role == 2 && UpdateMember.Role != 3 {
		//管理员只能修改普通用户的
		return nil, errors.New("你不能修改群主或者其他管理员的群昵称")
	}
	//剩下的都是合法的
	l.svcCtx.DB.Model(&UpdateMember).Updates(map[string]any{
		"member_nickname": req.Nickname,
	})
	return
}
