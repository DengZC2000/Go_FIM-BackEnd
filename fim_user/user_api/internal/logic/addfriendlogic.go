package logic

import (
	"FIM/common/models/ctype"
	"FIM/fim_user/user_models"
	"context"
	"errors"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Add_friendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdd_friendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Add_friendLogic {
	return &Add_friendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Add_friendLogic) Add_friend(req *types.AddFriendRequest) (resp *types.AddFriendResponse, err error) {
	var userConf user_models.UserConfModel
	var userInfo user_models.UserModel
	err = l.svcCtx.DB.Preload("UserConfModel").Take(&userInfo, req.UserID).Error
	if err != nil {
		return nil, errors.New("用户不存在，非法操作")
	}
	if userInfo.UserConfModel.RestrictAddUser {
		return nil, errors.New("该用户被限制加好友")
	}
	err = l.svcCtx.DB.Take(&userConf, "user_id = ?", req.FriendID).Error
	if err != nil {
		return nil, errors.New("没有此人")
	}
	var friend user_models.FriendModel
	if friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("你们已经是好友了~")
	}
	resp = &types.AddFriendResponse{}
	var verifyModel = user_models.FriendVerifyModel{
		SendUserID:         req.UserID,
		RevUserID:          req.FriendID,
		AdditionalMessages: req.AdditionalMessages,
	}

	switch userConf.Verification {
	case 0: //不允许任何人添加
		return nil, errors.New("该用户不允许任何人添加好友！")
	case 1: //允许任何人添加
		verifyModel.RevStatus = 1
		//加好友
		var userFriend = user_models.FriendModel{
			SendUserID: req.UserID,
			RevUserID:  req.FriendID,
		}
		l.svcCtx.DB.Create(&userFriend)
	case 2: //需要验证问题
		// verifyModel.Status = 0
	case 3:
		//需要回答问题
		if req.VerificationQuestion != nil {
			verifyModel.VerificationQuestion = &ctype.VerificationQuestion{
				Problem1: req.VerificationQuestion.Problem1,
				Problem2: req.VerificationQuestion.Problem2,
				Problem3: req.VerificationQuestion.Problem3,
				Answer1:  req.VerificationQuestion.Answer1,
				Answer2:  req.VerificationQuestion.Answer2,
				Answer3:  req.VerificationQuestion.Answer3,
			}
		}
	case 4: //需要正确回答问题
		var count = 0
		if userConf.VerificationQuestion != nil && req.VerificationQuestion != nil {
			if userConf.VerificationQuestion.Answer1 != nil && req.VerificationQuestion.Answer1 != nil {
				if *userConf.VerificationQuestion.Answer1 == *req.VerificationQuestion.Answer1 {
					count += 1
				}
			}
			if userConf.VerificationQuestion.Answer2 != nil && req.VerificationQuestion.Answer2 != nil {
				if *userConf.VerificationQuestion.Answer2 == *req.VerificationQuestion.Answer2 {
					count += 1

				}
			}
			if userConf.VerificationQuestion.Answer3 != nil && req.VerificationQuestion.Answer3 != nil {
				if *userConf.VerificationQuestion.Answer3 == *req.VerificationQuestion.Answer3 {
					count += 1

				}
			}

		}
		if count != userConf.GetQuestionCount() {
			return nil, errors.New("答案错误")
		}
		//直接加好友
		verifyModel.RevStatus = 1
		verifyModel.VerificationQuestion = userConf.VerificationQuestion
		//加好友
		var userFriend = user_models.FriendModel{
			SendUserID: req.UserID,
			RevUserID:  req.FriendID,
		}
		l.svcCtx.DB.Create(&userFriend)
	default:
		return nil, errors.New("不支持的验证参数")
	}
	err = l.svcCtx.DB.Create(&verifyModel).Error
	if err != nil {
		logx.Error(err.Error())
		return nil, errors.New("添加好友失败")
	}
	return
}
