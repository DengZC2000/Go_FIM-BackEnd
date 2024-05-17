package logic

import (
	"FIM/fim_user/user_models"
	"context"
	"errors"

	"FIM/fim_user/user_api/internal/svc"
	"FIM/fim_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Valid_statusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValid_statusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Valid_statusLogic {
	return &Valid_statusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Valid_statusLogic) Valid_status(req *types.FriendValidStatusRequest) (resp *types.FriendValidStatusResponse, err error) {
	var friendVerify user_models.FriendVerifyModel
	//我要操作状态，我得是接收者
	err = l.svcCtx.DB.Take(&friendVerify, "id = ? and rev_user_id = ?", req.VerifyID, req.UserID).Error
	if err != nil {
		return nil, errors.New("验证记录不存在")
	}
	if friendVerify.Status != 0 {
		return nil, errors.New("不可更改状态")
	}
	switch req.Status {
	case 1: //同意
		friendVerify.Status = 1
		//往好友表里去加
		l.svcCtx.DB.Create(&user_models.FriendModel{
			SendUserID: friendVerify.SendUserID,
			RevUserID:  friendVerify.RevUserID,
		})
	case 2: //拒绝
		friendVerify.Status = 2
	case 3: //忽略
		friendVerify.Status = 3
	case 4: //删除
		//一条验证记录，是两个人看的
		l.svcCtx.DB.Delete(&friendVerify)
		return nil, nil
	}
	l.svcCtx.DB.Save(&friendVerify)

	return
}
