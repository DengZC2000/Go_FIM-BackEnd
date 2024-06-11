package logs_model

import "FIM/common/models"

type LogModel struct {
	models.Model
	LogType      int8   `json:"log_type"` // 日志类型 2 操作日志 3 运行日志
	IP           string `gorm:"size:32" json:"ip"`
	Addr         string `gorm:"size:64" json:"addr"`
	UserID       uint   `json:"user_id"`
	UserNickname string `gorm:"size:64" json:"user_nickname"`
	UserAvatar   string `gorm:"size:128" json:"user_avatar"`
	Level        string `gorm:"size:32" json:"level"`
	Title        string `gorm:"size:32" json:"title"`
	Content      string `json:"content"`                // 日志详情
	Service      string `gorm:"size:32" json:"service"` // 服务 记录微服务的名称
}
