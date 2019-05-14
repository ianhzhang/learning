package main

import (
	"fmt"
	"log"
	"net/http"

	"gcp_functions"
)


func main() {
	fmt.Println("main")

	http.HandleFunc("/", gcp_functions.Entry)
	log.Fatal(http.ListenAndServe(":5050", nil))
}