package Admin

import (
	"context"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_restrictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_restrictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_restrictLogic {
	return &User_restrictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_restrictLogic) User_restrict(req *types.UserRestrictResquest) (resp *types.UserRestrictResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
