package handler

import (
	"FIM/common/response"
	"FIM/fim_auth/auth_api/internal/logic"
	"FIM/fim_auth/auth_api/internal/svc"

	"net/http"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		resp, err := l.Authentication(token)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)
	}
}
