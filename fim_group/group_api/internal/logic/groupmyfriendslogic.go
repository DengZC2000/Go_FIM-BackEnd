package logic

import (
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_my_friendsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_my_friendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_my_friendsLogic {
	return &Group_my_friendsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_my_friendsLogic) Group_my_friends(req *types.GroupFriendsListRequest) (resp *types.GroupFriendsListResponse, err error) {
	//需要去查我的好友列表
	friendResponse, err := l.svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
		User: uint32(req.UserID),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	//这个群的成员列表，组成一个map
	var memberList []group_models.GroupMemberModel
	l.svcCtx.DB.Find(&memberList, "group_id = ?", req.ID)
	var memberMap = map[uint]bool{}
	for _, model := range memberList {
		memberMap[model.UserID] = true
	}
	resp = &types.GroupFriendsListResponse{}
	for _, info := range friendResponse.FriendList {
		resp.List = append(resp.List, types.GroupFriendsResponse{
			UserID:    uint(info.UserId),
			Avatar:    info.Avatar,
			Nickname:  info.NickName,
			IsInGroup: memberMap[uint(info.UserId)],
		})
	}
	return
}
