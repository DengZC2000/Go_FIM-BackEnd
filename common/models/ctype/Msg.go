package ctype

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type MsgType = int8

const (
	TextMsgType MsgType = iota + 1
	ImageMsgType
	VideoMsgType
	FileMsgType
	VoiceMsgType
	VoiceCallMsgType
	VideoCallMsgType
	WithdrawMsgType
	ReplyMsgType
	QuoteMsgType
	AtMsgType
	TipMsgType
)

type Msg struct {
	Type         MsgType       `json:"type"`                     //消息类型
	TextMsg      *TextMsg      `json:"text_msg,omitempty"`       //文本消息
	ImageMsg     *ImageMsg     `json:"image_msg,omitempty"`      //图片
	VideoMsg     *VideoMsg     `json:"video_msg,omitempty"`      //视频
	FileMsg      *FileMsg      `json:"file_msg,omitempty"`       //文件
	VoiceMsg     *VoiceMsg     `json:"voice_msg,omitempty"`      //语音
	VoiceCallMsg *VoiceCallMsg `json:"voice_call_msg,omitempty"` //语音通话
	VideoCallMsg *VideoCallMsg `json:"video_call_msg,omitempty"` //视频通话
	WithdrawMsg  *WithdrawMsg  `json:"withdraw_msg,omitempty"`   //撤回消息
	ReplyMsg     *ReplyMsg     `json:"reply_msg,omitempty"`      //回复消息
	QuoteMsg     *QuoteMsg     `json:"quote_msg,omitempty"`      //引用消息
	AtMsg        *AtMsg        `json:"at_msg,omitempty"`         //at消息，只有群聊有
	TipMsg       *TipMsg       `json:"tip_msg,omitempty"`        //提示消息，一般是不入库的
}

// Scan 取出来的时候的数据
func (c *Msg) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 入库的数据
func (c *Msg) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type TipMsg struct {
	Status  string `json:"status"` //error warning success
	Content string `json:"content"`
}
type TextMsg struct {
	Content string `json:"content"`
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
	Content   string `json:"content"`    //撤回的提示符
	MsgID     uint   `json:"msg_id"`     //需要撤回的消息id 入参必填
	OriginMsg *Msg   `json:"origin_msg"` //原消息
}
type ReplyMsg struct {
	MsgID         uint      `json:"msg_id"`  //消息id
	Content       string    `json:"content"` //回复的文本消息
	Msg           *Msg      `json:"msg"`
	UserID        uint      `json:"user_id"`         //被回复人的用户id
	UserNickName  string    `json:"user_nick_name"`  //被回复人的用户昵称
	OriginMsgDate time.Time `json:"origin_msg_date"` //原消息的时间
}
type QuoteMsg struct {
	MsgID         uint      `json:"msg_id"`  //消息id
	Content       string    `json:"content"` //回复的文本消息，目前只能回复图片
	Msg           *Msg      `json:"msg"`
	UserID        uint      `json:"user_id"`         //被回复人的用户id
	UserNickName  string    `json:"user_nick_name"`  //被回复人的用户昵称
	OriginMsgDate time.Time `json:"origin_msg_date"` //原消息的时间
}
type AtMsg struct {
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
	Msg     *Msg   `json:"msg"`
}
