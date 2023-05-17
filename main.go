package main

import (
	"ginApi/common/config"
	"ginApi/common/enum"
	"ginApi/common/response"
	"ginApi/middleware"
	"ginApi/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 日志中间件
	r.Use(middleware.LoggerMiddleware{}.Handle())
	// 捕获异常转换成正常的http请求
	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		// 判断是不是主动抛出的错误类型
		res, ok := err.(response.Response)
		// 未知的类型统一返回服务器开小差
		if !ok {
			response.Fail(c, &response.Response{
				Code: enum.CodeSystemError,
				Msg:  "服务器开小差",
			})
			return
		}
		// 返回主动抛出的错误
		response.Fail(c, &res)
	}))
	// 环境配置
	if config.Viper.Get("env") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r.Static("/static", "./static")
	//二级目录写法
	r.LoadHTMLGlob("template/**/*")
	routers.Router{}.Router(r)
	r.Run(":8000")
}
