package logic

import (
	"FIM/fim_chat/chat_models"
	"context"

	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Chat_deleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChat_deleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Chat_deleteLogic {
	return &Chat_deleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Chat_deleteLogic) Chat_delete(req *types.ChatDeleteRequest) (resp *types.ChatDeleteResponse, err error) {
	var chatList []chat_models.ChatModel
	l.svcCtx.DB.Find(&chatList, req.IdList)

	var userDeleteChatList []chat_models.UserChatDeleteModel
	l.svcCtx.DB.Find(&userDeleteChatList, "chat_id IN (?)", req.IdList)
	chatDeleteMap := map[uint]struct{}{}
	for _, model := range userDeleteChatList {
		chatDeleteMap[model.ChatID] = struct{}{}
	}

	var deleteChatList []chat_models.UserChatDeleteModel

	if len(chatList) > 0 {
		for _, chat := range chatList {
			//不是自己的聊天记录
			if !(chat.SendUserID == req.UserID || chat.RevUserID == req.UserID) {
				continue
			}
			//已经删过的聊天记录
			_, ok := chatDeleteMap[chat.ID]
			if ok {
				continue
			}
			deleteChatList = append(deleteChatList, chat_models.UserChatDeleteModel{
				UserID: req.UserID,
				ChatID: chat.ID,
			})
		}
	}
	if len(deleteChatList) > 0 {
		l.svcCtx.DB.Create(&deleteChatList)
	}

	return
}
