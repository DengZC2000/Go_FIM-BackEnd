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

// Info 为什么是指针 因为要改值
func (p *Pusher) Info(title string, content string) {
	p.Title = title
	p.Content = content
	p.Save()
}
func (p *Pusher) Save() {
	if p.Client == nil {
		return
	}
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
func NewActionPusher(ctx context.Context, client *kq.Pusher, serviceName string) *Pusher {
	return NewPusher(ctx, client, 2, "Action", serviceName)

}
func NewRuntimePusher(ctx context.Context, client *kq.Pusher, serviceName string) *Pusher {
	return NewPusher(ctx, client, 3, "Runtime", serviceName)
}

func NewPusher(ctx context.Context, client *kq.Pusher, LogType int8, level string, serviceName string) *Pusher {
	var userID uint
	userIDs := ctx.Value("UserID")
	if userIDs != nil {
		ID, _ := strconv.Atoi(userIDs.(string))
		userID = uint(ID)
	}
	clientIP := ctx.Value("ClientIP").(string)
	return &Pusher{
		IP:      clientIP,
		UserID:  userID,
		LogType: LogType,
		Level:   level,
		Service: serviceName,
		Client:  client,
	}

}
