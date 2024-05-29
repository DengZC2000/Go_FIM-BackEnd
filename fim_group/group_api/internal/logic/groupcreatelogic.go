package logic

import (
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils/set"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_createLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_createLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_createLogic {
	return &Group_createLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_createLogic) Group_create(req *types.GroupCreateRequest) (resp *types.GroupCreateResponse, err error) {
	var groupModel = group_models.GroupModel{
		Creator:      req.UserID,
		Abstract:     fmt.Sprintf("本群创建于%s，群主很懒，什么都没有留下", time.Now().Format("2006-01-02")),
		IsSearch:     false,
		Verification: 2,
		Size:         50,
	}
	groupUserList := []uint{req.UserID}
	switch req.Mode {
	case 1: //直接创建模式
		if req.Name == "" {
			return nil, errors.New("群名不可为空")
		}
		if req.Size > 1000 {
			return nil, errors.New("群规模错误")
		}
		groupModel.Title = req.Name
		groupModel.Size = req.Size
		groupModel.IsSearch = req.IsSearch
	case 2: //选人创建模式
		if len(req.UserIDList) == 0 {
			return nil, errors.New("没有要选择的好友")
		}
		// 去算选择的用户昵称，是不是超过最大长度
		// 群名是32
		// 调用户信息列表
		var userIDList = []uint32{uint32(req.UserID)} //先把自己放进去
		for _, u := range req.UserIDList {
			userIDList = append(userIDList, uint32(u))
			groupUserList = append(groupUserList, u)
		}
		//判断邀请的这些人是不是你的好友，只要有一个不是，那就说明用户是乱填的
		userFriendResponse, err := l.svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
			User: uint32(req.UserID),
		})
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		var friendIDList []uint
		for _, i2 := range userFriendResponse.FriendList {
			friendIDList = append(friendIDList, uint(i2.UserId))
		}
		//判断两个是不是一致的
		slice := set.Difference(req.UserIDList, friendIDList)
		if len(slice) != 0 {
			return nil, errors.New("选择的用户列表中有人不是你的好友")
		}
		userListRes, err1 := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
			UserIdList: userIDList,
		})
		if err1 != nil {
			logx.Error(err1)
			return nil, errors.New("用户服务错误")
		}
		//去算昵称的长度，算到第几个人大于32,《xxx、xxx、xxx的群聊》
		var nameList []string
		for _, info := range userListRes.UserInfo {
			if len([]rune(strings.Join(nameList, "、"))) > 29 {
				nameList = nameList[:len(nameList)-1]
				break
			}
			nameList = append(nameList, info.NickName)
		}
		groupModel.Title = strings.Join(nameList, "、") + "的群聊"
	default:
		return nil, errors.New("不支持的群聊创建方式")
	}
	//群头像
	// 1.默认头像 2.文字头像
	groupModel.Avatar = string([]rune(groupModel.Title)[0])
	err = l.svcCtx.DB.Create(&groupModel).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("群组创建失败")
	}
	var members []group_models.GroupMemberModel
	for i, u := range groupUserList {
		member := group_models.GroupMemberModel{
			GroupID: groupModel.ID,
			UserID:  u,
			Role:    3,
		}
		if i == 0 {
			member.Role = 1
		}
		members = append(members, member)
	}
	err = l.svcCtx.DB.Create(&members).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("群组成员创建失败")
	}
	return
}
