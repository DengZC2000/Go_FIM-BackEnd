package logic

import (
	"FIM/fim_settings/settings_model"
	"context"

	"FIM/fim_settings/settings_api/internal/svc"
	"FIM/fim_settings/settings_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Settings_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSettings_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Settings_infoLogic {
	return &Settings_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Settings_infoLogic) Settings_info(req *types.SettingsInfoRequest) (resp *settings_model.SettingsModel, err error) {
	// 有且只有一条记录
	// 查询的时候，查不到就添加一条记录
	// 在系统启动的时候，查一下有没有，没有就加一条
	resp = &settings_model.SettingsModel{}
	l.svcCtx.DB.First(resp)
	resp.QQ.Key = "******"
	resp.QQ.WebPath = resp.QQ.GetPath()
	return
}
