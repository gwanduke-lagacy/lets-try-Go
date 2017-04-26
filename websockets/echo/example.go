package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func runEchoWebsocketExample() {
	// -> https://jeremywho.com/simple-websocket-client-in-go-using-gorilla-websocket/
	// http://www.websocket.org/echo.html
	URL := "ws://echo.websocket.org"

	var dialer *websocket.Dialer

	// 메시지를 보내고 받을 수 있는 *websocket.Conn 을 리턴한다
	conn, _, err := dialer.Dial(URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go timeWriter(conn)

	// https://godoc.org/github.com/gorilla/websocket#Conn.ReadMessage
	// https://godoc.org/github.com/gorilla/websocket#Conn.ReadJSON
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}
		fmt.Printf("received: %s\n", message)
	}
	// 에러핸들링을 하지 않는 경우
	// for {
	// 	_, message, _ := conn.ReadMessage()
	// 	fmt.Printf("received: %s\n", message)
	// }
}

func timeWriter(conn *websocket.Conn) {
	for {
		// 2초간 쉬면서 conn으로 데이터를 전송한다.
		time.Sleep(time.Second * 2)

		// messsageType, 보낼메시지
		// https://godoc.org/github.com/gorilla/websocket#Conn.WriteMessage
		// https://godoc.org/github.com/gorilla/websocket#Conn.WriteJSON
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	}
}
