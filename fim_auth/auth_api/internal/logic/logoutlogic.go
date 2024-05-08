package logic

import (
	"FIM/utils/jwt"
	"context"
	"errors"
	"fmt"
	"time"

	"FIM/fim_auth/auth_api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp string, err error) {
	if token == "" {
		err = errors.New("请输入token")
		return
	}
	claims, err := jwt.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("token错误")
		return
	}
	key := fmt.Sprintf("logout_%d", claims.UserID)
	now := time.Now()
	expiration := claims.ExpiresAt.Sub(now)
	l.svcCtx.Redis.SetNX(context.Background(), key, "", expiration)
	resp = "注销成功"
	return
}
