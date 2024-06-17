package Admin

import (
	"FIM/fim_chat/chat_models"
	"context"

	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Chat_admin_history_removeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChat_admin_history_removeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Chat_admin_history_removeLogic {
	return &Chat_admin_history_removeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Chat_admin_history_removeLogic) Chat_admin_history_remove(req *types.ChatAdminHistoryRemoveRequest) (resp *types.ChatAdminHistoryRemoveResponse, err error) {
	var msgList []chat_models.ChatModel
	l.svcCtx.DB.Find(&msgList, "id in ?", req.IDList).Delete(&msgList)
	logx.Infof("删除聊天记录个数 %d", len(msgList))
	var userChatDeleteList []chat_models.UserChatDeleteModel
	l.svcCtx.DB.Find(&userChatDeleteList, "chat_id in ?", req.IDList).Delete(&userChatDeleteList)
	logx.Infof("删除关联用户删除聊天记录个数 %d", len(userChatDeleteList))
	return
}
