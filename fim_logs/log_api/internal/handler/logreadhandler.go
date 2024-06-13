package handler

import (
	"FIM/common/response"
	"FIM/fim_logs/log_api/internal/logic"
	"FIM/fim_logs/log_api/internal/svc"
	"FIM/fim_logs/log_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func log_readHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogReadRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewLog_readLogic(r.Context(), svcCtx)
		resp, err := l.Log_read(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
