package gcp_functions

import (
	"fmt"
	"net/http"
	"gcp_functions/utils"
)

func Hello1Get(w http.ResponseWriter, r *http.Request) {
	utils.Hello()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, `{"result": 1}`)
}