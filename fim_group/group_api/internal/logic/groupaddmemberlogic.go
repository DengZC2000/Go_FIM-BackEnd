package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
