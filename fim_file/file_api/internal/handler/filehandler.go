package handler

import (
	"FIM/common/response"
	"FIM/fim_file/file_api/internal/logic"
	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"
	"FIM/fim_file/file_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
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

		//文件大小限制
		mSize := float64(fileHead.Size) / float64(1024) / float64(1024)
		if mSize > svcCtx.Config.FileSize {
			response.Response(r, w, nil, errors.New(fmt.Sprintf("文件上传超过大小限制 %.2f MB", svcCtx.Config.FileSize)))
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
		l := logic.NewFileLogic(r.Context(), svcCtx)
		resp, err := l.File(&req)
		fileData, _ := io.ReadAll(file)
		fileHash := utils.MD5(fileData)
		var fileModel file_models.FileModel
		err = svcCtx.DB.Take(&fileModel, "hash = ?", fileHash).Error
		if err == nil {
			//已经有这个文件了
			resp.Src = fileModel.WebPath()
			response.Response(r, w, resp, nil)
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

		dirPath := path.Join(svcCtx.Config.UpLoadDir, "file", dirName)
		_, err = os.ReadDir(dirPath)
		if err != nil {
			os.MkdirAll(dirPath, 0666)
		}
		UID := uuid.New()
		newfileModel := file_models.FileModel{
			UserID:   req.UserID,
			FileName: fileHead.Filename,
			Size:     fileHead.Size,
			Path:     path.Join(dirPath, fmt.Sprintf("%s.%s", UID, nameList[len(nameList)-1])),
			Hash:     fileHash,
			Uid:      UID,
		}

		err = os.WriteFile(newfileModel.Path, fileData, 0666)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		err = svcCtx.DB.Create(&newfileModel).Error
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		resp.Src = newfileModel.WebPath()
		response.Response(r, w, resp, err)

	}
}
