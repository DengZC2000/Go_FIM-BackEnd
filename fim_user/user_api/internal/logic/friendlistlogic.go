package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_user/user_models"
	"context"

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

	//var count int64
	//l.svcCtx.DB.Model(user_models.FriendModel{}).Where("send_user_id = ? or rev_user_id = ?", req.UserID, req.UserID).Count(&count)
	//if req.Limit <= 0 {
	//	req.Limit = 10
	//}
	//if req.Page <= 0 {
	//	req.Page = 1
	//}
	//offset := (req.Page - 1) * req.Limit
	//l.svcCtx.DB.Preload("SendUserModel").Preload("RevUserModel").Limit(req.Limit).Offset(offset).Find(&friends, "send_user_id = ? or rev_user_id = ?", req.UserID, req.UserID)
	friends, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.FriendModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preloads: []string{"SendUserModel", "RevUserModel"},
	})
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
		}
		if friend.RevUserID == req.UserID {
			//我是接受方,则要把发起方信息作为好友信息返回
			info.UserID = friend.SendUserID
			info.Nickname = friend.SendUserModel.NickName
			info.Avatar = friend.SendUserModel.Avatar
			info.Profile = friend.SendUserModel.Profile
			info.Notice = friend.RevUserNotice
		}
		list = append(list, info)
	}
	resp = &types.FriendListResponse{Count: int(count), List: list}
	return
}
