package logic

import (
	"FIM/fim_group/group_models"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_validLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_validLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_validLogic {
	return &Group_validLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_validLogic) Group_valid(req *types.GroupValidRequest) (resp *types.GroupValidResponse, err error) {
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
	resp = &types.GroupValidResponse{}
	resp.Verification = group.Verification
	switch group.Verification {
	case 0: //不允许任何人添加
	case 1: //允许任何人添加
		// 直接成为好友
		//先往验证表中添加一条数据，然后通过
	case 2: //需要验证问题

	case 3, 4: //需要正确回答问题
		if group.VerificationQuestion != nil {
			resp.VerificationQuestion = types.VerificationQuestion{
				Problem1: group.VerificationQuestion.Problem1,
				Problem2: group.VerificationQuestion.Problem2,
				Problem3: group.VerificationQuestion.Problem3,
			}
		}
	default:
	}
	return
}
