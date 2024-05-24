package logic

import (
	"FIM/fim_chat/chat_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"errors"

	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_topLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_topLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_topLogic {
	return &User_topLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_topLogic) User_top(req *types.UserTopRequest) (resp *types.UserTopResponse, err error) {
	if req.UserID != req.FriendID {

		//是否是好友
		res, err := l.svcCtx.UserRpc.IsFriend(context.Background(), &user_rpc.IsFriendRequest{
			User1: uint32(req.UserID),
			User2: uint32(req.FriendID),
		})
		if err != nil {
			return nil, err
		}
		if !res.IsFriend {
			return nil, errors.New("你们还不是好友呢")
		}
	}
	var topUser chat_models.TopUserModel
	err1 := l.svcCtx.DB.Take(&topUser, "user_id = ? and top_user_id = ?", req.UserID, req.FriendID).Error
	if err1 != nil {
		//没有置顶
		l.svcCtx.DB.Create(&chat_models.TopUserModel{
			UserID:    req.UserID,
			TopUserID: req.FriendID,
		})
		return
	}
	//已经有置顶了
	l.svcCtx.DB.Delete(&topUser, "user_id = ? and top_user_id = ?", req.UserID, req.FriendID)
	return
}
