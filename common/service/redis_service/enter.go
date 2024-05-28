package redis_service

import (
	"FIM/common/models/ctype"
	"FIM/fim_user/user_rpc/types/user_rpc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// GetUserBaseInfoByRedis 通过redis获取用户基本信息
func GetUserBaseInfoByRedis(client *redis.Client, UserRpc user_rpc.UsersClient, userID uint) (userInfo ctype.UserInfo, err error) {
	key := fmt.Sprintf("fim_server_user_%d", userID)
	str, err := client.Get(context.Background(), key).Result()
	fmt.Println(str)
	fmt.Println(err)
	if err != nil {
		//没找到
		userBaseResponse, err1 := UserRpc.UserBaseInfo(context.Background(), &user_rpc.UserBaseInfoRequest{
			UserId: uint32(userID),
		})
		if err1 != nil {
			err = err1
			return
		}
		userInfo.ID = userID
		userInfo.Avatar = userBaseResponse.Avatar
		userInfo.Nickname = userBaseResponse.NickName
		//设置进缓存
		byteData, _ := json.Marshal(userInfo)
		client.Set(context.Background(), key, string(byteData), time.Hour) //存一个小时的用户基本信息
		return userInfo, nil
	}
	userInfo = ctype.UserInfo{}
	err = json.Unmarshal([]byte(str), &userInfo)

	return
}
