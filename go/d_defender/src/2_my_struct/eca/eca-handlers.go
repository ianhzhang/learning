package eca

import (
	"net/http"
	"fmt"
	"reflect"
	"github.com/gorilla/mux"
	"strconv"
)

func GetHello1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ECA GetHello1")
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(w, "Hello World 1")
}

func GetHello2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ECA GetHello2")
	var firstName ="Ian"
	var lastName = "Zhang"

	// multiple way to make json string
	str := `{ "ln": "Zhang", "fn":"Ian","son":{"ln":"Zhang","fn":"Kevin"}, "email":["izhang@extreme.com", "izhang@gmail.com"] }`
	fmt.Println(reflect.TypeOf(str))  // string
	str = `{ "ln": "` + lastName + `", "fn": "` + firstName + `"}`
	str = fmt.Sprintf(`{ "ln": "%v", "fn":"%v"}`, lastName, firstName)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, str)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	name := vars["name"]
	seq := vars["seq"]
	fmt.Println("ECA GetData", name)
	fmt.Println(reflect.TypeOf(seq))
	fmt.Println(seq)
	i, err := strconv.Atoi(seq)
	fmt.Println(err)
	if err == nil {
		fmt.Println(i)
	}
	fmt.Fprintf(w, "Hello GetData " + name)
}
