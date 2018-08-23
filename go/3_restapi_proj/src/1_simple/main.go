package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle graphql")
		fmt.Fprintf(w, "Hello World")
	})

	http.ListenAndServe(":5000", nil)
}