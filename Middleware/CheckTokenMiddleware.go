package Middleware

import (
	"ginApi/Common/Enum"
	"ginApi/Common/Tools"
	"ginApi/Models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type CheckTokenMiddleware struct{}

func (this CheckTokenMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Tools.Config.GetString("token.type") != "token" {
			c.Next()
			return
		}
		param := make(map[string]interface{})
		if err := c.ShouldBindBodyWith(&param, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeParamError,
				"msg":  "参数解析错误",
			})
			return
		}
		if _, ok := param["token"]; !ok {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeTokenError,
				"msg":  Enum.ErrMsg[Enum.CodeTokenError],
			})
			return
		}
		users, _ := Models.RedisDb.HGetAll("token:" + param["token"].(string)).Result()
		if len(users) == 0 {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeTokenError,
				"msg":  Enum.ErrMsg[Enum.CodeTokenError],
			})
			return
		}
		c.Set("users", users)
		c.Next()
	}
}
