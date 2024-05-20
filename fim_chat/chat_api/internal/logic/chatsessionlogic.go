package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_chat/chat_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"errors"
	"fmt"

	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Chat_sessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChat_sessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Chat_sessionLogic {
	return &Chat_sessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Data struct {
	SU         uint   `gorm:"column:sU"`
	RU         uint   `gorm:"column:rU"`
	MaxDate    string `gorm:"column:maxDate"`
	MaxPreview string `gorm:"column:maxPreview"`
}

func (l *Chat_sessionLogic) Chat_session(req *types.ChatSessionRequest) (resp *types.ChatSessionResponse, err error) {
	chatList, count, _ := list_query.ListQuery(l.svcCtx.DB, Data{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "maxDate desc",
		},
		Table: func() (string, any) {
			return "(?) as u", l.svcCtx.DB.Model(&chat_models.ChatModel{}).
				Select("least(send_user_id, rev_user_id) as sU",
					" greatest(send_user_id, rev_user_id) as rU",
					"max(created_at) as maxDate",
					"(select msg_preview from chat_models where (send_user_id = sU and rev_user_id = rU) or (send_user_id = rU and rev_user_id = sU) order by created_at desc limit 1 ) as maxPreview").
				Where("send_user_id = ? or rev_user_id = ?", req.UserID, req.UserID).
				Group("least(send_user_id, rev_user_id)").
				Group("greatest(send_user_id, rev_user_id)")
		},
	})
	var userIDList []uint32
	for _, data := range chatList {
		if data.RU != req.UserID {
			userIDList = append(userIDList, uint32(data.RU))
		}
		if data.SU != req.UserID {
			userIDList = append(userIDList, uint32(data.SU))
		}
	}
	response, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})
	if err != nil {
		logx.Error(err)
		return nil, errors.New("用户服务错误")
	}
	var list = make([]types.ChatSession, 0)
	for _, data := range chatList {
		s := types.ChatSession{
			CreateAt:   data.MaxDate,
			MsgPreview: data.MaxPreview,
		}
		fmt.Println(data)
		if data.RU != req.UserID {
			s.UserID = data.RU
			s.Avatar = response.UserInfo[uint32(s.UserID)].Avatar
			s.Nickname = response.UserInfo[uint32(s.UserID)].NickName
		}
		if data.SU != req.UserID {
			s.UserID = data.SU
			s.Avatar = response.UserInfo[uint32(s.UserID)].Avatar
			s.Nickname = response.UserInfo[uint32(s.UserID)].NickName
		}
		list = append(list, s)
	}
	resp = &types.ChatSessionResponse{
		List:  list,
		Count: count,
	}
	return
}