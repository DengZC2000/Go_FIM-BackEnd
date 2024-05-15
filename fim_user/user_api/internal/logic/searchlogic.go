package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_user/user_models"
	"context"
	"fmt"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// 先找所有用户
	friends, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.UserConfModel{
		Online: req.Online,
	}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preloads: []string{"UserModel"},
		Joins:    "left join user_models um on um.id = user_conf_models.user_id",
		Where:    l.svcCtx.DB.Where("(user_conf_models.search_user <> 0 or user_conf_models.search_user is not null) and (user_conf_models.search_user = 1 and um.id = ?) or (user_conf_models.search_user = 2 and (um.id = ? or um.nick_name like ?))", req.Key, req.Key, fmt.Sprintf("%%%s%%", req.Key)),
	})
	var list []types.SearchInfo
	for _, friend := range friends {
		list = append(list, types.SearchInfo{
			UserID:   friend.UserID,
			Nickname: friend.UserModel.NickName,
			Profile:  friend.UserModel.Profile,
			Avatar:   friend.UserModel.Avatar,
			//IsFriend:,
		})
	}
	resp = &types.SearchResponse{Count: count, List: list}

	return
}
