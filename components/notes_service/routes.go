package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	APIVersion  string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.APIVersion + route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"IndexHandler",
		"/api/v1",
		"/",
		"GET",
		IndexHandler,
	},
	Route{
		"LogsHandler",
		"/api/v1",
		"/logs",
		"GET",
		LogsHandler,
	},
	Route{
		"NotesHandler",
		"/api/v1",
		"/notes",
		"GET",
		NotesHandler,
	},
}
