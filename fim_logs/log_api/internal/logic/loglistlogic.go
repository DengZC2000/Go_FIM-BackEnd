package logic

import (
	"context"

	"FIM/fim_logs/log_api/internal/svc"
	"FIM/fim_logs/log_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Log_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLog_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Log_listLogic {
	return &Log_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Log_listLogic) Log_list(req *types.LogListRequest) (resp *types.LogListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
