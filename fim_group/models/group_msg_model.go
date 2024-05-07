package models

import (
	"FIM/common/models"
	"FIM/common/models/ctype"
)

// GroupMsgModel 群消息表
type GroupMsgModel struct {
	models.Model
	GroupID    uint             `json:"group_id"` //群ID
	GroupModel GroupModel       `gorm:"foreignKey:GroupID" json:"-"`
	SendUserID uint             `json:"send_user_id"`               //发送者ID
	MsgType    int8             `json:"msg_type"`                   //消息类型 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回消息 9 回复消息 10 引用消息 11 @消息
	MsgPreview string           `gorm:"size:64" json:"msg_preview"` //消息预览
	Msg        ctype.Msg        `json:"msg"`                        //消息内容
	SystemMsg  *ctype.SystemMsg `json:"system_msg"`                 //系统提示
}
