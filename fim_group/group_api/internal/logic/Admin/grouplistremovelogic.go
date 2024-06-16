package Admin

import (
	"FIM/fim_group/group_models"
	"context"
	"gorm.io/gorm"

	"FIM/fim_group/group_api/internal/svc"
	"FIM/fim_group/group_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Group_list_removeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroup_list_removeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Group_list_removeLogic {
	return &Group_list_removeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Group_list_removeLogic) Group_list_remove(req *types.GroupListRemoveRequest) (resp *types.GroupListRemoveResponse, err error) {
	var groupList []group_models.GroupModel
	l.svcCtx.DB.Preload("MemberList").Preload("GroupMsgList").Find(&groupList, "id in ?", req.IdList)
	for _, model := range groupList {
		logx.Infof("开始删除id为 %d 的群,群昵称 %s", model.ID, model.Title)
		err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
			if len(model.GroupMsgList) > 0 {
				err = tx.Delete(&model.GroupMsgList).Error
				if err != nil {
					return err
				}
			}

			logx.Infof("删除群消息数量: %d条", len(model.GroupMsgList))
			if len(model.MemberList) > 0 {
				err = tx.Delete(&model.MemberList).Error
				if err != nil {
					return err
				}
			}

			logx.Infof("删除群成员数量: %d条", len(model.MemberList))

			var topModelList []group_models.GroupUserTopModel
			err = tx.Find(&topModelList, "group_id = ?", model.ID).Delete(&topModelList).Error
			if err != nil {
				return err
			}
			logx.Infof("删除群置顶数量: %d条", len(topModelList))

			var verifyList []group_models.GroupVerifyModel
			err = tx.Find(&verifyList, "group_id = ?", model.ID).Delete(&verifyList).Error
			if err != nil {
				return err
			}
			logx.Infof("删除群验证数量: %d条", len(verifyList))

			var userDeleteMsgList []group_models.GroupUserMsgDeleteModel
			err = tx.Find(&userDeleteMsgList, "group_id = ?", model.ID).Delete(&userDeleteMsgList).Error
			if err != nil {
				return err
			}
			logx.Infof("删除群用户删除聊天记录数量: %d条", len(userDeleteMsgList))
			err = tx.Delete(&model).Error
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			logx.Error(err)
			continue
		}
		logx.Infof("删除成功id为 %d 的群", model.ID)
	}
	return
}
