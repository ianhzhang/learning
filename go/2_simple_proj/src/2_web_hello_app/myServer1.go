package main

import (
	"fmt"
	"net/http"
)

func handler1(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World. go server")
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "1.0")
}

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/api", handler2)
	http.ListenAndServe(":8080", nil)
}
