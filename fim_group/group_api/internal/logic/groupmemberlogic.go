package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/common/models/ctype"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_memberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_memberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_memberLogic {
	return &Group_memberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Data struct {
	GroupID        uint   `gorm:"column:group_id"`
	UserID         uint   `gorm:"column:user_id"`
	Role           int8   `gorm:"column:role"`
	CreatedAt      string `gorm:"column:created_at"`
	MemberNickname string `gorm:"column:member_nickname"`
	NewMsgDate     string `gorm:"column:new_msg_date"`
}

func (l *Group_memberLogic) Group_member(req *types.GroupMemberRequest) (resp *types.GroupMemberResponse, err error) {
	switch req.Sort {
	case "":
	case "new_msg_date desc", "new_msg_date asc":
	case "role asc":
	case "created_at desc", "created_at asc":
	default:
		return nil, errors.New("不支持的排序方式")
	}
	resp = &types.GroupMemberResponse{}
	column := fmt.Sprintf("(select group_msg_models.created_at from group_msg_models where group_member_models.group_id = %d and group_msg_models.send_user_id = user_id) as new_msg_date", req.ID)
	memberList, count, _ := list_query.ListQuery(l.svcCtx.DB, Data{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  req.Sort,
		},
		Table: func() (string, any) {
			return "(?) as u", l.svcCtx.DB.Model(&group_models.GroupMemberModel{
				GroupID: req.ID,
			}).Select("group_id",
				"user_id",
				"role",
				"created_at",
				"member_nickname",
				column)
		},
		Where: l.svcCtx.DB.Where("group_id = ?", req.ID),
	})
	fmt.Println(memberList, count)
	//拿到id
	var userIDList []uint32
	for _, data := range memberList {
		userIDList = append(userIDList, uint32(data.UserID))
	}

	userListResponse, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	//关于降级
	var userInfoMap = map[uint]ctype.UserInfo{}
	if err == nil {
		for u, info := range userListResponse.UserInfo {
			userInfoMap[uint(u)] = ctype.UserInfo{
				ID:       uint(u),
				Nickname: info.NickName,
				Avatar:   info.Avatar,
			}
		}
	} else {
		logx.Error(err)
	}
	//拿在线
	var userOnlineMap = map[uint]bool{}
	userOnlineResponse, err := l.svcCtx.UserRpc.UserOnlineList(l.ctx, &user_rpc.UserOnlineListRequest{})
	if err == nil {
		for _, u := range userOnlineResponse.UserIdList {
			userOnlineMap[uint(u)] = true
		}
	} else {
		logx.Error(err)
	}
	for _, data := range memberList {
		resp.List = append(resp.List, types.GroupMemberInfo{
			UserID:         data.UserID,
			UserNickname:   userInfoMap[data.UserID].Nickname,
			Avatar:         userInfoMap[data.UserID].Avatar,
			IsOnline:       userOnlineMap[data.UserID],
			Role:           data.Role,
			MemberNickname: data.MemberNickname,
			CreatedAt:      data.CreatedAt,
			NewMsgDate:     data.NewMsgDate,
		})
	}
	resp.Count = int(count)
	return
}
