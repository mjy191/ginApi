package Models

import (
	"ginApi/Common/Enum"
	"github.com/go-redis/redis"
)

var RedisDb *redis.Client

func init() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = RedisDb.Ping().Result()
	if err != nil {
		panic(map[string]interface{}{
			"code": Enum.CodeSystemError,
			"msg":  "redis服务器错误",
		})
	}
}
