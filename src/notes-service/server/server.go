package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/notes/index", NotesIndexHandler).Methods(http.MethodGet)
	// router.HandleFunc("/notes/index/{note: [a-zA-Z0-9]+}", LogsHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8001", router))
}
