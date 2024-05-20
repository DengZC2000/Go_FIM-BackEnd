package handler

import (
	"FIM/common/response"
	"FIM/fim_chat/chat_api/internal/logic"
	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func user_topHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserTopRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewUser_topLogic(r.Context(), svcCtx)
		resp, err := l.User_top(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
