package handler

import (
	"FIM/common/response"
	"FIM/fim_file/file_api/internal/logic"
	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"
	"FIM/fim_file/file_models"
	"FIM/utils"
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

func imageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		imageType := r.FormValue("imageType")

		switch imageType {
		case "avatar", "group_avatar", "chat":
		case "":
			response.Response(r, w, nil, errors.New("imageType不能为空"))
			return
		default:
			response.Response(r, w, nil, errors.New("imageType只能为avatar,group_avatar,chat"))
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
		//先去算hash
		imageData, _ := io.ReadAll(file)
		imageHash := utils.MD5(imageData)
		var fileModel file_models.FileModel
		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)
		err = svcCtx.DB.Take(&fileModel, "hash = ?", imageHash).Error
		if err == nil {
			//找到了，有hash一摸一样的
			resp.Url = fileModel.WebPath()
			response.Response(r, w, resp, nil)
			return
		}
		// 拼路径 /uploads/imageType/{uid}.{后缀}
		dirPath := path.Join(svcCtx.Config.UpLoadDir, imageType)
		_, err = os.ReadDir(dirPath)
		if err != nil {
			os.MkdirAll(dirPath, 0666)
		}
		UID := uuid.New()
		filePath := path.Join(dirPath, fmt.Sprintf("%s.%s", UID, nameList[len(nameList)-1]))

		//fileName := fileHead.Filename

		err = os.WriteFile(filePath, imageData, 0666)
		if err != nil {
			response.Response(r, w, nil, err)
			return
		}
		newfileModel := file_models.FileModel{
			UserID:   req.UserID,
			FileName: fileHead.Filename,
			Size:     fileHead.Size,
			Path:     filePath,
			Hash:     utils.MD5(imageData),
			Uid:      UID,
		}
		//图片信息入库
		err = svcCtx.DB.Create(&newfileModel).Error
		if err != nil {
			logx.Error(err)
			response.Response(r, w, nil, err)
			return
		}
		resp.Url = newfileModel.WebPath()
		response.Response(r, w, resp, err)

	}
}
