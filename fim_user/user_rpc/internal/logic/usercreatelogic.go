package logic

import (
	"FIM/fim_user/user_models"
	"context"
	"log"

	"FIM/fim_user/user_rpc/internal/svc"
	"FIM/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc UserInfo(UserInfoRequest)returns(UserInfoResponse);
func (l *UserCreateLogic) UserCreate(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	user := user_models.UserModel{
		NickName: in.Nickname,
		Avatar:   in.Avatar,
		Role:     int8(in.Role),
		OpenID:   in.Openid,
	}
	err := l.svcCtx.DB.Create(&user).Error
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &user_rpc.UserCreateResponse{UserId: int32(user.ID)}, nil
}
