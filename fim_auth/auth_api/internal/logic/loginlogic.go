package logic

import (
	"FIM/fim_auth/auth_models"
	"FIM/utils/jwt"
	"FIM/utils/pwd"
	"context"
	"errors"
	"fmt"

	"FIM/fim_auth/auth_api/internal/svc"
	"FIM/fim_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	var user auth_models.UserModel

	count := l.svcCtx.DB.Take(&user, "id = ?", req.Username).RowsAffected
	if count != 1 {
		err = errors.New("用户名或密码错误")
		return
	}
	fmt.Println(user)
	if !pwd.CheckPwd(user.Password, req.Password) {
		err = errors.New("用户名或密码错误")
		return
	}
	//判断用户的注册来源，第三方登录来的不能通过用户名和密码登录
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   user.ID,
		Nickname: user.NickName,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		err = errors.New("生成token失败")
		return
	}
	err = l.svcCtx.KqPusherClient.Push(fmt.Sprintf("昵称：%s -- 登陆成功", user.NickName))
	if err != nil {
		logx.Error(err)
	}
	return &types.LoginResponse{Token: token}, nil
}
