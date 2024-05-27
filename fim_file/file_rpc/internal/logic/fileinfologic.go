package logic

import (
	"FIM/fim_file/file_models"
	"context"
	"errors"
	"strings"

	"FIM/fim_file/file_rpc/internal/svc"
	"FIM/fim_file/file_rpc/types/file_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileInfoLogic {
	return &FileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileInfoLogic) FileInfo(in *file_rpc.FileInfoRequest) (*file_rpc.FileInfoResponse, error) {
	var fileModel file_models.FileModel
	err := l.svcCtx.DB.Take(&fileModel, "uid = ?", in.FileId).Error
	if err != nil {
		return nil, errors.New("文件不存在")
	}

	var tp string //文件类型
	nameList := strings.Split(fileModel.FileName, ".")
	if len(nameList) > 1 {
		tp = nameList[len(nameList)-1]
	}
	return &file_rpc.FileInfoResponse{
		FileName: fileModel.FileName,
		FileHash: fileModel.Hash,
		FileSize: fileModel.Size,
		FileType: tp,
	}, nil
}
