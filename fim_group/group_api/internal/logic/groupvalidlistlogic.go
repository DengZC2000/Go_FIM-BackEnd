package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_valid_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_valid_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_valid_listLogic {
	return &Group_valid_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_valid_listLogic) Group_valid_list(req *types.GroupValidListRequest) (resp *types.GroupValidListResponse, err error) {
	// 群验证列表 自己得是群管理员或者群主才行
	var groupIDList []uint //我管理的群
	l.svcCtx.DB.Model(&group_models.GroupMemberModel{}).Where("user_id = ? and (role = 1 or role  = 2)", req.UserID).Select("group_id").Scan(&groupIDList)
	//先去查自己管理了哪些群，然后去找这些群的验证表
	groups, count, _ := list_query.ListQuery(l.svcCtx.DB, group_models.GroupVerifyModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preloads: []string{"GroupModel"},
		Where:    l.svcCtx.DB.Where("group_id in ?", groupIDList),
	})
	resp = &types.GroupValidListResponse{}
	var userIDList []uint32
	for _, group := range groups {
		userIDList = append(userIDList, uint32(group.UserID))
	}
	userList, err1 := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})

	for _, group := range groups {
		info := types.GroupValidInfoResponse{
			ID:                 group.ID,
			GroupID:            group.GroupID,
			UserID:             group.UserID,
			Status:             group.Status,
			AdditionalMessages: group.AdditionalMessages,
			Title:              group.GroupModel.Title,
			CreatedAt:          group.CreatedAt.String(),
			Type:               group.Type,
		}
		if group.VerificationQuestion != nil {
			info.VerificationQuestion = &types.VerificationQuestion{
				Problem1: group.VerificationQuestion.Problem1,
				Problem2: group.VerificationQuestion.Problem2,
				Problem3: group.VerificationQuestion.Problem3,
				Answer1:  group.VerificationQuestion.Answer1,
				Answer2:  group.VerificationQuestion.Answer2,
				Answer3:  group.VerificationQuestion.Answer3,
			}
		}
		if err1 == nil {
			info.UserNickname = userList.UserInfo[uint32(info.UserID)].NickName
			info.UserAvatar = userList.UserInfo[uint32(info.UserID)].Avatar
		}
		resp.List = append(resp.List, info)
	}
	resp.Count = int(count)
	return
}
