package handler

import (
	"FIM/common/response"
	"FIM/fim_group/group_api/internal/logic"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func group_update_nicknameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupUpdateMemberNicknameRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewGroup_update_nicknameLogic(r.Context(), svcCtx)
		resp, err := l.Group_update_nickname(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
