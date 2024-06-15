package Admin

import (
	"context"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_listLogic {
	return &Group_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_listLogic) Group_list(req *types.GroupListRequest) (resp *types.GroupListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
