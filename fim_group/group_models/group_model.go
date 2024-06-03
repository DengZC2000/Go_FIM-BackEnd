package group_models

import (
	"FIM/common/models"
	"FIM/common/models/ctype"
)

type GroupModel struct {
	models.Model
	Title                string                      `gorm:"size:32" json:"title"`        //群名
	Abstract             string                      `gorm:"size:128" json:"abstract"`    //简介
	Avatar               string                      `gorm:"size:256" json:"avatar"`      //群头像
	Creator              uint                        `json:"creator"`                     //群主
	IsSearch             bool                        `json:"is_search"`                   //是否可以搜索
	Verification         int8                        `json:"verification"`                //群验证,0 不允许任何人 1允许任何人 2 需要验证消息 3 需要回答问题 4 需要正确回答设置的问题
	VerificationQuestion *ctype.VerificationQuestion `json:"verification_question"`       //验证问题,不一定需要，验证方式为3和4的时候需要
	IsInvite             bool                        `json:"is_invite"`                   //允许邀请
	IsTemporarySession   bool                        `json:"is_temporary_session"`        //是否开启临时会话
	IsProhibition        bool                        `json:"is_prohibition"`              //是否开启全员禁言
	Size                 int                         `json:"size"`                        //群规模 20 100 1000 2000
	MemberList           []GroupMemberModel          `gorm:"foreignKey:GroupID" json:"-"` //群成员
}

func (uc GroupModel) GetQuestionCount() (count int) {
	if uc.VerificationQuestion != nil {
		if uc.VerificationQuestion.Answer1 != nil {
			count += 1
		}
		if uc.VerificationQuestion.Answer2 != nil {
			count += 1
		}
		if uc.VerificationQuestion.Answer3 != nil {
			count += 1
		}
	}
	return count
}
