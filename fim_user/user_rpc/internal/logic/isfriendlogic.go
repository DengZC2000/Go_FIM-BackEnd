package logic

import (
	"FIM/fim_user/user_models"
	"context"

	"FIM/fim_user/user_rpc/internal/svc"
	"FIM/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendLogic {
	return &IsFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFriendLogic) IsFriend(in *user_rpc.IsFriendRequest) (*user_rpc.IsFriendResponse, error) {
	var friend user_models.FriendModel
	if friend.IsFriend(l.svcCtx.DB, uint(in.User1), uint(in.User2)) {
		return &user_rpc.IsFriendResponse{IsFriend: true}, nil
	}
	return &user_rpc.IsFriendResponse{IsFriend: false}, nil
}
