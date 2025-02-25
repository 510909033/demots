package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := "ws://localhost:8080/ws"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var wg sync.WaitGroup

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			wg.Add(10)
			for i := 0; i < 10; i++ {
				// go func() {
				wg.Done()

				err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
				if err != nil {
					log.Println("write:", err)
					return
				}
				// }()
			}

			wg.Wait()

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
