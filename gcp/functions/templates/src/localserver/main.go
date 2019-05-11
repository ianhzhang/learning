package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"gcp_functions"
)


func main() {
	fmt.Println("main")

	router := mux.NewRouter()
	router.HandleFunc("/hello1", gcp_functions.Hello1Get).Methods("GET")
	router.HandleFunc("/hello2", gcp_functions.Hello2Get).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":5050", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}