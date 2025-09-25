package models

import (
	"fmt"
	"ginApi/common/config"
	"ginApi/common/enum"
	"ginApi/common/response"
	"github.com/go-redis/redis"
)

var RedisDb *redis.Client

func init() {
	addr := fmt.Sprintf("%v:%v", config.Viper.Get("redis.host"), config.Viper.Get("redis.port"))
	password := config.Viper.GetString("redis.password")
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	_, err = RedisDb.Ping().Result()
	if err != nil {
		fmt.Printf("redis error: %v\n", err)
		panic(&response.Response{
			Code: enum.CodeSystemError,
			Msg:  "redis服务器错误",
		})
	}
}
