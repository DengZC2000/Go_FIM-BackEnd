package logic

import (
	"FIM/common/list_query"
	"FIM/common/models"
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"FIM/fim_user/user_models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_valid_listLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_valid_listLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_valid_listLogic {
	return &User_valid_listLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_valid_listLogic) User_valid_list(req *types.FriendValidResquest) (resp *types.FriendValidResponse, err error) {
	fvls, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.FriendVerifyModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Where:    l.svcCtx.DB.Where("send_user_id = ? or rev_user_id", req.UserID, req.UserID),
		Preloads: []string{"RevUserModel.UserConfModel", "SendUserModel.UserConfModel"},
	})
	var list []types.FriendValidInfo
	for _, fvl := range fvls {
		info := types.FriendValidInfo{
			AdditionalMessages: fvl.AdditionalMessages,
			ID:                 fvl.ID,
			CreatedAt:          fvl.CreatedAt.String(),
		}
		if fvl.SendUserID == req.UserID {
			//我是发送方
			info.UserID = fvl.RevUserID
			info.Nickname = fvl.RevUserModel.NickName
			info.Avatar = fvl.RevUserModel.Avatar
			info.Verification = fvl.RevUserModel.UserConfModel.Verification
			info.Flag = "send"
			info.Status = fvl.SendStatus
		}
		if fvl.RevUserID == req.UserID {
			//我是接收方
			info.UserID = fvl.SendUserID
			info.Nickname = fvl.SendUserModel.NickName
			info.Avatar = fvl.SendUserModel.Avatar
			info.Verification = fvl.SendUserModel.UserConfModel.Verification
			info.Flag = "rev"
			info.Status = fvl.RevStatus
		}
		if fvl.VerificationQuestion != nil {
			info.VerificationQuestion = &types.VerificationQuestion{
				Problem1: fvl.VerificationQuestion.Problem1,
				Problem2: fvl.VerificationQuestion.Problem2,
				Problem3: fvl.VerificationQuestion.Problem3,
				Answer1:  fvl.VerificationQuestion.Answer1,
				Answer2:  fvl.VerificationQuestion.Answer2,
				Answer3:  fvl.VerificationQuestion.Answer3,
			}
		}
		list = append(list, info)
	}
	resp = &types.FriendValidResponse{List: list, Count: count}
	return
}
