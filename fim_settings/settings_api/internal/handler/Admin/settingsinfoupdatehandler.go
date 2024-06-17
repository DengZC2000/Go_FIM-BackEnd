package Admin

import (
	"FIM/common/response"
	"FIM/fim_settings/settings_api/internal/logic/Admin"
	"FIM/fim_settings/settings_api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func Settings_info_updateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Admin.MySettingsInfoUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := Admin.NewSettings_info_updateLogic(r.Context(), svcCtx)
		resp, err := l.Settings_info_update(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
