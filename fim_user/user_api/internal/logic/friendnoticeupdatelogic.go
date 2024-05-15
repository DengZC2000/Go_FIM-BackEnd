package logic

import (
	"FIM/fim_user/user_models"
	"context"
	"errors"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Friend_notice_updateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriend_notice_updateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Friend_notice_updateLogic {
	return &Friend_notice_updateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Friend_notice_updateLogic) Friend_notice_update(req *types.FriendNoticeUpdateRequest) (resp *types.FriendNoticeUpdateResponse, err error) {
	var friend user_models.FriendModel
	if !friend.IsFriend(l.svcCtx.DB, req.FriendID, req.UserID) {
		return nil, errors.New("他（她）还不是你的好友哦~")
	}
	if friend.SendUserID == req.UserID {
		//我是发起方
		if friend.SendUserNotice == req.Notice {
			return
		}
		l.svcCtx.DB.Model(&friend).Update("send_user_notice", req.Notice)
	}
	if friend.RevUserID == req.UserID {
		//我是接受方
		if friend.RevUserNotice == req.Notice {
			return
		}
		l.svcCtx.DB.Model(&friend).Update("rev_user_notice", req.Notice)
	}
	return
}
