package main

import (
	"fmt"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Hello")
	str := `{ "ln": "Zhang", "fn":"Ian","son":{"ln":"Zhang","fn":"Kevin"}, "email":["izhang@extreme.com", "izhang@gmail.com"] }`
	fmt.Println(str)
	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	fmt.Fprintf(w, str)
}


func main() {
	fmt.Println("Start")

	http.HandleFunc("/api/hello", handleHello)

	err := http.ListenAndServe("0.0.0.0:5000", nil)
	if err != nil {
		log.Fatal("err starting http server: ", err)
		return
	}
}
