package middleware

import (
	"bytes"
	"fmt"
	"ginApi/common/enum"
	"ginApi/common/myLogger"
	_ "ginApi/common/myLogger"
	"ginApi/common/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

type LoggerMiddleware struct{}

func (this LoggerMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		myLogger.Logid = strings.Replace(uuid.New().String(), "-", "", -1)
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				panic(&response.Response{
					Code: enum.CodeParamError,
					Msg:  "获取body错误",
				})
				return
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
		myLogger.Println(accessStr)
		blw := &bodyLogWrite{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		response := blw.body.String()
		response = fmt.Sprintf("url[%s] response[%s]",
			c.Request.URL,
			response,
		)
		myLogger.Println(response)
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
