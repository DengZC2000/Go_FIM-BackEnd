package logic

import (
	"FIM/fim_auth/auth_api/internal/svc"
	"FIM/fim_auth/auth_api/internal/types"
	"FIM/fim_auth/auth_models"
	"FIM/utils/jwt"
	"FIM/utils/pwd"
	"context"
	"errors"
	"fmt"

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
	// 启用http request写入
	l.svcCtx.ActionPusher.IsRequest()
	l.svcCtx.ActionPusher.IsHeaders()
	l.svcCtx.ActionPusher.IsResponse()

	l.svcCtx.ActionPusher.SetItemInfo("Username", req.Username)
	l.svcCtx.ActionPusher.PushInfo("用户登陆操作")
	defer l.svcCtx.ActionPusher.Commit(l.ctx)

	count := l.svcCtx.DB.Take(&user, "id = ?", req.Username).RowsAffected
	if count != 1 {
		l.svcCtx.ActionPusher.PushError(fmt.Sprintf("%s 用户名不存在", req.Username))
		err = errors.New("用户名不存在")
		return
	}
	if !pwd.CheckPwd(user.Password, req.Password) {

		l.svcCtx.ActionPusher.PushError(fmt.Sprintf("%s 密码错误", req.Username))
		l.svcCtx.ActionPusher.SetItemInfo("password", req.Password)

		err = errors.New("密码错误")
		return
	}
	//判断用户的注册来源，第三方登录来的不能通过用户名和密码登录
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   user.ID,
		Nickname: user.NickName,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		l.svcCtx.ActionPusher.SetItemError("生成token失败", err.Error())
		l.svcCtx.ActionPusher.PushError("生成token失败")
		err = errors.New("生成token失败")
		return
	}

	l.svcCtx.ActionPusher.PushInfo(fmt.Sprintf("%s 用户登陆成功", user.NickName))
	ctx := context.WithValue(l.ctx, "UserID", fmt.Sprintf("%d", user.ID))
	l.svcCtx.ActionPusher.SetCtx(ctx)
	return &types.LoginResponse{Token: token}, nil
}
