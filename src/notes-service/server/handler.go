package server

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Notes Service"))
}

func NotesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Notes Service"))
}