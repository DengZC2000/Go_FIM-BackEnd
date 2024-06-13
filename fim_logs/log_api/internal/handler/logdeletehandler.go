package handler

import (
	"FIM/common/response"
	"FIM/fim_logs/log_api/internal/logic"
	"FIM/fim_logs/log_api/internal/svc"
	"FIM/fim_logs/log_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func log_deleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewLog_deleteLogic(r.Context(), svcCtx)
		resp, err := l.Log_delete(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
