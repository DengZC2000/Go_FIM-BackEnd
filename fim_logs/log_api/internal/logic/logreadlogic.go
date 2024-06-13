package logic

import (
	"FIM/fim_logs/logs_model"
	"context"
	"errors"

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
	var logModel logs_model.LogModel
	err = l.svcCtx.DB.Take(&logModel, req.ID).Error
	if err != nil {
		return nil, errors.New("日志记录不存在")
	}
	// 前端要判断一下，如果已经读取了，就不要再调接口了
	if logModel.IsRead {
		return
	}
	l.svcCtx.DB.Model(&logModel).Update("is_read", true)

	return
}
