package log_stash

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type Pusher struct {
	LogType int8   `json:"log_type"` // 日志类型 2 操作日志 3 运行日志
	IP      string `json:"ip"`
	UserID  uint   `json:"user_id"`
	Level   string `json:"level"`
	Title   string `json:"title"`
	Content string `json:"content"` // 日志详情
	Service string `json:"service"` // 服务 记录微服务的名称
	Client  *kq.Pusher
}

// Push 为什么是指针 因为要改值
func (p *Pusher) Push(title string, content string) {
	p.Title = title
	p.Content = content

}
func (p *Pusher) Commit(ctx context.Context) {
	if p.Client == nil {
		return
	}
	var userID uint
	userIDs := ctx.Value("UserID")
	if userIDs != nil {
		ID, _ := strconv.Atoi(userIDs.(string))
		userID = uint(ID)
	}
	clientIP := ctx.Value("ClientIP").(string)
	p.IP = clientIP
	p.UserID = userID
	byteData, err := json.Marshal(p)
	if err != nil {
		logx.Error(err)
		return
	}
	err = p.Client.Push(string(byteData))
	if err != nil {
		logx.Error(err)
		return
	}
}
func NewActionPusher(client *kq.Pusher, serviceName string) *Pusher {
	return NewPusher(client, 2, "Action", serviceName)

}
func NewRuntimePusher(client *kq.Pusher, serviceName string) *Pusher {
	return NewPusher(client, 3, "Runtime", serviceName)
}

func NewPusher(client *kq.Pusher, LogType int8, level string, serviceName string) *Pusher {
	return &Pusher{
		LogType: LogType,
		Level:   level,
		Service: serviceName,
		Client:  client,
	}

}
