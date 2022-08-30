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

func (this CheckSignMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramSign := c.Query("sign")
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeSignError,
				"msg":  "签名错误",
			})
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		body := string(data[0:len(data)])
		key := "abc123"
		sign := Tools.Sha1(key + body + key)
		signPreStr := fmt.Sprintf("url[%s] logid[%s] signPre[%s]\n",
			c.Request.URL,
			Logger.Logid,
			key+body+key,
		)
		Logger.Println(signPreStr)
		if sign != paramSign {
			c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
				"code": Enum.CodeSignError,
				"msg":  "签名错误",
			})
		} else {
			c.Next()
		}
	}
}
