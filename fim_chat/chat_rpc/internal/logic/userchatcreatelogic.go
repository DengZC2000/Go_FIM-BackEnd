package logic

import (
	"FIM/common/models/ctype"
	"FIM/fim_chat/chat_models"
	"context"
	"encoding/json"

	"FIM/fim_chat/chat_rpc/internal/svc"
	"FIM/fim_chat/chat_rpc/types/chat_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserChatCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatCreateLogic {
	return &UserChatCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserChatCreateLogic) UserChatCreate(in *chat_rpc.UserChatRequest) (*chat_rpc.UserChatResponse, error) {
	var msg *ctype.Msg
	err := json.Unmarshal(in.Msg, &msg)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	var systemMsg *ctype.SystemMsg
	if systemMsg != nil {
		err := json.Unmarshal(in.SystemMsg, systemMsg)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
	}
	err = l.svcCtx.DB.Create(&chat_models.ChatModel{
		SendUserID: uint(in.SendUserId),
		RevUserID:  uint(in.RevUserId),
		MsgType:    msg.Type,
		MsgPreview: "",
		Msg:        msg,
		SystemMsg:  systemMsg,
	}).Error

	return &chat_rpc.UserChatResponse{}, err
}
