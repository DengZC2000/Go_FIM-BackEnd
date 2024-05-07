package main

import (
	"FIM/core"
	models2 "FIM/fim_chat/models"
	models3 "FIM/fim_group/models"
	"FIM/fim_user/models"
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
		db := core.InitGorm()
		err := db.AutoMigrate(
			&models.UserModel{},
			&models.UserConfModel{},
			&models.FriendModel{},
			&models.FriendVerifyModel{},

			&models2.ChatModel{},

			&models3.GroupModel{},
			&models3.GroupVerifyModel{},
			&models3.GroupMsgModel{},
			&models3.GroupMemberModel{},
		)
		if err != nil {
			log.Fatal("表结构生成失败！")
			return
		}
		fmt.Println("表结构生成成功！")
	}

}
