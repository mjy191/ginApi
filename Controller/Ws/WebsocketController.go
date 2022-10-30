package Ws

import (
	"encoding/json"
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
var connMap = make(map[string]*websocket.Conn)

type userInfo struct {
	Uuid string `json:"uuid""`
}

var user userInfo

var msg = make(chan string, 20)

func (this WebsocketController) Handel(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//报错uuid和conn
	//connMap[uuid] = conn
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
		if er := json.Unmarshal(p, &user); er == nil {
			connMap[user.Uuid] = conn
		}
		//保存对应关系
		log.Println("connMap", connMap)
		// 遍历链接，广播发送消息
		log.Println(string(p))
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
	// 删除connmap
	for key, val := range connMap {
		if val == conn {
			delete(connMap, key)
			break
		}
	}
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

//单发消息
func singleMsg(conn *websocket.Conn, msg string) {
	conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

// 通过http协议推送消息
func (this WebsocketController) SendMsg(c *gin.Context) {
	// 从body获取消息推送
	//body := request.Body
	//b, _ := io.ReadAll(body)
	//body.Close()
	//fmt.Println(string(b))

	// 从form表单获取消息推送
	// 群发消息
	//b := c.Request.FormValue("msg")
	//fmt.Println(b)
	//msg <- b
	//fmt.Println(conns)
	//go broadcast(conns)

	//发送单个消息
	mg := c.Request.FormValue("msg")
	uuid := c.Request.FormValue("uuid")
	log.Println(connMap)
	conn, ok := connMap[uuid]
	if ok {
		singleMsg(conn, mg)
	}
	c.Writer.Write([]byte(fmt.Sprintf("%v", "发送成功")))
}
