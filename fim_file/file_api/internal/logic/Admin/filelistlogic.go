package Admin

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_file/file_models"
	"context"

	"FIM/fim_file/file_api/internal/svc"
	"FIM/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type File_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFile_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *File_listLogic {
	return &File_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *File_listLogic) File_list(req *types.FileListRequest) (resp *types.FileListResponse, err error) {
	list, count, _ := list_query.ListQuery(l.svcCtx.DB, file_models.FileModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Key:   req.Key,
		},
		Likes: []string{"file_name"},
	})
	resp = &types.FileListResponse{}
	resp.Count = int(count)
	for _, model := range list {
		resp.List = append(resp.List, types.FileListInfoResponse{
			FileName:  model.FileName,
			Size:      model.Size,
			Path:      model.Path,
			CreatedAt: model.CreatedAt.String(),
			ID:        model.ID,
			WebPath:   model.WebPath(),
		})
	}
	return
}
