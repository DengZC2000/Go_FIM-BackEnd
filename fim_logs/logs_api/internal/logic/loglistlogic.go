package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_logs/logs_api/internal/svc"
	"FIM/fim_logs/logs_api/internal/types"
	"FIM/fim_logs/logs_model"
	"context"

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
	logList, count, _ := list_query.ListQuery(l.svcCtx.DB, logs_model.LogModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},
		Likes: []string{"ip", "user_nickname", "title"},
	})
	resp = &types.LogListResponse{}
	for _, log := range logList {
		info := types.LogInfoResponse{
			ID:           log.ID,
			CreatedAt:    log.CreatedAt.String(),
			LogType:      log.LogType,
			IP:           log.IP,
			Addr:         log.Addr,
			UserID:       log.UserID,
			UserNickname: log.UserNickname,
			UserAvatar:   log.UserAvatar,
			Level:        log.Level,
			Title:        log.Title,
			Content:      log.Content,
			Service:      log.Service,
			IsRead:       log.IsRead,
		}
		resp.List = append(resp.List, info)
	}
	resp.Count = int(count)

	return
}
