package Admin

import (
	"FIM/fim_group/group_models"
	"context"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_message_removeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_message_removeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_message_removeLogic {
	return &Group_message_removeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_message_removeLogic) Group_message_remove(req *types.GroupMessageRemoveRequest) (resp *types.GroupMessageRemoveResponse, err error) {
	// 系统普通成员删除是软删除
	// 系统管理员删除就是真删除了
	// 值得注意的是，现在原消息没有了，前端没有办法点击跳转了（针对回复消息、引用消息等）
	var messageList []group_models.GroupMsgModel
	l.svcCtx.DB.Find(&messageList, "id in ? ", req.IdList).Delete(&messageList)
	var userDeleteMessageList []group_models.GroupUserMsgDeleteModel
	l.svcCtx.DB.Find(&userDeleteMessageList, "msg_id in ?", req.IdList).Delete(&userDeleteMessageList)
	logx.Infof("删除聊天记录个数 %d ，关联用户删除聊天记录个数 %d", len(messageList), len(userDeleteMessageList))
	return
}
