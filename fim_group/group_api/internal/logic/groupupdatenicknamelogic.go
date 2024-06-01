package logic

import (
	"context"

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

	return
}
