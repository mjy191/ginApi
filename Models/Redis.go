package Models

import (
	"fmt"
	"ginApi/Common/Enum"
	"ginApi/Common/Tools"
	"github.com/go-redis/redis"
)

var RedisDb *redis.Client

func init() {
	addr := fmt.Sprintf("%v:%v", Tools.Config.Get("redis.host"), Tools.Config.Get("redis.port"))
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     addr,
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
