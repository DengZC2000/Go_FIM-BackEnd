package handler

import (
	"FIM/common/response"
	"FIM/fim_user/user_api/internal/logic"
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func friend_listHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendListRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewFriend_listLogic(r.Context(), svcCtx)
		resp, err := l.Friend_list(&req)

		response.Response(r, w, resp, err)
	}
}
