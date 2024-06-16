package Admin

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/common/models/ctype"
	"FIM/fim_chat/chat_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"FIM/utils"
	"context"
	"errors"

	"FIM/fim_chat/chat_api/internal/svc"
	"FIM/fim_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Chat_admin_historyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChat_admin_historyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Chat_admin_historyLogic {
	return &Chat_admin_historyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ChatHistory struct {
	ID        uint             `json:"id"`
	SendUser  ctype.UserInfo   `json:"send_user"`
	RevUser   ctype.UserInfo   `json:"rev_user"`
	IsMe      bool             `json:"is_me"` //哪条消息是我发的
	CreateAt  string           `json:"create_at"`
	Msg       *ctype.Msg       `json:"msg"`
	SystemMsg *ctype.SystemMsg `json:"system_msg"`
}
type ChatHistoryResponse struct {
	List  []ChatHistory `json:"list"`
	Count int64         `json:"count"`
}

// Chat_admin_history 用户与用户的聊天记录
func (l *Chat_admin_historyLogic) Chat_admin_history(req *types.ChatAdminHistoryRequest) (resp *ChatHistoryResponse, err error) {

	chatList, count, _ := list_query.ListQuery(l.svcCtx.DB, chat_models.ChatModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
			Sort:  "created_at desc",
		},

		Where: l.svcCtx.DB.Where("(send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)",
			req.SendUserID, req.RevUserID, req.RevUserID, req.SendUserID),
	})
	var userIDList []uint32
	for _, chat := range chatList {
		userIDList = append(userIDList, uint32(chat.SendUserID))
		userIDList = append(userIDList, uint32(chat.RevUserID))
	}

	//去重
	userIDList = utils.DeduplicationList(userIDList)
	//去调用户服务的rpc方法，获取用户信息{用户id:用户信息}
	response, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &user_rpc.UserListInfoRequest{
		UserIdList: userIDList,
	})

	if err != nil {
		logx.Error(err.Error())
		return nil, errors.New("用户服务错误")
	}

	var list = make([]ChatHistory, 0)
	for _, chat := range chatList {
		sendUser := ctype.UserInfo{
			Nickname: response.UserInfo[uint32(chat.SendUserID)].NickName,
			Avatar:   response.UserInfo[uint32(chat.SendUserID)].Avatar,
			ID:       chat.SendUserID,
		}
		revUser := ctype.UserInfo{
			Nickname: response.UserInfo[uint32(chat.RevUserID)].NickName,
			Avatar:   response.UserInfo[uint32(chat.RevUserID)].Avatar,
			ID:       chat.RevUserID,
		}
		info := ChatHistory{
			ID:        chat.ID,
			CreateAt:  chat.CreatedAt.String(),
			Msg:       chat.Msg,
			SystemMsg: chat.SystemMsg,
			SendUser:  sendUser,
			RevUser:   revUser,
		}
		if info.SendUser.ID == req.RevUserID {
			info.IsMe = true
		}
		list = append(list, info)
	}
	resp = &ChatHistoryResponse{
		List:  list,
		Count: count,
	}

	return
}
