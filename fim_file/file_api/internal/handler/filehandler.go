package handler

import (
	"FIM/common/response"
	"FIM/fim_file/file_api/internal/logic"
	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils"
	"FIM/utils/random"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func fileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileRequest
		if err := httpx.ParseHeaders(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		file, fileHead, err := r.FormFile("file")
		if err != nil {
			response.Response(r, w, nil, errors.New(err.Error()))
			return
		}

		//文件上传用黑名单限制

		//exe
		nameList := strings.Split(fileHead.Filename, ".")
		if len(nameList) <= 1 {
			//文件没有后缀名
			response.Response(r, w, nil, errors.New("文件非法"))
			return
		}
		if utils.InList(svcCtx.Config.BlackList, nameList[len(nameList)-1]) {
			response.Response(r, w, nil, errors.New("上传的文件非法"))
			return
		}
		//先去拿用户信息
		userResponse, err := svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
			UserIdList: []uint32{uint32(req.UserID)},
		})
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		dirName := fmt.Sprintf("%d_%s", req.UserID, userResponse.UserInfo[uint32(req.UserID)].NickName)

		//文件重名处理
		dirPath := path.Join(svcCtx.Config.UpLoadDir, "file", dirName)
		dir, err := os.ReadDir(dirPath)
		if err != nil {
			os.MkdirAll(dirPath, 0666)
		}
		filePath := path.Join(dirPath, fileHead.Filename)
		imageData, _ := io.ReadAll(file)
		//fileName := fileHead.Filename
		l := logic.NewFileLogic(r.Context(), svcCtx)
		resp, err := l.File(&req)
		resp.Src = "/" + filePath
		if utils.InDir(dir, fileHead.Filename) {
			//改名
			prefix := utils.GetFilePrefix(filePath)
			filePath = prefix + "_" + random.RandStr(4) + "." + nameList[len(nameList)-1]
		}

		err = os.WriteFile(filePath, imageData, 0666)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		resp.Src = "/" + filePath
		response.Response(r, w, resp, err)

	}
}
