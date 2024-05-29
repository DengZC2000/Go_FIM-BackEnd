package logic

import (
	"context"
	"strconv"

	"FIM/fim_user/user_rpc/internal/svc"
	"FIM/fim_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOnlineListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOnlineListLogic {
	return &UserOnlineListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOnlineListLogic) UserOnlineList(in *user_rpc.UserOnlineListRequest) (resp *user_rpc.UserOnlineListResponse, err error) {
	// 查哪些用户在线
	resp = &user_rpc.UserOnlineListResponse{}
	onlineMap := l.svcCtx.MyRedis.HGetAll(context.Background(), "online").Val()
	for key, _ := range onlineMap {
		val, err1 := strconv.Atoi(key)
		if err1 != nil {
			logx.Error(err1)
			continue
		}
		resp.UserIdList = append(resp.UserIdList, uint32(val))
	}
	return
}
