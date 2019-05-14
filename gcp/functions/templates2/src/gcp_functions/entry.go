package gcp_functions

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)
func defaultNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("REST:", r.URL.Path)
	http.NotFound(w, r)
}

func Entry(w http.ResponseWriter, r *http.Request) {
	log.Println("Entry")
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/hello1/{id}", Hello1Get).Methods("GET")
	router.HandleFunc("/hello2", Hello2Get).Methods("GET")
	router.HandleFunc("/", defaultNotFound)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"})
	handlers.CORS(headersOk, originsOk, methodsOk)(router).ServeHTTP(w, r)
}