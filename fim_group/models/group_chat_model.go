package models

import (
	"FIM/common/models"
	"FIM/common/models/ctype"
)

type GroupChatModel struct {
	models.Model
	SendUserID uint             `json:"send_user_id"` //发起验证方
	MsgType    int8             `json:"msg_type"`     //消息类型 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回消息 9 回复消息 10 引用消息 11 @消息
	MsgPreview string           `json:"msg_preview"`  //消息预览
	Msg        ctype.Msg        `json:"msg"`          //消息内容
	SystemMsg  *ctype.SystemMsg `json:"system_msg"`   //系统提示
}
