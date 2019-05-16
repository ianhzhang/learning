package main

import (
        "fmt"
        "encoding/json"

        "bytes"
        "myfirestore"
)


// HelloGet is an HTTP Cloud Function.
func main() {
        fmt.Println("Start1:")
        myfirestore.Test_firestore();
}

func dumpObj(obj interface{}) {
        buf := bytes.Buffer{}
        _ = json.NewEncoder(&buf).Encode(obj)
        fmt.Println(buf.String())
}

