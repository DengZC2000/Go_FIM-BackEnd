package logic

import (
	"FIM/fim_group/group_models"
	"FIM/fim_group/group_rpc/internal/svc"
	"FIM/fim_group/group_rpc/types/group_rpc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGroupSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGroupSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGroupSearchLogic {
	return &UserGroupSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGroupSearchLogic) UserGroupSearch(in *group_rpc.UserGroupSearchRequest) (resp *group_rpc.UserGroupSearchResponse, err error) {
	type Data struct {
		UserID uint32 `gorm:"column:user_id"`
		Count  uint32 `gorm:"column:count"`
	}
	var data []Data
	switch in.Mode {
	case 1: // 查询创建的个数
		l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
			Where("user_id in ? and role = ?", in.UserIdList, 1).
			Group("user_id").
			Select("user_id", "count(id) as count").
			Scan(&data)
	case 2: // 查询加入的个数
		l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
			Where("user_id in ?", in.UserIdList).
			Group("user_id").
			Select("user_id", "count(id) as count").
			Scan(&data)
	}
	var groupUserMap = map[uint32]uint32{}
	for _, u2 := range data {
		groupUserMap[u2.UserID] = u2.Count
	}
	resp = &group_rpc.UserGroupSearchResponse{}
	resp.Result = map[int32]int32{}
	for _, uid := range in.UserIdList {
		resp.Result[int32(uid)] = int32(groupUserMap[uid])
	}
	return resp, nil
}
