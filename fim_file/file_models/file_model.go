package file_models

import (
	"FIM/common/models"
	"github.com/google/uuid"
)

type FileModel struct {
	models.Model
	Uid      uuid.UUID //文件唯一id /api/file/{uuid}
	UserID   uint      `json:"user_id"`   //用户id
	FileName string    `json:"file_name"` //文件的名称
	Size     int64     `json:"size"`      //文件的大小
	Path     string    `json:"path"`      //文件的路径
	Hash     string    `json:"hash"`      //文件的哈希
}

func (file *FileModel) WebPath() string {
	return "/api/file/" + file.Uid.String()
}
