package middleware

import (
	"FIM/common/log_stash"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LogActionMiddleware(pusher *log_stash.Pusher) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ClientIP := httpx.GetRemoteAddr(r)

			// 设置入参
			pusher.SetRequest(r)
			pusher.SetHeaders(r)

			ctx := context.WithValue(r.Context(), "ClientIP", ClientIP)
			ctx = context.WithValue(ctx, "UserID", r.Header.Get("User-ID"))
			next(w, r.WithContext(ctx))
			// 设置响应
			pusher.SetResponse(w)
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
