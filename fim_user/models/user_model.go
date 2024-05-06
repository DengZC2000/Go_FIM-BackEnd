package models

import "FIM/common/models"

type UserModel struct {
	models.Model
	NickName string `json:"nick_name"`
	Password string `json:"password"`
	Profile  string `json:"profile"` //个人简介
	Avatar   string `json:"avatar"`  //头像
}
