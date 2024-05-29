package logic

import (
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils/set"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_infoLogic {
	return &Group_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_infoLogic) Group_info(req *types.GroupInfoRequest) (resp *types.GroupInfoResponse, err error) {
	// 谁能调这个接口，必须得是这个群的成员
	var groupModel group_models.GroupModel
	err = l.svcCtx.DB.Preload("MemberList").Take(&groupModel, req.ID).Error
	if err != nil {
		return nil, errors.New("群不存在")
	}
	resp = &types.GroupInfoResponse{
		GroupID:     groupModel.ID,
		Title:       groupModel.Title,
		Abstract:    groupModel.Abstract,
		Avatar:      groupModel.Avatar,
		MemberCount: len(groupModel.MemberList),
	}
	// 查用户列表信息
	var userIDList []uint32
	var AlluserIDList []uint32
	for _, model := range groupModel.MemberList {
		if model.Role == 1 || model.Role == 2 {
			userIDList = append(userIDList, uint32(model.UserID))
		}
		AlluserIDList = append(AlluserIDList, uint32(model.UserID))
	}
	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	if err != nil {
		return nil, err
	}
	var creator types.UserInfo
	var adminList []types.UserInfo
	for _, model := range groupModel.MemberList {
		if model.Role == 3 {
			continue
		}
		user := types.UserInfo{
			UserID:   model.UserID,
			Avatar:   userListResponse.UserInfo[uint32(model.UserID)].Avatar,
			Nickname: userListResponse.UserInfo[uint32(model.UserID)].NickName,
		}

		if model.Role == 1 {
			creator = user
			continue
		}
		if model.Role == 2 {
			adminList = append(adminList, user)
		}
	}
	// 算在线用户总数
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(context.Background(), &user_rpc.UserOnlineListRequest{})
	if err == nil {
		slice := set.Intersect(AlluserIDList, userOnlineResponse.UserIdList)
		resp.MemberOnlineCount = len(slice)
	}
	// 用户服务需要去写一个在线的用户列表的方法
	resp.Creator = creator
	resp.AdminList = adminList

	return
}
