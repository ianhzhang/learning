

// Package helloworld provides a set of Cloud Functions samples.
package gcptest

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"gcptest/myfirestore"
	"bytes"
)


// HelloGet is an HTTP Cloud Function.
func IanFirestoreTest(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	myfirestore.Test_firestore()
	fmt.Fprint(w, "Hello, World final2: " + t.Format(time.RFC3339))

}

func dumpObj(obj interface{}) {
	buf := bytes.Buffer{}
	_ = json.NewEncoder(&buf).Encode(obj)
	fmt.Println(buf.String())
}

// [END functions_helloworld_get]

