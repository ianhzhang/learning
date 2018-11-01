package restapi

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func StartRest() {
	fmt.Println("Start RestAPI1")

	router := NewRouter()
	
	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe("localhost:8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}