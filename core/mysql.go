package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	dsn := "root:1310138359@tcp(127.0.0.1:3306)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("%s 连接失败", dsn))
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
