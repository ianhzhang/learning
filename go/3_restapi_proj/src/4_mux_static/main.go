package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
type Person struct {
	Fn string   `json:"fn,omitempty"`
	Ln  string   `json:"ln,omitempty"`
}

func main() {
	fmt.Println("Start")
	router := mux.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
/*
	router.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path[1:])
			fmt.Println("-----3")
			
			http.ServeFile(w, r, "static/"+r.URL.Path[1:])
			//http.ServeFile(w, r, "dist/"+r.URL.Path[1:])
			//fmt.Println(http.Dir("./static/"))
			//router.PathPrefix("/").Handler(http.FileServer(http.Dir("/static")))
		},
	)
	*/
	router.HandleFunc(
		"/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello World 2 mux")
		},
	).Methods("GET")

	router.HandleFunc(
		"/getJsonBytes", func(w http.ResponseWriter, r *http.Request) {
			str := `{ "ln": "Zhang", "fn":"Ian","email":["izhang@extreme.com", "izhang@gmail.com"] }`
			fmt.Fprintf(w, str)
		},
	).Methods("GET")

	router.HandleFunc(
		"/getJsonType", func(w http.ResponseWriter, r *http.Request) {
			person := Person{Fn:"Ian", Ln:"Zhang"}
			//str := `{ "ln": "Zhang", "fn":"Ian","email":["izhang@extreme.com", "izhang@gmail.com"] }`
			//fmt.Fprintf(w, str)
			json.NewEncoder(w).Encode(person)
		},
	).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting http server ::", err)
	}
}