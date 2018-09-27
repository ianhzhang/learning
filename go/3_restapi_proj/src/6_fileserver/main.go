package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)



func handleHello(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Rslt string
	}

	fmt.Println("Handle Hello 6")
	rslt := Response{Rslt:"succ3"}
	fmt.Println(rslt)
	json.NewEncoder(w).Encode(rslt)
}


func main() {
	fmt.Println("Start 6")
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", handleHello).Methods("GET")

	// has to be at the end to catch all
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	

	err := http.ListenAndServe("localhost:8081", r)
	if err != nil {
		log.Fatal("err starting http server: ", err)
		return
	}
}