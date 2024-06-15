package Admin

import (
	"context"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_message_removeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_message_removeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_message_removeLogic {
	return &Group_message_removeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_message_removeLogic) Group_message_remove(req *types.GroupMessageRemoveRequest) (resp *types.GroupMessageRemoveResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
