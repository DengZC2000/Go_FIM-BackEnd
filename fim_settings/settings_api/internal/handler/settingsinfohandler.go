package handler

import (
	"FIM/common/response"
	"FIM/fim_settings/settings_api/internal/logic"
	"FIM/fim_settings/settings_api/internal/svc"
	"FIM/fim_settings/settings_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func settings_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SettingsInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewSettings_infoLogic(r.Context(), svcCtx)
		resp, err := l.Settings_info(&req)
		//if err != nil {
		//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//httpx.OkJsonCtx(r.Context(), w, resp)
		//}

		response.Response(r, w, resp, err)
	}
}
