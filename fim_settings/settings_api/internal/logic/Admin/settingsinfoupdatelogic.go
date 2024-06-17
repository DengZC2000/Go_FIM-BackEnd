package Admin

import (
	"FIM/common/models/ctype"
	"FIM/fim_settings/settings_model"
	"context"
	"errors"

	"FIM/fim_settings/settings_api/internal/svc"
	"FIM/fim_settings/settings_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Settings_info_updateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSettings_info_updateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Settings_info_updateLogic {
	return &Settings_info_updateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type MySettingsInfoUpdateRequest struct {
	Site *ctype.SiteType `json:"site"`
	QQ   *ctype.QQType   `json:"qq"`
}

func (l *Settings_info_updateLogic) Settings_info_update(req *MySettingsInfoUpdateRequest) (resp *types.SettingsInfoUpdateResponse, err error) {

	var settingsModel settings_model.SettingsModel
	l.svcCtx.DB.First(&settingsModel)
	if req.QQ.Key == "******" {
		req.QQ.Key = settingsModel.QQ.Key
	}
	err = l.svcCtx.DB.Model(&settingsModel).Updates(settings_model.SettingsModel{
		Site: req.Site,
		QQ:   req.QQ,
	}).Error
	if err != nil {
		return nil, errors.New("修改配置失败！")
	}
	return
}
