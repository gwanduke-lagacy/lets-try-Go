package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func startLoopWebsocketRoute() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wshandler(w, r)
	})

	http.Handle("/", http.FileServer(http.Dir("loop_websocket_client")))

	http.ListenAndServe(":12312", nil)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// for 루프가 커넥션이 끝날때까지 실행되며 클아이언트가 보낼때마다 기록한다.
// 일단 에러가 나면 커넥션은 깨지고 루프가 끝난다.
func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
