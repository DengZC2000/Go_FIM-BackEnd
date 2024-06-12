package middleware

import (
	"FIM/common/log_stash"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Writer struct {
	http.ResponseWriter
	Body []byte
}

func (w *Writer) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}
func LogActionMiddleware(pusher *log_stash.Pusher) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ClientIP := httpx.GetRemoteAddr(r)

			// 设置入参
			pusher.SetRequest(r)
			pusher.SetHeaders(r)

			ctx := context.WithValue(r.Context(), "ClientIP", ClientIP)
			ctx = context.WithValue(ctx, "UserID", r.Header.Get("User-ID"))
			var nw = Writer{
				ResponseWriter: w,
			}
			next(&nw, r.WithContext(ctx))
			// 设置响应
			if pusher.GetResponse() {
				// 读响应体
				pusher.SetResponse(string(nw.Body))

			}
		}
	}
}

// LogMiddleware 自定义的中间件
func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := httpx.GetRemoteAddr(r)
		ctx := context.WithValue(r.Context(), "ClientIP", ClientIP)
		ctx = context.WithValue(ctx, "UserID", r.Header.Get("User-ID"))
		next(w, r.WithContext(ctx))
	}
}
