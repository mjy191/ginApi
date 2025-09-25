package middleware

import (
	"bytes"
	"fmt"
	"ginApi/common/config"
	"ginApi/common/enum"
	"ginApi/common/myLogger"
	"ginApi/common/response"
	"ginApi/common/tools"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type CheckSignMiddleware struct {
}

// Handle 验证签名中间件
func (this CheckSignMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramSign := c.Query("sign")
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(&response.Response{
				Code: enum.CodeSignError,
				Msg:  enum.ErrMsg[enum.CodeSignError],
			})
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		body := string(data[0:len(data)])
		key := config.Viper.GetString("signKey")
		sign := tools.Sha1(key + body + key)
		signPreStr := fmt.Sprintf("url[%s] signPre[%s]",
			c.Request.URL,
			key+body+key,
		)
		myLogger.Println(signPreStr)
		if sign != paramSign {
			panic(&response.Response{
				Code: enum.CodeSignError,
				Msg:  enum.ErrMsg[enum.CodeSignError],
			})
		} else {
			c.Next()
		}
	}
}
