package Admin

import (
	"context"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_listLogic {
	return &User_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_listLogic) User_list(req *types.UserListResquest) (resp *types.UserListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
