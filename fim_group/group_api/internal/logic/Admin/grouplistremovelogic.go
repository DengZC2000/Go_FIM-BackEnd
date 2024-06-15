package Admin

import (
	"context"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_list_removeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_list_removeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_list_removeLogic {
	return &Group_list_removeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_list_removeLogic) Group_list_remove(req *types.GroupListRemoveRequest) (resp *types.GroupListRemoveResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
