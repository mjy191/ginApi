package Middleware

import (
	"bytes"
	"fmt"
	"ginApi/Common/Enum"
	"ginApi/Common/Logger"
	"ginApi/Common/Tools"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type CheckSignMiddleware struct {
}

// Handle 验证签名中间件
func (this CheckSignMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramSign := c.Query("sign")
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeSignError,
				"msg":  Enum.ErrMsg[Enum.CodeSignError],
			})
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		body := string(data[0:len(data)])
		key := Tools.Config.GetString("signKey")
		sign := Tools.Sha1(key + body + key)
		signPreStr := fmt.Sprintf("url[%s] signPre[%s]",
			c.Request.URL,
			key+body+key,
		)
		Logger.Println(signPreStr)
		if sign != paramSign {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeSignError,
				"msg":  Enum.ErrMsg[Enum.CodeSignError],
			})
		} else {
			c.Next()
		}
	}
}
