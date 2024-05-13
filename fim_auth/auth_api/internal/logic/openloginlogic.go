package logic

import (
	"FIM/fim_auth/auth_api/internal/svc"
	"FIM/fim_auth/auth_api/internal/types"
	"FIM/fim_auth/auth_models"
	"FIM/fim_user/user_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils/jwt"
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
	type OpenInfo struct {
		Nickname string
		OpenID   string
		Avatar   string
	}
	var info OpenInfo

	switch req.Flag {
	case "qq":
		QQInfo, openError := open_login.NewQQLogin(req.Code, open_login.QQConfig{
			AppID:    l.svcCtx.Config.QQ.AppID,
			AppKey:   l.svcCtx.Config.QQ.AppKey,
			Redirect: l.svcCtx.Config.QQ.Redirect,
		})
		info = OpenInfo{
			Nickname: QQInfo.Nickname,
			OpenID:   QQInfo.OpenID,
			Avatar:   QQInfo.FigureurlQQ,
		}
		err = openError

	default:
		err = errors.New("不支持的三方登录")
	}

	if err != nil {
		err = errors.New("登陆失败")
		return
	}
	fmt.Println(info)
	var user auth_models.UserModel
	err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenID).Error
	if err != nil {
		//注册逻辑
		fmt.Println("注册服务")
		res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
			Nickname:       info.Nickname,
			Password:       "",
			Role:           2,
			Avatar:         info.Avatar,
			Openid:         info.OpenID,
			RegisterSource: "QQ",
		})
		if err != nil {
			return nil, errors.New("注册失败")
		}
		user.Model.ID = uint(res.UserId)
		user.Role = 2
		user.NickName = info.Nickname

		l.svcCtx.DB.Create(&user_models.UserConfModel{
			UserID:        user.ID,
			RecallMessage: nil, //撤回消息的提示
			FriendOnline:  false,
			Sound:         true,
			SecureLink:    false,
			SavePwd:       false,
			SearchUser:    2, //允许别人找到你的方式
			Verification:  2, //需要验证问题
			Online:        true,
		})
	}

	//登录逻辑
	fmt.Println("登录服务")
	token, errr := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   user.ID,
		Nickname: user.NickName,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if errr != nil {
		err = errors.New("生成token失败")
		return
	}
	return &types.LoginResponse{Token: token}, nil

}
