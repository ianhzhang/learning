package main

// https://tutorialedge.net/golang/parsing-json-with-golang/

import (
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
	"reflect"
)


/////////////////////////////////////////////////////////////////////////////////
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

/////////////////////////////////////////////////////////////////////////////////
func main() {
	fmt.Println("Read Json file start 2")

	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var myUsers = new(Users)

	json.Unmarshal([]byte(byteValue), myUsers)
	fmt.Println(myUsers)

	for index, user := range(myUsers.Users) {
		fmt.Println(index)
		fmt.Println(user)
		fmt.Println(reflect.TypeOf(user))
	}
	/* wrong cannot range over interface
	for index, item := range(result["users"]) {
		fmt.Println(index)
		fmt.Println(item)
	}*/


}