package handler

import (
	"FIM/common/response"
	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"
	"FIM/fim_file/file_models"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"os"
)

func file_showHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		var fileModel file_models.FileModel
		err := svcCtx.DB.Take(&fileModel, "uid = ?", req.FileName).Error
		if err != nil {
			response.Response(r, w, nil, errors.New("文件不存在"))
			return
		}
		byteData, err := os.ReadFile(fileModel.Path)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		w.Write(byteData)
	}
}
