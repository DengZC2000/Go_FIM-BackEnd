package Admin

import (
	"FIM/fim_chat/chat_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"errors"

	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Chat_admin_sessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChat_admin_sessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Chat_admin_sessionLogic {
	return &Chat_admin_sessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Chat_admin_session 作为接收者的用户，哪些人和他聊过
func (l *Chat_admin_sessionLogic) Chat_admin_session(req *types.ChatAdminSessionRequest) (resp *types.ChatAdminSessionResponse, err error) {
	var sendUserIDList []uint32
	l.svcCtx.DB.Model(chat_models.ChatModel{}).
		Where("rev_user_id = ?", req.RevUserId).
		Group("send_user_id").
		Select("send_user_id").Scan(&sendUserIDList)
	userList, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIdList: sendUserIDList,
	})
	if err != nil {
		return nil, errors.New("用户服务错误")
	}
	resp = &types.ChatAdminSessionResponse{}
	for u, info := range userList.UserInfo {
		resp.List = append(resp.List, types.UserInfo{
			UserID:   uint(u),
			Avatar:   info.Avatar,
			Nickname: info.NickName,
		})
	}
	resp.Count = len(userList.UserInfo)

	return
}
