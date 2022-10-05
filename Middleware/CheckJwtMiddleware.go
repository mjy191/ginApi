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
		}
		token := c.Request.Header.Get("x-token")
		Logger.Println(fmt.Sprintf("x-token:%s", token))
		userId, err := Tools.Jwt{}.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeTokenError,
				"msg":  Enum.ErrMsg[Enum.CodeTokenError],
			})
			return
		}

		c.Set("users", map[string]int{
			"userId": userId,
		})
		c.Next()
	}
}
