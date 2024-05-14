package handler

import (
	"FIM/common/response"
	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"os"
	"path"
)

func image_showHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		filePath := path.Join("uploads", req.ImageType, req.ImageName)
		byteData, err := os.ReadFile(filePath)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		w.Write(byteData)

	}
}
