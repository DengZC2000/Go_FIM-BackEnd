package handler

import (
	"FIM/common/response"
	"net/http"

	"FIM/fim_auth/auth_api/internal/logic"
	"FIM/fim_auth/auth_api/internal/svc"
)

func open_loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewOpen_loginLogic(r.Context(), svcCtx)
		resp, err := l.Open_login()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)
	}
}
