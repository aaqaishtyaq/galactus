package main

import (
	"log"
	"net/http"
)

var port string = "8001"

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+port, router))
}
