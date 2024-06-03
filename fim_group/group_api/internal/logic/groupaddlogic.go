package logic

import (
	"FIM/common/models/ctype"
	"FIM/fim_group/group_models"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_addLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_addLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_addLogic {
	return &Group_addLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_addLogic) Group_add(req *types.AddGroupRequest) (resp *types.AddGroupResponse, err error) {
	var group group_models.GroupModel
	err = l.svcCtx.DB.Take(&group, req.GroupID).Error
	if err != nil {
		return nil, errors.New("群不存在")
	}

	var member group_models.GroupMemberModel
	err1 := l.svcCtx.DB.Take(&member, "group_id = ? and user_id = ?", req.GroupID, req.UserID).Error
	if err1 == nil {
		return nil, errors.New("你已经在群里了，请勿重复加群")
	}
	resp = &types.AddGroupResponse{}
	verifyModel := group_models.GroupVerifyModel{
		GroupID:            req.GroupID,
		UserID:             req.UserID,
		Status:             0,
		AdditionalMessages: req.AdditionalMessages,
		Type:               1,
	}
	switch group.Verification {
	case 0: //不允许任何人添加
		return nil, errors.New("不允许任何人加群")
	case 1: //允许任何人添加
		// 直接成为好友
		//先往验证表中添加一条数据，然后通过
		verifyModel.Status = 1
		//把用户加到群中
		var userGroup = group_models.GroupMemberModel{
			GroupID: req.GroupID,
			UserID:  req.UserID,
			Role:    3,
		}
		l.svcCtx.DB.Create(&userGroup)
	case 2: //需要验证问题

	case 3: //需要回答问题
		if req.VerificationQuestion != nil {
			verifyModel.VerificationQuestion = &ctype.VerificationQuestion{
				Problem1: group.VerificationQuestion.Problem1,
				Problem2: group.VerificationQuestion.Problem2,
				Problem3: group.VerificationQuestion.Problem3,
				Answer1:  req.VerificationQuestion.Answer1,
				Answer2:  req.VerificationQuestion.Answer2,
				Answer3:  req.VerificationQuestion.Answer3,
			}
		}
	case 4: //需要正确回答问题
		var count = 0
		if group.VerificationQuestion != nil && req.VerificationQuestion != nil {
			if group.VerificationQuestion.Answer1 != nil && req.VerificationQuestion.Answer1 != nil {
				if *group.VerificationQuestion.Answer1 == *req.VerificationQuestion.Answer1 {
					count += 1
				}
			}
			if group.VerificationQuestion.Answer2 != nil && req.VerificationQuestion.Answer2 != nil {
				if *group.VerificationQuestion.Answer2 == *req.VerificationQuestion.Answer2 {
					count += 1

				}
			}
			if group.VerificationQuestion.Answer3 != nil && req.VerificationQuestion.Answer3 != nil {
				if *group.VerificationQuestion.Answer3 == *req.VerificationQuestion.Answer3 {
					count += 1

				}
			}
			if count != group.GetQuestionCount() {
				return nil, errors.New("答案错误")
			}
			//直接加群
			verifyModel.Status = 1
			verifyModel.VerificationQuestion = group.VerificationQuestion
			//把用户加到群中
			var userGroup = group_models.GroupMemberModel{
				GroupID: req.GroupID,
				UserID:  req.UserID,
				Role:    3,
			}
			l.svcCtx.DB.Create(&userGroup)
		} else {
			return nil, errors.New("答案错误")

		}
	default:
	}
	err = l.svcCtx.DB.Create(&verifyModel).Error
	return
}
