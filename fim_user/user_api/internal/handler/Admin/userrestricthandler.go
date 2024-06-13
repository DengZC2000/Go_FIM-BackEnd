package Admin

import (
	"FIM/common/response"
	"FIM/fim_user/user_api/internal/logic/Admin"
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func User_restrictHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRestrictResquest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewUser_restrictLogic(r.Context(), svcCtx)
		resp, err := l.User_restrict(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
