package logic

import (
	"context"

	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type File_showLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFile_showLogic(ctx context.Context, svcCtx *svc.ServiceContext) *File_showLogic {
	return &File_showLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *File_showLogic) File_show(req *types.FileShowRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
