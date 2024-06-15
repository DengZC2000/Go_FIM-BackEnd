package Admin

import (
	"context"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_message_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_message_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_message_listLogic {
	return &Group_message_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_message_listLogic) Group_message_list(req *types.GroupMessageListRequest) (resp *types.GroupMessageListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
