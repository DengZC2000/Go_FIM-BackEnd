package logic

import (
	"FIM/common/models/ctype"
	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"
	"FIM/fim_user/user_models"
	"FIM/utils/maps"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_updateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_updateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_updateLogic {
	return &User_updateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_updateLogic) User_update(req *types.UserInfoUpdateRequest) (resp *types.UserInfoUpdateResponse, err error) {

	fmt.Println(req.UserID)

	var user user_models.UserModel
	userMaps := maps.RefToMap(*req, "user")
	if len(userMaps) != 0 {
		err = l.svcCtx.DB.Take(&user, req.UserID).Error
		if err != nil {
			logx.Error("没有此人")
			return nil, errors.New("没有此人")
		}
		err = l.svcCtx.DB.Model(&user).Updates(userMaps).Error
		if err != nil {
			logx.Error("更新用户信息失败")
			return nil, errors.New("更新用户信息失败")
		}
	}
	userConfMaps := maps.RefToMap(*req, "user_conf")
	var userConf user_models.UserConfModel
	if len(userConfMaps) != 0 {
		err = l.svcCtx.DB.Take(&userConf, req.UserID).Error
		if err != nil {
			logx.Error("没有此人")
			return nil, errors.New("没有此人")
		}
	}
	verificationQuestion, ok := userConfMaps["verification_question"]
	if ok {
		data := ctype.VerificationQuestion{}
		delete(userConfMaps, "verification_question")
		maps.MapToStruct(verificationQuestion.(map[string]any), &data)
		if val, ok := verificationQuestion.(map[string]any)["problem1"]; ok {
			s := val.(string)
			data.Problem1 = &s
		}
		if val, ok := verificationQuestion.(map[string]any)["problem2"]; ok {
			s := val.(string)
			data.Problem2 = &s
		}
		if val, ok := verificationQuestion.(map[string]any)["problem3"]; ok {
			s := val.(string)
			data.Problem3 = &s
		}
		if val, ok := verificationQuestion.(map[string]any)["answer1"]; ok {
			s := val.(string)
			data.Answer1 = &s
		}
		if val, ok := verificationQuestion.(map[string]any)["answer2"]; ok {
			s := val.(string)
			data.Answer2 = &s
		}
		if val, ok := verificationQuestion.(map[string]any)["answer3"]; ok {
			s := val.(string)
			data.Answer3 = &s
		}
		err = l.svcCtx.DB.Model(&userConf).Updates(&user_models.UserConfModel{VerificationQuestion: &data}).Error
		if err != nil {
			logx.Error("更新用户配置Q&A信息失败")
			return nil, errors.New("更新用户配置Q&A信息失败")
		}
	}
	err = l.svcCtx.DB.Model(&userConf).Updates(userConfMaps).Error
	if err != nil {
		logx.Error("更新用户配置信息失败")
		return nil, errors.New("更新用户配置信息失败")
	}

	return
}
