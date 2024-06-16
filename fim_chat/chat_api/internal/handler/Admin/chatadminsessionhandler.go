package Admin

import (
	"FIM/common/response"
	"FIM/fim_chat/chat_api/internal/logic/Admin"
	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func Chat_admin_sessionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatAdminSessionRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewChat_admin_sessionLogic(r.Context(), svcCtx)
		resp, err := l.Chat_admin_session(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
