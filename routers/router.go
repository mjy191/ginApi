package routers

import "github.com/gin-gonic/gin"

type Router struct{}

// Router 路由
func (this Router) Router(r *gin.Engine) {
	ApiRouter{}.Router(r)
	AdminRouter{}.Router(r)
	Websocket{}.Router(r)
}
