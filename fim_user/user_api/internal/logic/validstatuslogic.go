package logic

import (
	"FIM/common/models/ctype"
	"FIM/fim_chat/chat_rpc/chat"
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"FIM/fim_user/user_models"
	"context"
	"encoding/json"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type Valid_statusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValid_statusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Valid_statusLogic {
	return &Valid_statusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Valid_statusLogic) Valid_status(req *types.FriendValidStatusRequest) (resp *types.FriendValidStatusResponse, err error) {
	var friendVerify user_models.FriendVerifyModel
	//我要操作状态，我得是接收者
	err = l.svcCtx.DB.Take(&friendVerify, "id = ? and rev_user_id = ?", req.VerifyID, req.UserID).Error
	if err != nil {
		return nil, errors.New("验证记录不存在")
	}
	if friendVerify.RevStatus != 0 {
		return nil, errors.New("不可更改状态")
	}
	switch req.Status {
	case 1: //同意
		friendVerify.RevStatus = 1
		//往好友表里去加
		l.svcCtx.DB.Create(&user_models.FriendModel{
			SendUserID: friendVerify.SendUserID,
			RevUserID:  friendVerify.RevUserID,
		})

		msg := ctype.Msg{
			Type: 1,
			TextMsg: &ctype.TextMsg{
				Content: "我们已经是好友了，开始聊天吧！",
			},
		}
		byteData, _ := json.Marshal(msg)

		//给对方发个消息
		_, err = l.svcCtx.ChatRpc.UserChatCreate(context.Background(), &chat.UserChatRequest{
			SendUserId: uint32(friendVerify.SendUserID),
			RevUserId:  uint32(friendVerify.RevUserID),
			Msg:        byteData,
			SystemMsg:  nil,
		})

		if err != nil {
			logx.Error(err)
			return nil, errors.New(err.Error())
		}
	case 2: //拒绝
		friendVerify.RevStatus = 2
	case 3: //忽略
		friendVerify.RevStatus = 3
	case 4: //删除
		//一条验证记录，是两个人看的
		l.svcCtx.DB.Delete(&friendVerify)
		return nil, nil
	}
	l.svcCtx.DB.Save(&friendVerify)

	return
}
