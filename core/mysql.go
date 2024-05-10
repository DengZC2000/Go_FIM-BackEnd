package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitGorm(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{
		// 设置Logger，启用所有日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("[%s] 连接失败！", MysqlDataSource))
		return nil
	}
	fmt.Println(fmt.Sprintf("[%s] mysql连接成功！", MysqlDataSource))
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
