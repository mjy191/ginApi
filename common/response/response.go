package response

import (
	"ginApi/common/enum"
	"ginApi/common/myLogger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Response 返回api接口的结构体
// code 返回状态码
// msg 返回错误信息
// data 数据
// logid 日志id
// timestamp 时间戳

type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Logid     string      `json:"logid"`
	Timestamp int64       `json:"timestamp"`
}

// New 初始化默认值
func New(res *Response) *Response {
	if res.Code == 0 {
		res.Code = enum.CodeSuccess
	}
	if res.Msg == "" {
		res.Msg = enum.ErrMsg[enum.CodeSuccess]
	}
	if res.Logid == "" {
		res.Logid = myLogger.LogId
	}
	res.Timestamp = time.Now().Unix()
	return res
}

// Success 成功的返回
func Success(c *gin.Context, res *Response) {
	res = New(res)
	c.JSON(http.StatusOK, &res)
}

// Fail 失败的返回
func Fail(c *gin.Context, res *Response) {
	res = New(res)
	c.AbortWithStatusJSON(http.StatusOK, &res)
}
