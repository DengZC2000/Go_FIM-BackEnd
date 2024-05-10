package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitRedis(addr, password string, db, PoolSize int) (client *redis.Client, err error) {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: PoolSize, //连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	if Rdb != nil {
		fmt.Println(fmt.Sprintf("[%s] redis连接成功！", addr))
	} else {
		return nil, errors.New("redis连接失败")
	}
	return Rdb, nil
}
