package service

import (
	"FIM/common/models/ctype"
	"FIM/fim_settings/settings_api/internal/svc"
	"FIM/fim_settings/settings_model"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func InitSettings(ctx *svc.ServiceContext) {
	var settingModel settings_model.SettingsModel
	err := ctx.DB.First(&settingModel).Error
	if err != nil {
		err = ctx.DB.Create(&settings_model.SettingsModel{
			Site: &ctype.SiteType{
				CreatedAt: time.Now().Format("2006-01-02"),
				Version:   "1.0.0",
			},
			QQ: &ctype.QQType{},
		}).Error
	}
	if err != nil {
		panic(err)
	}
	logx.Info("插入默认系统数据成功")
}
