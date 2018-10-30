package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


func NewRouter() *mux.Router {
	fmt.Println("NewRouter")
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range GetRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		//ihz handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path("/defender/api" + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}