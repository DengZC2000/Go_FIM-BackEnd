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
		if r.Header.Get("User-Role") != "1" {
			response.Response(r, w, nil, errors.New("用户操作鉴权失败"))
			return
		}
		next(w, r)
	}
}
