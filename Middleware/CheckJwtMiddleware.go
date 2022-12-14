package Middleware

import (
	"fmt"
	"ginApi/Common/Enum"
	"ginApi/Common/Logger"
	"ginApi/Common/Tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CheckJwtMiddleware struct{}

func (this CheckJwtMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Tools.Config.GetString("token.type") != "jwt" {
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
		userId, err := Tools.Jwt{}.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeTokenError,
				"msg":  Enum.ErrMsg[Enum.CodeTokenError],
			})
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
