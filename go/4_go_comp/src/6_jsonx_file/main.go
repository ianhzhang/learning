package main

// https://tutorialedge.net/golang/parsing-json-with-golang/

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
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

	jsonxFile, err := os.Open("data.jsonx")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonxFile.Close()

	reader := bufio.NewReader(jsonxFile)
	for {
		line, _, err := reader.ReadLine()

		if err!= nil {
			break
		} else {
			fmt.Println(string(line))
			dataItem := DataItem{}
			json.Unmarshal([]byte(line), &dataItem)
			fmt.Println(dataItem)
			fmt.Println("------------------------------")
		}
	}

}