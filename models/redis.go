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
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err = RedisDb.Ping().Result()
	if err != nil {
		panic(response.Response{
			Code: enum.CodeSystemError,
			Msg:  "redis服务器错误",
		})
	}
}
