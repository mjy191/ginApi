package Routers

import (
	"ginApi/Controller/Ws"
	"github.com/gin-gonic/gin"
)

type Websocket struct{}

func (this Websocket) Router(r *gin.Engine) {
	// websocket接口
	r.GET("/websocket", Ws.WebsocketController{}.Handel)
	// 通过http推送消息给websocket
	r.POST("/websocket/send", Ws.WebsocketController{}.SendMsg)
}
