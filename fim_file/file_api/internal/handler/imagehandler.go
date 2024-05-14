package handler

import (
	"FIM/common/response"
	"FIM/fim_file/file_api/internal/logic"
	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"
	"FIM/utils"
	"FIM/utils/random"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func imageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		imageType := r.FormValue("imageType")
		if imageType == "" {
			response.Response(r, w, nil, errors.New("imageType不能为空"))
			return
		}
		file, fileHead, err := r.FormFile("image")
		if err != nil {
			response.Response(r, w, nil, errors.New(err.Error()))
			return
		}

		//文件大小限制
		mSize := float64(fileHead.Size) / float64(1024) / float64(1024)
		if mSize > svcCtx.Config.FileSize {
			response.Response(r, w, nil, errors.New(fmt.Sprintf("文件上传超过大小限制 %.2f MB", svcCtx.Config.FileSize)))
			return
		}
		//文件白名单限制
		nameList := strings.Split(fileHead.Filename, ".")
		if len(nameList) <= 1 {
			response.Response(r, w, nil, errors.New("文件非法"))
			return
		}
		if !utils.InList(svcCtx.Config.WriteList, nameList[len(nameList)-1]) {
			response.Response(r, w, nil, errors.New("上传的不是图片"))
			return
		}
		//文件重名处理
		dirPath := path.Join(svcCtx.Config.UpLoadDir, imageType)
		dir, err := os.ReadDir(dirPath)
		if err != nil {
			os.MkdirAll(dirPath, 0666)
		}
		filePath := path.Join(svcCtx.Config.UpLoadDir, imageType, fileHead.Filename)
		imageData, _ := io.ReadAll(file)
		//fileName := fileHead.Filename
		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)
		resp.Url = "/" + filePath
		if InDir(dir, fileHead.Filename) {
			byteData, _ := os.ReadFile(filePath)

			if utils.MD5(byteData) == utils.MD5(imageData) {
				//两个文件是一样的,名字一样，内容也一样，虚惊一场，直接返回
				response.Response(r, w, resp, nil)
				return
			}
			//否则此时名字相同，但是内容不同，其中一个需要改名
			prefix := utils.GetFilePrefix(filePath)
			filePath = prefix + "_" + random.RandStr(4) + "." + nameList[len(nameList)-1]
		}

		err = os.WriteFile(filePath, imageData, 0666)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		resp.Url = "/" + filePath
		response.Response(r, w, resp, err)

	}
}
func InDir(dir []os.DirEntry, file string) bool {
	for _, entry := range dir {
		if entry.Name() == file {
			return true
		}
	}
	return false
}
