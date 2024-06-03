package main

import (
	"FIM/core"
	"FIM/fim_chat/chat_models"
	"FIM/fim_file/file_models"
	"FIM/fim_group/group_models"
	"FIM/fim_user/user_models"
	"flag"
	"fmt"
	"log"
)

type Options struct {
	DB bool
}

func main() {

	var opt Options
	flag.BoolVar(&opt.DB, "db", false, "是否只建立表结构")
	flag.Parse()

	if opt.DB {
		db := core.InitGorm("root:1310138359@tcp(127.0.0.1:3306)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local")
		err := db.AutoMigrate(
			&user_models.UserModel{},
			&user_models.UserConfModel{},
			&user_models.FriendModel{},
			&user_models.FriendVerifyModel{},

			&chat_models.ChatModel{},
			&chat_models.TopUserModel{},
			&chat_models.UserChatDeleteModel{},
			&group_models.GroupModel{},
			&group_models.GroupVerifyModel{},
			&group_models.GroupMsgModel{},
			&group_models.GroupMemberModel{},
			&group_models.GroupUserMsgDeleteModel{},

			&file_models.FileModel{},
		)
		if err != nil {
			log.Fatal("表结构生成失败！")
			return
		}
		fmt.Println("表结构生成成功！")
	}

}
