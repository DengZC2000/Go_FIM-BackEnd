package logic

import (
	"FIM/fim_group/group_models"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_topLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_topLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_topLogic {
	return &Group_topLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_topLogic) Group_top(req *types.GroupTopRequest) (resp *types.GroupTopResponse, err error) {
	var CurrentMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Take(&CurrentMember, "group_id = ? and user_id = ?", req.GroupID, req.UserID).Error
	if err != nil {
		return nil, errors.New("该群不存在或者你不是群成员")
	}

	var userTop group_models.GroupUserTopModel
	err1 := l.svcCtx.DB.Take(&userTop, "user_id = ? and group_id = ?", req.UserID, req.GroupID).Error
	if err1 != nil {
		//没置顶
		if req.IsTop {
			l.svcCtx.DB.Create(&group_models.GroupUserTopModel{
				GroupID: req.GroupID,
				UserID:  req.UserID,
			})
		}
	} else {
		l.svcCtx.DB.Delete(&userTop)
	}
	return
}
