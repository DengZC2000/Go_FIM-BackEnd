package settings_model

import (
	"FIM/common/models"
	"FIM/common/models/ctype"
)

type SettingsModel struct {
	models.Model
	Site *ctype.SiteType `json:"site"`
	QQ   *ctype.QQType   `json:"qq"`
}
