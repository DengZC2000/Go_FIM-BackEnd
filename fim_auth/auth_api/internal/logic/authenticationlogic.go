package logic

import (
	"FIM/fim_auth/auth_api/internal/svc"
	"FIM/utils/jwt"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp string, err error) {
	if token == "" {
		err = errors.New("请输入token")
		return
	}

	claims, err := jwt.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("认证失败")
		return
	}
	_, err = l.svcCtx.Redis.Get(context.Background(), fmt.Sprintf("logout_%d", claims.UserID)).Result()
	if err == nil {
		err = errors.New("用户已注销，认证失败")
		return
	}
	resp = "认证成功"
	err = nil
	return
}
