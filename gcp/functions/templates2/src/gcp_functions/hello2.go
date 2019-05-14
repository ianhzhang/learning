package gcp_functions

import (
	"fmt"
	"net/http"
)

func Hello2Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, `{"result": 2}`)
}