package logic

import (
	"FIM/fim_logs/logs_model"
	"context"

	"FIM/fim_logs/logs_api/internal/svc"
	"FIM/fim_logs/logs_api/internal/types"

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
	var logList []logs_model.LogModel
	l.svcCtx.DB.Find(&logList, req.IdList)
	if len(logList) > 0 {
		l.svcCtx.DB.Delete(&logList)
		l.svcCtx.ActionPusher.SetItemInfo("删除日志条数", len(logList))
	}
	l.svcCtx.ActionPusher.PushInfo("删除日志操作")
	l.svcCtx.ActionPusher.Commit(l.ctx)
	return
}
