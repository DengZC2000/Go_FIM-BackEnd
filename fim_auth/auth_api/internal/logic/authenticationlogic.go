package logic

import (
	"FIM/fim_auth/auth_api/internal/svc"
	"FIM/fim_auth/auth_api/internal/types"
	"FIM/utils"
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

func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp string, err error) {

	if utils.InList(l.svcCtx.Config.WhiteList, req.ValidPath) {
		logx.Infof("%s 在白名单中", req.ValidPath)
		return "ok", nil
	}

	if req.Token == "" {
		err = errors.New("请输入token")
		return
	}

	_, err = jwt.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("认证失败")
		return
	}

	result := l.svcCtx.Redis.Exists(context.Background(), fmt.Sprintf("logout_%s", req.Token))
	if result.Val() == 1 || result.Err() != nil {
		err = errors.New("用户已注销，认证失败")
		return
	}

	resp = "认证成功"
	err = nil
	return
}
