package logic

import (
	"FIM/fim_user/user_models"
	"context"

	"FIM/fim_user/user_rpc/internal/svc"
	"FIM/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListInfoLogic {
	return &UserListInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListInfoLogic) UserListInfo(in *user_rpc.UserListInfoRequest) (*user_rpc.UserListInfoResponse, error) {

	var userList []user_models.UserModel
	l.svcCtx.DB.Find(&userList, in.UserIdList)

	resp := &user_rpc.UserListInfoResponse{}
	resp.UserInfo = make(map[uint32]*user_rpc.UserInfo)
	for _, i2 := range userList {
		resp.UserInfo[uint32(i2.ID)] = &user_rpc.UserInfo{
			NickName: i2.NickName,
			Avatar:   i2.Avatar,
		}
	}
	return resp, nil
}
