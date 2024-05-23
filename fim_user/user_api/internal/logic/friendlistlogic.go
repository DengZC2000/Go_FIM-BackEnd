package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_user/user_models"
	"context"
	"strconv"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Friend_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriend_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Friend_listLogic {
	return &Friend_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Friend_listLogic) Friend_list(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {

	friends, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.FriendModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preloads: []string{"SendUserModel", "RevUserModel"},
	})

	// 查哪些用户在线
	onlineMap := l.svcCtx.Redis.HGetAll(context.Background(), "online").Val()
	var onlineUserMap = map[uint]bool{}
	for key, _ := range onlineMap {
		val, err1 := strconv.Atoi(key)
		if err1 != nil {
			logx.Error(err1)
			continue
		}
		onlineUserMap[uint(val)] = true
	}
	var list []types.FriendInfoResponse
	for _, friend := range friends {
		info := types.FriendInfoResponse{}
		if friend.SendUserID == req.UserID {
			//我是发起方,则要把接收方信息返回
			info.UserID = friend.RevUserID
			info.Nickname = friend.RevUserModel.NickName
			info.Avatar = friend.RevUserModel.Avatar
			info.Profile = friend.RevUserModel.Profile
			info.Notice = friend.SendUserNotice
			info.IsOnline = onlineUserMap[friend.RevUserID]
		}
		if friend.RevUserID == req.UserID {
			//我是接受方,则要把发起方信息作为好友信息返回
			info.UserID = friend.SendUserID
			info.Nickname = friend.SendUserModel.NickName
			info.Avatar = friend.SendUserModel.Avatar
			info.Profile = friend.SendUserModel.Profile
			info.Notice = friend.RevUserNotice
			info.IsOnline = onlineUserMap[friend.SendUserID]
		}
		list = append(list, info)
	}
	resp = &types.FriendListResponse{Count: int(count), List: list}
	return
}
