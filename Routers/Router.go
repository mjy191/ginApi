package Routers

import "github.com/gin-gonic/gin"

type Router struct{}

func (this Router) Router(r *gin.Engine) {
	ApiRouter{}.Router(r)
	AdminRouter{}.Router(r)
	Websocket{}.Router(r)
}
