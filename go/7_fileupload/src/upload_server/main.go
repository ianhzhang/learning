package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"fmt"
	"os"
	"io"
	"github.com/gorilla/mux"
)

// https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.5.html

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------1")
	fmt.Println("method:", r.Method)		// POST
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	fmt.Println("-----------------2")
	fmt.Println(file)	//{0xc420095440}   file used at line #36.   io.Copy(f, file)
	fmt.Println(handler.Filename)		// source file name
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	fmt.Println("-----------------3")	// create a file
	f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)		// copy the file to newly created file
	
}

// our main function
func main() {
	router := mux.NewRouter()
	//router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/upload", upload).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", router))
}
