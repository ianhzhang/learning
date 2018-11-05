package main

// https://tutorialedge.net/golang/parsing-json-with-golang/

import (
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
	//"reflect"
)


/////////////////////////////////////////////////////////////////////////////////

type DataRcd struct {
	Data DataItem `json:"data"`
	Seq int `json:"seq"`
}

type DataItem struct {
	A   float32 `json:"a"`
	B   float32 `json:"b"`
}


/////////////////////////////////////////////////////////////////////////////////
func main() {
	fmt.Println("Read Json file file1")

	jsonFile, err := os.Open("data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Read all file contentes
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(byteValue))

	// Convert bytes array to Object
	myData := []DataRcd{}
	json.Unmarshal([]byte(byteValue), &myData)
	fmt.Println(myData)

}