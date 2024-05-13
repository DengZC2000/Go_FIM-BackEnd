package logic

import (
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"FIM/fim_user/user_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_infoLogic {
	return &User_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_infoLogic) User_info(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	userInfo, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: uint32(req.UserID),
	})
	if err != nil {
		return nil, err
	}
	var user user_models.UserModel
	err = json.Unmarshal(userInfo.Data, &user)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据错误")
	}

	fmt.Println(string(userInfo.Data))
	return &types.UserInfoResponse{
		UserID:         user.ID,
		Nickname:       user.NickName,
		Role:           user.Role,
		Profile:        user.Profile,
		Avatar:         user.Avatar,
		RegisterSource: user.RegisterSource,
	}, nil
}
