package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

// websocket测试
func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8000/websocket", nil)
	if err != nil {
		log.Println(err)
		return
	}
	go send(conn)
	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		fmt.Println(m, string(p))
	}
	defer conn.Close()
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}
