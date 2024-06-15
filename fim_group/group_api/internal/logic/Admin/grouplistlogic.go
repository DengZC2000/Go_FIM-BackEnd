package Admin

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/common/models/ctype"
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_listLogic {
	return &Group_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_listLogic) Group_list(req *types.GroupListRequest) (resp *types.GroupListResponse, err error) {
	resp = &types.GroupListResponse{}
	list, count, _ := list_query.ListQuery(l.svcCtx.DB, group_models.GroupModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Key:   req.Key,
			Sort:  "created_at desc",
		},
		Likes:    []string{"title"},
		Preloads: []string{"MemberList", "GroupMsgList"},
	})
	var userIDList []uint32
	for _, model := range list {
		for _, memberModel := range model.MemberList {
			userIDList = append(userIDList, uint32(memberModel.UserID))
		}
	}
	utils.DeduplicationList(userIDList)
	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	if err != nil {
		return nil, errors.New("用户服务错误")
	}
	var userInfoMap = map[uint]ctype.UserInfo{}
	for u, info := range userListResponse.UserInfo {
		userInfoMap[uint(u)] = ctype.UserInfo{
			ID:       uint(u),
			Nickname: info.NickName,
			Avatar:   info.Avatar,
		}
	}

	var userOnlineMap = map[uint]bool{}
	userOnline, err := l.svcCtx.UserRpc.UserOnlineList(l.ctx, &user_rpc.UserOnlineListRequest{})
	if err == nil {
		for _, u := range userOnline.UserIdList {
			userOnlineMap[uint(u)] = true
		}
	} else {
		return nil, errors.New("用户服务错误")
	}
	for _, model := range list {
		info := types.GroupListInfoResponse{
			ID:           model.ID,
			CreatedAt:    model.CreatedAt.String(),
			Title:        model.Title,
			Abstract:     model.Abstract,
			Avatar:       model.Avatar,
			MemberCount:  len(model.MemberList),
			MessageCount: len(model.GroupMsgList),
			Creater: types.UserInfo{
				UserID:   model.Creator,
				Avatar:   userInfoMap[model.Creator].Avatar,
				Nickname: userInfoMap[model.Creator].Nickname,
			},
		}
		var adminList []types.UserInfo
		for _, memberModel := range model.MemberList {
			_, ok := userOnlineMap[memberModel.UserID]
			if ok {
				info.MemberOnlineCount++
			}
			if memberModel.Role == 2 {
				adminList = append(adminList, types.UserInfo{
					UserID:   memberModel.UserID,
					Avatar:   userInfoMap[memberModel.UserID].Avatar,
					Nickname: userInfoMap[memberModel.UserID].Nickname,
				})
			}
		}
		info.AdminList = adminList
		resp.List = append(resp.List, info)
	}
	resp.Count = int(count)

	return
}
