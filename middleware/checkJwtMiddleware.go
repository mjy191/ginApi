package middleware

import (
	"fmt"
	"ginApi/common/config"
	"ginApi/common/enum"
	"ginApi/common/jwt"
	"ginApi/common/myLogger"
	"ginApi/common/response"
	"github.com/gin-gonic/gin"
)

type CheckJwtMiddleware struct{}

func (this CheckJwtMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Viper.GetString("token.type") != "jwt" {
			c.Next()
			return
		}
		token := c.Request.Header.Get("X-token")
		myLogger.Println(fmt.Sprintf("X-token[%s]", token))
		if token == "" {
			panic(response.Response{
				Code: enum.CodeTokenError,
				Msg:  enum.ErrMsg[enum.CodeTokenError],
			})
		}
		userId, err := jwt.Jwt{}.ValidateToken(token)
		if err != nil {
			panic(response.Response{
				Code: enum.CodeTokenError,
				Msg:  enum.ErrMsg[enum.CodeTokenError],
			})
		}

		c.Set("userId", userId)
		c.Next()
	}
}
