package logic

import (
	"context"

	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Image_showLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImage_showLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Image_showLogic {
	return &Image_showLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Image_showLogic) Image_show(req *types.ImageShowRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
