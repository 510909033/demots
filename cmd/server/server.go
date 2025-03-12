package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Connect struct {
	connMap *sync.Map
}

var connects = Connect{
	connMap: &sync.Map{},
}

// 增加一个连接
func (c *Connect) AddConn(conn *websocket.Conn) {
	c.connMap.Store(conn, time.Now())
	log.Printf("New connection added: %v", conn.RemoteAddr())
}

// 判断一个链接是否存在
func (c *Connect) IsConnExist(conn *websocket.Conn) bool {
	_, ok := c.connMap.Load(conn)
	return ok
}

// 删除一个连接
func (c *Connect) DelConn(conn *websocket.Conn) {
	startTime, ok := c.connMap.Load(conn)
	if ok {
		duration := time.Since(startTime.(time.Time))
		log.Printf("Connection closed: %v, Duration: %v", conn.RemoteAddr(), duration)
	}
	c.connMap.Delete(conn)
}

type Response struct {
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer func() {
		connects.DelConn(conn)
		conn.Close()
	}()

	connects.AddConn(conn)

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		time.Sleep(time.Second * 2)

		var resp = Response{
			Msg:  "server response: " + string(message),
			Time: time.Now().Format(time.DateTime),
		}
		vals, _ := json.Marshal(resp)
		err = conn.WriteMessage(mt, vals)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	http.HandleFunc("/ws", echo)
	fmt.Println("WebSocket server started at ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
