package logic

import (
	"FIM/common/models/ctype"
	"FIM/fim_group/group_models"
	"FIM/utils/maps"
	"context"
	"errors"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_updateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_updateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_updateLogic {
	return &Group_updateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_updateLogic) Group_update(req *types.GroupUpdateRequest) (resp *types.GroupUpdateResponse, err error) {
	// 只能是群主或者管理员才能调用
	var groupMember group_models.GroupMemberModel
	err = l.svcCtx.DB.Preload("GroupModel").Take(&groupMember, "user_id = ? and group_id = ?", req.UserID, req.ID).Error
	if err != nil {
		logx.Error(err)
		return nil, errors.New("群不存在或者用户不是该群成员")
	}
	if !(groupMember.Role == 1 || groupMember.Role == 2) {
		return nil, errors.New("只能群主和管理员才能修改群信息")
	}
	groupMaps := maps.RefToMap(*req, "conf")
	if len(groupMaps) != 0 {
		// 单独处理 verificationQuestion
		// 取出maps中的 verificationQuestion ，再把maps中的 verificationQuestion 删掉
		// 利用取出的 verificationQuestion 去更新这个字段
		// 随后用剩余的maps中的项更新其他字段
		verificationQuestion, ok := groupMaps["verification_question"]
		if ok {
			delete(groupMaps, "verification_question")
			data := ctype.VerificationQuestion{}
			maps.MapToStruct(verificationQuestion.(map[string]any), &data)
			l.svcCtx.DB.Model(&groupMember.GroupModel).Updates(&group_models.GroupModel{
				VerificationQuestion: &data,
			})
		}
		err = l.svcCtx.DB.Model(&groupMember.GroupModel).Updates(groupMaps).Error
		if err != nil {
			logx.Error(err)
			return nil, errors.New("群信息更新失败")
		}
	}
	return
}
