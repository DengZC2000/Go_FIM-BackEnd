package middleware

import (
	"FIM/common/response"
	"errors"
	"net/http"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 必须role为1的才能调用
		if r.Header.Get("User-Role") != "1" {
			response.Response(r, w, nil, errors.New("日志鉴权失败"))
			return
		}
		next(w, r)
	}
}
