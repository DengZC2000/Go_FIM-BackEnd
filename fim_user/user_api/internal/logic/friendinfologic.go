package logic

import (
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"FIM/fim_user/user_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"encoding/json"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type Friend_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriend_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Friend_infoLogic {
	return &Friend_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Friend_infoLogic) Friend_info(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	var friend user_models.FriendModel

	if !friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("你不是他(她)的好友哦!")
	}
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user_rpc.UserInfoRequest{
		UserId: uint32(req.FriendID),
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	var friendInfo user_models.UserModel

	json.Unmarshal(res.Data, &friendInfo)
	resp = &types.FriendInfoResponse{
		UserID:   friendInfo.ID,
		Nickname: friendInfo.NickName,
		Profile:  friendInfo.Profile,
		Avatar:   friendInfo.Avatar,
		Notice:   friend.FriendNotice(req.UserID),
	}

	return
}
