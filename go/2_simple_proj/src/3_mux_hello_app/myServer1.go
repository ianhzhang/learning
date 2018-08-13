package main

import (
	"fmt"
	"net/http"
)

func handler1(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "1.0")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler1)
	mux.HandleFunc("/api", handler2)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()

}
