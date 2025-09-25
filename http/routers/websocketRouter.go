package routers

import (
	"ginApi/http/controller/ws"
	"github.com/gin-gonic/gin"
)

type Websocket struct{}

func (this Websocket) Router(r *gin.Engine) {
	// websocket接口
	r.GET("/websocket", ws.WebsocketController{}.Handel)
	// 通过http推送消息给websocket
	r.POST("/websocket/send", ws.WebsocketController{}.SendMsg)
}
