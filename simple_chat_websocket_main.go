package main

import (
	"log"
	"net/http"
)

func simpleChatWebsocketRoute() {
	go h.run()
	http.Handle("/", http.FileServer(http.Dir("./chat_websocket_client")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
