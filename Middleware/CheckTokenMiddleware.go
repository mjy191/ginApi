package Middleware

import (
	"fmt"
	"ginApi/Common/Enum"
	"ginApi/Common/Logger"
	"ginApi/Common/Tools"
	"ginApi/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CheckTokenMiddleware struct{}

// Handle 验证token中间件
func (this CheckTokenMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Tools.Config.GetString("token.type") != "token" {
			c.Next()
			return
		}
		token := c.Request.Header.Get("X-token")
		Logger.Println(fmt.Sprintf("X-token[%s]", token))
		if token == "" {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeTokenError,
				"msg":  Enum.ErrMsg[Enum.CodeTokenError],
			})
			return
		}
		users, _ := Models.RedisDb.HGetAll("token:" + token).Result()
		if len(users) == 0 {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeTokenError,
				"msg":  Enum.ErrMsg[Enum.CodeTokenError],
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
