package logic

import (
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"FIM/fim_user/user_models"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_infoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_infoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_infoLogic {
	return &User_infoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_infoLogic) User_info(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	userInfo, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: uint32(req.UserID),
	})
	if err != nil {
		return nil, err
	}
	var user user_models.UserModel
	err = json.Unmarshal(userInfo.Data, &user)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据错误")
	}

	fmt.Println(string(userInfo.Data))
	resp = &types.UserInfoResponse{
		UserID:        user.ID,
		Nickname:      user.NickName,
		Profile:       user.Profile,
		Avatar:        user.Avatar,
		RecallMessage: user.UserConfModel.RecallMessage,
		FriendOnline:  user.UserConfModel.FriendOnline,
		Sound:         user.UserConfModel.Sound,
		SecureLink:    user.UserConfModel.SecureLink,
		SavePwd:       user.UserConfModel.SavePwd,
		SearchUser:    user.UserConfModel.SearchUser,
		Verification:  user.UserConfModel.Verification,
	}
	if user.UserConfModel.VerificationQuestion != nil {
		resp.VerificationQuestion = &types.VerificationQuestion{
			Problem1: user.UserConfModel.VerificationQuestion.Problem1,
			Problem2: user.UserConfModel.VerificationQuestion.Problem2,
			Problem3: user.UserConfModel.VerificationQuestion.Problem3,
			Answer1:  user.UserConfModel.VerificationQuestion.Answer1,
			Answer2:  user.UserConfModel.VerificationQuestion.Answer2,
			Answer3:  user.UserConfModel.VerificationQuestion.Answer3,
		}
	}
	return resp, nil
}
