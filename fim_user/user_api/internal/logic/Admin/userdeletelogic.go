package Admin

import (
	"context"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_deleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_deleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_deleteLogic {
	return &User_deleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_deleteLogic) User_delete(req *types.UserDeleteResquest) (resp *types.UserDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
