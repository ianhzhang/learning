package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)
type Person struct {
	Fn string   `json:"fn,omitempty"`
	Ln  string   `json:"ln,omitempty"`
}

func main() {
	fmt.Println("Start")

	http.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path[1:])
			fmt.Println("-----")
			
			//http.ServeFile(w, r, "static/"+r.URL.Path[1:])
			http.ServeFile(w, r, "dist/"+r.URL.Path[1:])
		},
	)
	http.HandleFunc(
		"/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World")
		},
	)

	http.HandleFunc(
		"/getJsonBytes", func(w http.ResponseWriter, r *http.Request) {
			str := `{ "ln": "Zhang", "fn":"Ian","email":["izhang@extreme.com", "izhang@gmail.com"] }`
			fmt.Fprintf(w, str)
		},
	)
	http.HandleFunc(
		"/getJsonType", func(w http.ResponseWriter, r *http.Request) {
			person := Person{Fn:"Ian", Ln:"Zhang"}
			//str := `{ "ln": "Zhang", "fn":"Ian","email":["izhang@extreme.com", "izhang@gmail.com"] }`
			//fmt.Fprintf(w, str)
			json.NewEncoder(w).Encode(person)
		},
	)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting http server ::", err)
	}
}