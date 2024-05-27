package logic

import (
	"FIM/fim_user/user_models"
	"context"
	"errors"

	"FIM/fim_user/user_rpc/internal/svc"
	"FIM/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserBaseInfoLogic {
	return &UserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserBaseInfoLogic) UserBaseInfo(in *user_rpc.UserBaseInfoRequest) (*user_rpc.UserBaseInfoResponse, error) {
	var user user_models.UserModel
	err := l.svcCtx.DB.Take(&user, in.UserId).Error
	if err != nil {
		return nil, errors.New("该用户不存在")
	}

	return &user_rpc.UserBaseInfoResponse{
		UserId:   uint32(user.ID),
		NickName: user.NickName,
		Avatar:   user.Avatar,
	}, nil
}
