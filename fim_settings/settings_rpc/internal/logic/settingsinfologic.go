package logic

import (
	"FIM/fim_settings/settings_model"
	"context"
	"encoding/json"

	"FIM/fim_settings/settings_rpc/internal/svc"
	"FIM/fim_settings/settings_rpc/types/settings_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SettingsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSettingsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SettingsInfoLogic {
	return &SettingsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SettingsInfoLogic) SettingsInfo(in *settings_rpc.SettingsInfoRequest) (*settings_rpc.SettingsInfoResponse, error) {
	var settingModel settings_model.SettingsModel
	l.svcCtx.DB.First(&settingModel)
	byteData, _ := json.Marshal(settingModel)

	return &settings_rpc.SettingsInfoResponse{Data: byteData}, nil
}
