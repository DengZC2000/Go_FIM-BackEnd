package user_models

import (
	"FIM/common/models"
	"FIM/common/models/ctype"
)

type UserConfModel struct {
	models.Model
	UserID               uint                        `json:"user_id"`
	UserModel            UserModel                   `gorm:"foreignKey:UserID" json:"-"`
	RecallMessage        *string                     `gorm:"size:128" json:"recall_message"` //撤回消息的提示内容,xxx撤回了一条消息并亲了你一下
	FriendOnline         bool                        `json:"friend_online"`                  //好友上线提示
	Sound                bool                        `json:"sound"`                          //声音
	SecureLink           bool                        `json:"secure_link"`                    //安全链接
	SavePwd              bool                        `json:"save_pwd"`                       //是否记住密码
	SearchUser           int8                        `json:"search_user"`                    //允许别人查找到你的方式,0:不允许别人查找到我，1：通过用户号找到我 2：可以通过昵称搜索找到我
	Verification         int8                        `json:"verification"`                   //好友验证,0 不允许任何人 1允许任何人 2 需要验证消息 3 需要回答问题 4 需要正确回答设置的问题
	VerificationQuestion *ctype.VerificationQuestion `json:"friend_question"`                //验证问题,不一定需要，验证方式为3和4的时候需要
	Online               bool                        `json:"online"`                         //是否在线
}
