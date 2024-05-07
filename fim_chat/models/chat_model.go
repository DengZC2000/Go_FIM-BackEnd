package models

import (
	"FIM/common/models"
	"time"
)

type ChatModel struct {
	models.Model
	SendUserID uint       `json:"send_user_id"` //发起验证方
	RevUserID  uint       `json:"rev_user_id"`  //接收验证方
	MsgType    int8       `json:"msg_type"`     //消息类型 1 文本 2 图片 3 视频 4 文件 5 语音 6 语音通话 7 视频通话 8 撤回消息 9 回复消息 10 引用消息
	MsgPreview string     `json:"msg_preview"`  //消息预览
	Msg        Msg        `json:"msg"`          //消息内容
	SystemMsg  *SystemMsg `json:"system_msg"`   //系统提示
}
type SystemMsg struct {
	Type int8 `json:"type"` //违规类型 1 涉黄 2 涉恐 3 涉政 4 不正当言论
}
type Msg struct {
	Type         int8          `json:"type"`           //消息类型
	Content      *string       `json:"content"`        //为1的时候使用
	ImageMsg     *ImageMsg     `json:"image_msg"`      //图片
	VideoMsg     *VideoMsg     `json:"video_msg"`      //视频
	FileMsg      *FileMsg      `json:"file_msg"`       //文件
	VoiceMsg     *VoiceMsg     `json:"voice_msg"`      //语音
	VoiceCallMsg *VoiceCallMsg `json:"voice_call_msg"` //语音通话
	VideoCallMsg *VideoCallMsg `json:"video_call_msg"` //视频通话
	WithdrawMsg  *WithdrawMsg  `json:"withdraw_msg"`   //撤回消息
	ReplyMsg     *ReplyMsg     `json:"reply_msg"`      //回复消息
	QuoteMsg     *QuoteMsg     `json:"quote_msg"`      //引用消息
}
type ImageMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}
type VideoMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` //时长，单位 s
}
type FileMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"` //文件大小
	Type  string `json:"type"` //文件类型
}
type VoiceMsg struct {
	Src  string `json:"src"`
	Time int    `json:"time"` //语音消息的时长,单位 s
}
type VoiceCallMsg struct {
	StartTime time.Time `json:"start_time"` //开始时间
	EndTime   time.Time `json:"end_time"`   //结束时间
	EndReason int8      `json:"end_reason"` //结束原因 0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}
type VideoCallMsg struct {
	StartTime time.Time `json:"start_time"` //开始时间
	EndTime   time.Time `json:"end_time"`   //结束时间
	EndReason int8      `json:"end_reason"` //结束原因 0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}
type WithdrawMsg struct {
	Content   string `json:"content"` //撤回的提示符
	OriginMsg *Msg   `json:"-"`       //原消息
}
type ReplyMsg struct {
	MsgID   uint   `json:"msg_id"`  //消息id
	Content string `json:"content"` //回复的文本消息，目前只能回复图片
	Msg     *Msg   `json:"msg"`
}
type QuoteMsg struct {
	MsgID   uint   `json:"msg_id"`  //消息id
	Content string `json:"content"` //回复的文本消息，目前只能回复图片
	Msg     *Msg   `json:"msg"`
}
