package logic

import (
	"context"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_ws_chatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_ws_chatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_ws_chatLogic {
	return &Group_ws_chatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_ws_chatLogic) Group_ws_chat(req *types.GroupChatRequest) (resp *types.GroupChatResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
