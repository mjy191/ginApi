package ws

import (
	"encoding/json"
	"fmt"
	"ginApi/controller"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
)

type WebsocketController struct {
	controller.BaseController
}

var upgrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 用户的uuid
type userInfo struct {
	Uuid string `json:"uuid"`
}

// 服务结构
type websocketServer struct {
	userInfo
	connMap map[string]*websocket.Conn
	conns   []*websocket.Conn
	msg     chan string
}

var ws = &websocketServer{
	connMap: make(map[string]*websocket.Conn),
	msg:     make(chan string),
}

func (this WebsocketController) Handel(c *gin.Context) {
	// 判断是否是websocket
	if c.IsWebsocket() {

	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 保存链接
	ws.conns = append(ws.conns, conn)
	fmt.Println("打印链接", ws.conns)
	go listenBroadcast()
	//广播消息上线了
	ws.msg <- fmt.Sprintf("%v", conn.RemoteAddr()) + " 上线了"
	// 阻塞主进程
	for {
		_, p, e := conn.ReadMessage()
		if e != nil {
			break
		}

		// 绑定uuid和conn关系
		if strings.Contains(string(p), "uuid") {
			if er := json.Unmarshal(p, &ws.userInfo); er == nil {
				ws.connMap[ws.userInfo.Uuid] = conn
				log.Println("connMap", ws.connMap)
			}
		} else {
			// 发送广播消息
			ws.msg <- fmt.Sprintf("%v：%s", conn.RemoteAddr(), string(p))
		}
		log.Println(string(p))
	}
	// 关闭链接
	conn.WriteControl(websocket.CloseMessage, []byte("关闭链接"), time.Now().Add(time.Second))
	// 删除conn从conns切片中
	var index int
	for key, val := range ws.conns {
		if val == conn {
			index = key
			break
		}
	}
	ws.conns = append(ws.conns[:index], ws.conns[index+1:]...)
	// 删除connmap
	for key, val := range ws.connMap {
		if val == conn {
			delete(ws.connMap, key)
			break
		}
	}
	defer conn.Close()
	log.Println(ws.conns)
	log.Println("服务关闭")
}

// 监听广播消息，发送消息
func listenBroadcast() {
	for {
		msg := <-ws.msg
		for i := range ws.conns {
			ws.conns[i].WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}
}

// 单发消息
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

	//发送消息给指定用户
	msg := c.Request.FormValue("msg")
	uuid := c.Request.FormValue("uuid")
	log.Println(ws.connMap)
	conn, ok := ws.connMap[uuid]
	if ok {
		singleMsg(conn, msg)
	}
	c.Writer.Write([]byte(fmt.Sprintf("%v", "发送成功")))
}
