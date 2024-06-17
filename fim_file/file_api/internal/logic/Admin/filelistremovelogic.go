package Admin

import (
	"FIM/fim_file/file_models"
	"context"

	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type File_list_removeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFile_list_removeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *File_list_removeLogic {
	return &File_list_removeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *File_list_removeLogic) File_list_remove(req *types.FileListRemoveRequest) (resp *types.FileListRemoveResponse, err error) {
	var fileList []file_models.FileModel
	l.svcCtx.DB.Find(&fileList, "id in ? ", req.IDList).Delete(&fileList)
	logx.Infof("删除文件个数 %s", len(fileList))
	return
}
