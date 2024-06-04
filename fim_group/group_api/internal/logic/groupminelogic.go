package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"FIM/fim_group/group_models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_mineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_mineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_mineLogic {
	return &Group_mineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_mineLogic) Group_mine(req *types.GroupMineRequest) (resp *types.GroupMineListResponse, err error) {
	//查群id列表
	var groupIDList []uint
	query := l.svcCtx.DB.Model(&group_models.GroupMemberModel{}).Where("user_id = ?", req.UserID)
	if req.Mode == 1 {
		//我创建的群聊
		query.Where("role = ?", 1)
	}
	query.Select("group_id").Scan(&groupIDList)
	groups, count, _ := list_query.ListQuery(l.svcCtx.DB, group_models.GroupModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preloads: []string{"MemberList"},
		Where:    l.svcCtx.DB.Where("id in ?", groupIDList),
	})
	resp = &types.GroupMineListResponse{}
	for _, model := range groups {
		var role int8
		for _, member := range model.MemberList {
			if req.UserID == member.UserID {
				role = int8(member.Role)
			}
		}
		resp.List = append(resp.List, types.GroupMineResponse{
			GroupID:          model.ID,
			GroupTitle:       model.Title,
			GroupAvatar:      model.Avatar,
			GroupMemberCount: len(model.MemberList),
			Role:             role,
			Mode:             req.Mode,
		})
	}
	resp.Count = int(count)
	return
}
