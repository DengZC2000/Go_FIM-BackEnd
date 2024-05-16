package logic

import (
	"FIM/fim_user/user_models"
	"context"
	"errors"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_validLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_validLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_validLogic {
	return &User_validLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_validLogic) User_valid(req *types.UserValidRequest) (resp *types.UserValidResponse, err error) {
	var userConf user_models.UserConfModel
	err = l.svcCtx.DB.Take(&userConf, "user_id = ?", req.FriendID).Error
	if err != nil {
		return nil, errors.New("没有此人")
	}
	var friend user_models.FriendModel
	if friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("你们已经是好友了~")
	}
	resp = &types.UserValidResponse{}
	resp.Verification = userConf.Verification
	switch userConf.Verification {
	case 0: //不允许任何人添加
	case 1: //允许任何人添加
		// 直接成为好友
		//先往验证表中添加一条数据，然后通过
	case 2: //需要验证问题

	case 3, 4: //需要正确回答问题
		if userConf.VerificationQuestion != nil {
			resp.VerificationQuestion = types.VerificationQuestion{
				Problem1: userConf.VerificationQuestion.Problem1,
				Problem2: userConf.VerificationQuestion.Problem2,
				Problem3: userConf.VerificationQuestion.Problem3,
			}
		}
	default:
	}
	return
}
