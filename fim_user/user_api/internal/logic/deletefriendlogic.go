package logic

import (
	"FIM/fim_user/user_models"
	"context"
	"errors"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Delete_friendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete_friendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Delete_friendLogic {
	return &Delete_friendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete_friendLogic) Delete_friend(req *types.DeleteFriendRequest) (resp *types.DeleteFriendResponse, err error) {
	var friend user_models.FriendModel

	if !friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("你和他（她）不是好友，无法删除")
	}
	l.svcCtx.DB.Delete(&friend)
	return
}
