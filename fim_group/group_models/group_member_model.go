package group_models

import (
	"FIM/common/models"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type GroupMemberModel struct {
	models.Model
	GroupID         uint       `json:"group_id"` //群id
	GroupModel      GroupModel `gorm:"foreignKey:GroupID" json:"-"`
	UserID          uint       `json:"user_id"`                        //用户id
	MemberNickname  string     `gorm:"size:32" json:"member_nickname"` //群昵称
	Role            int        `json:"role"`                           //1 群主 2 管理员 3 普通用户
	ProhibitionTime *int       `json:"prohibition_time"`               //禁言时间，单位 min
}

func (gm GroupMemberModel) GetProhibitionTime(client *redis.Client, DB *gorm.DB) *int {
	if gm.ProhibitionTime == nil {
		return nil
	}
	t, err := client.TTL(context.Background(), fmt.Sprintf("prohibition__%d", gm.ID)).Result()
	fmt.Println(t, err)
	if err != nil || t == -2*time.Nanosecond {
		// 将表改回去
		DB.Model(&gm).Update("prohibition_time", nil)
		return nil
	}
	res := int(t / time.Minute)
	return &res
}
