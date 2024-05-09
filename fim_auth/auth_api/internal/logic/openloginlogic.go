package logic

import (
	"FIM/fim_auth/auth_api/internal/svc"
	"FIM/fim_auth/auth_api/internal/types"
	"FIM/fim_auth/auth_models"
	"FIM/utils/open_login"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_loginLogic) Open_login(req *types.OpenLoginRequest) (resp *types.LoginResponse, err error) {
	switch req.Flag {
	case "qq":
		info, errr := open_login.NewQQLogin(req.Code, open_login.QQConfig{
			AppID:    l.svcCtx.Config.QQ.AppID,
			AppKey:   l.svcCtx.Config.QQ.AppKey,
			Redirect: l.svcCtx.Config.QQ.Redirect,
		})
		if errr != nil {
			err = errors.New("登陆失败")
			return
		}
		fmt.Println(info)
		var user auth_models.UserModel
		err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenID).Error
		if err != nil {
			//注册逻辑
			fmt.Println("注册服务")
		}

		//登录逻辑
		fmt.Println("登录服务")

	default:
		break
	}
	return
}
