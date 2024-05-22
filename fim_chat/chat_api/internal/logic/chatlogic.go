package logic

import (
	"context"

	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Chat_Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChat_Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Chat_Logic {
	return &Chat_Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Chat_Logic) Chat_(req *types.ChatRequest) (resp *types.ChatResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
