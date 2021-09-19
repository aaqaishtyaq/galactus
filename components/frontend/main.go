package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	port        = ":7000"
	staticDir   = "./static/"
	staticRoute = "/"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WorkspaceHandler(rw http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Panic("Websocket connection failed")
		return
	}

	for {
		fmt.Println("Awaiting message....")
		_, payload, err := wsConn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("Message received: \"%s\"\n", payload)
		time.Sleep(1 * time.Second)
		echo := fmt.Sprintf("Starting workspace with Git Remote address: %s", payload)
		err = wsConn.WriteMessage(websocket.TextMessage, []byte(echo))
		if err != nil {
			break
		}

		fmt.Printf("Sending message \"%s\" to \"%s\"\n", echo, wsConn.RemoteAddr())
		fmt.Println("-------")
	}
}

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logs\n"))
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Notes\n"))
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/logs", LogsHandler)
	router.HandleFunc("/notes", NotesHandler)
	router.HandleFunc("/workspace", WorkspaceHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	handleRequests()
}
