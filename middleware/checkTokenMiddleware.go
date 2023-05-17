package middleware

import (
	"fmt"
	"ginApi/common/config"
	"ginApi/common/enum"
	"ginApi/common/myLogger"
	"ginApi/common/response"
	"ginApi/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CheckTokenMiddleware struct{}

// Handle 验证token中间件
func (this CheckTokenMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Viper.GetString("token.type") != "token" {
			c.Next()
			return
		}
		token := c.Request.Header.Get("X-token")
		myLogger.Println(fmt.Sprintf("X-token[%s]", token))
		if token == "" {
			panic(&response.Response{
				Code: enum.CodeTokenError,
				Msg:  enum.ErrMsg[enum.CodeTokenError],
			})
		}
		users, _ := models.RedisDb.HGetAll("token:" + token).Result()
		if len(users) == 0 {
			panic(&response.Response{
				Code: enum.CodeTokenError,
				Msg:  enum.ErrMsg[enum.CodeTokenError],
			})
			return
		}
		if value, ok := users["userId"]; ok {
			userId, _ := strconv.Atoi(value)
			c.Set("userId", userId)
		}
		c.Next()
	}
}
