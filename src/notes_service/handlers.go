package main

import (
	"fmt"
	"net/http"

	"github.com/aaqaishtyaq/dvito/src/notes_service/markdown"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logs!")
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := markdown.ListNotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, notes)
}
