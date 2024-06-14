package Admin

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_user/user_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_listLogic {
	return &User_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_listLogic) User_list(req *types.UserListResquest) (resp *types.UserListResponse, err error) {
	list, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.UserModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Key:   req.Key,
		},
		Likes: []string{"nickname", "ip"},
	})
	resp = &types.UserListResponse{}
	var userIDList []uint32
	for _, model := range list {
		userIDList = append(userIDList, uint32(model.ID))
	}
	// 去查用户在线状态
	var userOnlineMap = map[uint]bool{}
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(l.ctx, &user_rpc.UserOnlineListRequest{})
	if err == nil {
		for _, u := range userOnlineResponse.UserIdList {
			userOnlineMap[uint(u)] = true
		}
	} else {
		logx.Error(err)
	}
	// 查用户创建的群聊个数
	// 查用户加入的群聊个数

	//查用户发送的消息个数
	for _, model := range list {
		info := types.UserListInfoResponse{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.String(),
			Nickname:  model.NickName,
			Avatar:    model.Avatar,
			IP:        model.IP,
			Addr:      model.Address,
			IsOnline:  userOnlineMap[model.ID],
		}
		resp.List = append(resp.List, info)
	}
	resp.Count = int(count)

	return
}
