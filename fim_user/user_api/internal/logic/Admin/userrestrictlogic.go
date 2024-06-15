package Admin

import (
	"FIM/fim_user/user_models"
	"context"
	"errors"

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
	var user user_models.UserModel
	err = l.svcCtx.DB.Preload("UserConfModel").Take(&user, req.UserID).Error
	if err != nil {
		// 没找到
		return nil, errors.New("用户不存在")
	}
	l.svcCtx.DB.Model(&user.UserConfModel).Updates(map[string]any{
		"restrict_chat":          req.RestrictChat,
		"restrict_add_user":      req.RestrictAddUser,
		"restrict_create_group":  req.RestrictCreateGroup,
		"restrict_in_group_chat": req.RestrictInGroupChat,
	})

	return
}
