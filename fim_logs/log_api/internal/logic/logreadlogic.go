package logic

import (
	"context"

	"FIM/fim_logs/log_api/internal/svc"
	"FIM/fim_logs/log_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Log_readLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLog_readLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Log_readLogic {
	return &Log_readLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Log_readLogic) Log_read(req *types.LogReadRequest) (resp *types.LogReadResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
