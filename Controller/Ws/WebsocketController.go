package Ws

import (
	"fmt"
	"ginApi/Controller"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type WebsocketController struct {
	Controller.BaseController
}

var upgrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var conns []*websocket.Conn

var msg = make(chan string, 20)

func (this WebsocketController) Handel(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 保存所有的链接
	msg <- "有人上线：" + fmt.Sprintf("%v", conn.RemoteAddr())
	go broadcast(conns)
	conns = append(conns, conn)
	fmt.Println("打印链接", conns)
	for {
		_, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		// 遍历链接，广播发送消息
		msg <- fmt.Sprintf("%v：%s", conn.RemoteAddr(), string(p))
		go broadcast(conns)
	}
	conn.WriteControl(websocket.CloseMessage, []byte("关闭链接"), time.Now().Add(time.Second))
	// 删除conn从conns切片中
	var index int
	for key, val := range conns {
		if val == conn {
			index = key
			break
		}
	}
	conns = append(conns[:index], conns[index+1:]...)
	defer conn.Close()
	log.Println(conns)
	log.Println("服务关闭")
}

//广播消息
func broadcast(conns []*websocket.Conn) {
	m := <-msg
	for i := range conns {
		conns[i].WriteMessage(websocket.TextMessage, []byte(m))
	}
}

// 通过http协议推送消息
func (this WebsocketController) SendMsg(c *gin.Context) {
	// 从body获取消息推送
	//body := request.Body
	//b, _ := io.ReadAll(body)
	//body.Close()
	//fmt.Println(string(b))
	// 从form表单获取消息推送
	b := c.Request.FormValue("msg")
	fmt.Println(b)
	msg <- b
	fmt.Println(conns)
	go broadcast(conns)
	c.Writer.Write([]byte(fmt.Sprintf("%v", "发送成功")))
}
