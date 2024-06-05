package logic

import (
	"FIM/fim_group/group_models"
	"context"
	"errors"
	"fmt"
	"time"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_user_prohibitionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_user_prohibitionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_user_prohibitionLogic {
	return &Group_user_prohibitionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_user_prohibitionLogic) Group_user_prohibition(req *types.GroupUpdateUserProhibitionRequest) (resp *types.GroupUpdateUserProhibitionResponse, err error) {

	var CurrentMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&CurrentMember, "user_id = ? and group_id = ?", req.UserID, req.GroupID).Error
	if err != nil {
		return nil, errors.New("该群不存在或者你不在群中")
	}
	if !(CurrentMember.Role == 1 || CurrentMember.Role == 2) {
		return nil, errors.New("你只是普通用户，没有权限")
	}
	var member group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&member, "user_id = ? and group_id = ?", req.MemberID, req.GroupID).Error
	if err != nil {
		return nil, errors.New("操作的用户不在该群")
	}
	if req.UserID == req.MemberID {
		return nil, errors.New("自己不能禁言自己")
	}

	if !((CurrentMember.Role == 1 && (member.Role == 2 || member.Role == 3)) || (CurrentMember.Role == 2 && member.Role == 3)) {
		return nil, errors.New("角色错误")
	}
	l.svcCtx.DB.Model(&member).Update("prohibition_time", req.ProhibitionTime)

	// 利用redis的过期时间去做这个动态禁言时间
	key := fmt.Sprintf("prohibition__%d", member.ID)
	if req.ProhibitionTime != nil {
		// 给redis设置一个key，过期时间是xxxx
		l.svcCtx.Redis.Set(context.Background(), key, "1", time.Duration(*req.ProhibitionTime)*time.Minute)
	} else {
		l.svcCtx.Redis.Del(context.Background(), key)
	}
	return
}
