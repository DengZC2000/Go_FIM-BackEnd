package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils/set"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_searchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_searchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_searchLogic {
	return &Group_searchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_searchLogic) Group_search(req *types.GroupSearchRequest) (resp *types.GroupSearchListResponse, err error) {
	//先找所有的用户
	//IsSearch 为false的就表示不能被搜索到 IsSearch = true 就可以通过搜索加入群聊

	groups, count, _ := list_query.ListQuery(l.svcCtx.DB, group_models.GroupModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preloads: []string{"MemberList"},
		Where:    l.svcCtx.DB.Where("is_search = 1 and (id = ? or title like ?)", req.Key, fmt.Sprintf("%%%s%%", req.Key)),
	})
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(context.Background(), &user_rpc.UserOnlineListRequest{})
	if err != nil {
		return nil, err
	}
	//称之为： 服务降级，如果用户的rpc方法挂了，只是页面上看到的人数是0而已，不会影响群搜索这个功能
	var userOnlineIDList []uint
	for _, u := range userOnlineResponse.UserIdList {
		userOnlineIDList = append(userOnlineIDList, uint(u))
	}
	resp = &types.GroupSearchListResponse{}
	for _, group := range groups {
		var IDList []uint
		var isInGroup = false
		for _, member := range group.MemberList {
			IDList = append(IDList, member.UserID)
			if member.UserID == req.UserID {
				isInGroup = true
			}
		}
		resp.List = append(resp.List, types.GroupSearchResponse{
			GroupID:         group.ID,
			Title:           group.Title,
			Abstract:        group.Abstract,
			Avatar:          group.Avatar,
			UserCount:       len(group.MemberList),
			UserOnlineCount: len(set.Intersect(IDList, userOnlineIDList)), //这个群的在线用户总数
			IsInGroup:       isInGroup,
		})
	}
	resp.Count = int(count)
	return
}
