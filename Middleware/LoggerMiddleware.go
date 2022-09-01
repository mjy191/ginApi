package Middleware

import (
	"bytes"
	"fmt"
	"ginApi/Common/Enum"
	"ginApi/Common/Logger"
	_ "ginApi/Common/Logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

type LoggerMiddleware struct{}

func (this LoggerMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		Logger.Logid = strings.Replace(uuid.New().String(), "-", "", -1)
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
					"code": Enum.CodeParamError,
					"msg":  "获取body错误",
				})
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		}
		accessStr := fmt.Sprintf("url[%s] ip[%s] method[%s] post_data[%s] body[%s]",
			c.Request.URL,
			c.ClientIP(),
			c.Request.Method,
			c.Request.PostForm.Encode(),
			string(body),
		)
		Logger.Println(accessStr)
		blw := &bodyLogWrite{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		response := blw.body.String()
		response = fmt.Sprintf("url[%s] response[%s]",
			c.Request.URL,
			response,
		)
		Logger.Println(response)
	}
}

type bodyLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWrite) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
