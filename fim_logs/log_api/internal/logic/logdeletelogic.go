package logic

import (
	"context"

	"FIM/fim_logs/log_api/internal/svc"
	"FIM/fim_logs/log_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Log_deleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLog_deleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Log_deleteLogic {
	return &Log_deleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Log_deleteLogic) Log_delete(req *types.LogDeleteRequest) (resp *types.LogDeleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
