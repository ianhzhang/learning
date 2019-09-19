package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // n << x is n times 2, x times.   
    // n * 2^x
    // 1 M is 2^20
    // 10 M is 10 * 2^20
    // 200M is 200 * 2^20  ( 200 << 20)   

    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()

    fmt.Printf("\n\nUploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    if _, err := os.Stat("./upload-files/"); os.IsNotExist(err) {
        os.Mkdir("./upload-files/", 0700)
    }
    ioutil.WriteFile("upload-files/" + handler.Filename, fileBytes, 0644)
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	println("setupRoutes")
    http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
	println("setupRoutes end -======")
}

func main() {
    fmt.Println("Hello World")

    setupRoutes()
}
