package main

import (
	"log"
	"net/http"
	"os"
	"secretsanta/drafter"
	"secretsanta/static"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/icza/gox/osx"
)

var FileServer = http.FileServerFS(static.StaticFS)

var port string = ":8080"

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var wsClient *int32 = new(int32)

func CloseOnNoConnect(conn *websocket.Conn) {
	atomic.AddInt32(wsClient, 1)
	log.Println("New client, total", *wsClient)
	f := conn.CloseHandler()
	conn.SetCloseHandler(func(code int, text string) error {
		atomic.AddInt32(wsClient, -1)
		if *wsClient == 0 {
			log.Println("No client, exit.")
			os.Exit(0)
		}
		return f(code, text)
	})
}

func wsHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error opening Websocket.")
		w.WriteHeader(http.StatusInternalServerError)
	}
	CloseOnNoConnect(conn)
	switch websocket.Subprotocols(r)[0] {
	case "draft":
		go drafter.DrafterWSApi(conn)
	}
}

func main() {
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	var err error
	http.Handle("/", FileServer)
	http.HandleFunc("/ws", wsHandle)
	err = osx.OpenDefault("http://localhost" + port)
	if err != nil {
		log.Println("Error opening browser:", err)
	}
	log.Println("Listening on port", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Println(err)
	}
}
