package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// LogMiddleware 自定义的中间件
func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ClientIP := httpx.GetRemoteAddr(r)
		ctx := context.WithValue(r.Context(), "ClientIP", ClientIP)
		ctx = context.WithValue(ctx, "UserID", r.Header.Get("User-ID"))
		next(w, r.WithContext(ctx))
	}
}
