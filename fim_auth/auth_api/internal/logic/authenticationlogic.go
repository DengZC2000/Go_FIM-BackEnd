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

func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp *types.AuthenticationResponse, err error) {
	if utils.InListRegex(l.svcCtx.Config.WhiteList, req.ValidPath) {
		logx.Infof("%s 在白名单中", req.ValidPath)
		return
	}

	if req.Token == "" {
		err = errors.New("请输入token")
		return
	}

	claims, err := jwt.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("认证失败")
		return
	}

	result := l.svcCtx.Redis.Exists(context.Background(), fmt.Sprintf("logout_%s", req.Token))
	if result.Err() != nil {
		// 处理Redis操作中的错误，可能不是用户已注销
		// 例如，你可以记录错误并返回一个通用的错误消息
		err = errors.New(result.Err().Error())
		return
	}

	if result.Val() == 1 {
		// 键存在，表示用户已注销
		err = errors.New("用户已注销，认证失败")
		return
	}

	resp = &types.AuthenticationResponse{UserID: claims.UserID, Role: int(claims.Role)}
	err = nil
	return
}
