package logic

import (
	"context"

	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileLogic {
	return &FileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileLogic) File(req *types.FileRequest) (resp *types.FileResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.FileResponse{}
	return
}
