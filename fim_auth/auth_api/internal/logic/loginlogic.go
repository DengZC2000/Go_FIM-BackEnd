package logic

import (
	"FIM/fim_auth/auth_models"
	"FIM/utils/jwt"
	"FIM/utils/pwd"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

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
	ctx := context.WithValue(l.ctx, "UserID", fmt.Sprintf("%d", user.ID))
	type Request1 struct {
		LogType int8   `json:"log_type"` // 日志类型 2 操作日志 3 运行日志
		IP      string `json:"ip"`
		UserID  uint   `json:"user_id"`
		Level   string `json:"level"`
		Title   string `json:"title"`
		Content string `json:"content"` // 日志详情
		Service string `json:"service"` // 服务 记录微服务的名称
	}
	userID := ctx.Value("UserID").(string)
	userIDInt, _ := strconv.Atoi(userID)
	req1 := Request1{
		LogType: 2,
		IP:      ctx.Value("ClientIP").(string),
		Level:   "info",
		UserID:  uint(userIDInt),
		Title:   fmt.Sprintf("%s 登陆成功", user.NickName),
		Content: "xxx",
		Service: l.svcCtx.Config.Name,
	}
	byteData, _ := json.Marshal(req1)
	err = l.svcCtx.KqPusherClient.Push(string(byteData))
	if err != nil {
		logx.Error(err)
	}
	return &types.LoginResponse{Token: token}, nil
}
