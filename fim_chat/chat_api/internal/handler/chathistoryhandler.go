package handler

import (
	"FIM/common/response"
	"FIM/fim_chat/chat_api/internal/logic"
	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func chat_historyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatHistoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewChat_historyLogic(r.Context(), svcCtx)
		resp, err := l.Chat_history(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
