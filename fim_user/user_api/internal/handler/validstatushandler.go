package handler

import (
	"FIM/common/response"
	"FIM/fim_user/user_api/internal/logic"
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func valid_statusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendValidStatusRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewValid_statusLogic(r.Context(), svcCtx)
		resp, err := l.Valid_status(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
