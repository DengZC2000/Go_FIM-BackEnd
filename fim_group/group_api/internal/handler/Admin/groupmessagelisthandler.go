package Admin

import (
	"FIM/common/response"
	"FIM/fim_group/group_api/internal/logic/Admin"
	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func Group_message_listHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupMessageListRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewGroup_message_listLogic(r.Context(), svcCtx)
		resp, err := l.Group_message_list(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
