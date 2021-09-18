package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FrontendServer struct {

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home\n"))
}

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logs\n"))
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Notes\n"))
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/logs", LogsHandler)
	router.HandleFunc("/notes", NotesHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}
