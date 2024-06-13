package log_stash

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Pusher struct {
	LogType    int8   `json:"log_type"` // 日志类型 2 操作日志 3 运行日志
	IP         string `json:"ip"`
	UserID     uint   `json:"user_id"`
	Level      string `json:"level"`
	Title      string `json:"title"`
	Content    string `json:"content"` // 日志详情
	Service    string `json:"service"` // 服务 记录微服务的名称
	Ctx        context.Context
	Client     *kq.Pusher
	Items      []string
	isRequest  bool
	isHeaders  bool
	isResponse bool
	request    string
	headers    string
	response   string
	Count      int
}

// PushInfo 为什么是指针 因为要改值
func (p *Pusher) PushInfo(title string) {
	p.Title = title
	p.Level = "Info"

}

// PushWarning 为什么是指针 因为要改值
func (p *Pusher) PushWarning(title string) {
	p.Title = title
	p.Level = "Warning"

}

// PushError 为什么是指针 因为要改值
func (p *Pusher) PushError(title string) {
	p.Title = title
	p.Level = "Error"

}
func (p *Pusher) IsRequest() {
	p.isRequest = true
}
func (p *Pusher) IsHeaders() {
	p.isHeaders = true
}
func (p *Pusher) IsResponse() {
	p.isResponse = true
}
func (p *Pusher) GetResponse() bool {
	return p.isResponse
}

// SetRequest 设置一组入参
func (p *Pusher) SetRequest(r *http.Request) {
	// 请求头
	// 请求体
	// 请求路径，请求方法
	// 关于请求体的问题，拿了之后要还回去
	// 一定要在参数绑定之前调用
	method := r.Method
	path := r.URL.String()
	byteData, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(byteData))
	p.request = fmt.Sprintf(`<div class="log_request">
<div class="log_request_head">
	<span class="log_request_method %s">%s</span>
	<span class="log_request_path">%s</span>
</div>
	<div class="log_request_body">
		<pre class="log_json_body">%s</pre>
	</div>
</div>`, strings.ToLower(method), method, path, string(byteData))

}
func (p *Pusher) SetHeaders(r *http.Request) {
	byteData, _ := json.Marshal(r.Header)
	p.headers = fmt.Sprintf(`<div class="log_request_header">
<div class="log_request_body">
   <pre class="log_json_body">%s</pre>
</div>
</div>`, string(byteData))
}
func (p *Pusher) SetResponse(w string) {
	fmt.Println("邓智超", w)

	p.response = fmt.Sprintf(`
<div class="log_response">
	<pre class="log_json_body">%s</pre>
</div>`, w)
	p.Commit(p.Ctx)
}
func (p *Pusher) SetItemInfo(label string, val any) {
	p.setItem("Info", label, val)
}

func (p *Pusher) SetItemWarning(label string, val any) {
	p.setItem("Warning", label, val)

}

func (p *Pusher) SetItemError(label string, val any) {
	p.setItem("Error", label, val)

}
func (p *Pusher) setItem(level, label string, val any) {
	p.Level = level
	var str string
	switch value := val.(type) {
	case string:
		str = fmt.Sprintf("<div class=\"log_item_label\">%s</div> <div class=\"log_item_content\">%s</div>", label, value)
	case int, uint, uint32, uint64, int32, int8:
		str = fmt.Sprintf("<div class=\"log_item_label\">%s</div> <div class=\"log_item_content\">%d</div>", label, value)
	default:
		byteData, _ := json.Marshal(val)
		str = fmt.Sprintf("<div class=\"log_item_label\">%s</div> <div class=\"log_item_content\">%s</div>", label, string(byteData))
	}
	logItem := fmt.Sprintf("<div class=\"log_item_%s\">%s</div>", level, str)
	p.Items = append(p.Items, logItem)

}
func (p *Pusher) SetCtx(ctx context.Context) {
	p.Ctx = ctx
}
func (p *Pusher) Commit(ctx context.Context) {
	if p.Ctx == nil {
		p.Ctx = ctx
	}
	if p.isResponse && p.Count == 0 {
		p.Count = 1
		return
	}
	if p.Client == nil {
		return
	}
	var items []string
	if p.isRequest {
		items = append(items, p.request)
	}
	if p.isHeaders {
		items = append(items, p.headers)
	}
	if p.isResponse {

		items = append(items, p.response)
	}
	items = append(items, p.Items...)
	for _, content := range items {
		p.Content += content
	}

	p.Items = []string{}
	var userID uint
	userIDs := p.Ctx.Value("UserID")
	if userIDs != nil {
		ID, _ := strconv.Atoi(userIDs.(string))
		userID = uint(ID)
	}
	clientIP := p.Ctx.Value("ClientIP").(string)
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
	p.Content = ""
}
func NewActionPusher(client *kq.Pusher, serviceName string) *Pusher {
	return NewPusher(client, 2, serviceName)

}
func NewRuntimePusher(client *kq.Pusher, serviceName string) *Pusher {
	return NewPusher(client, 3, serviceName)
}

func NewPusher(client *kq.Pusher, LogType int8, serviceName string) *Pusher {
	return &Pusher{
		LogType: LogType,
		Service: serviceName,
		Client:  client,
	}

}
