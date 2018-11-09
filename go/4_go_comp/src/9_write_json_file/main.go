package main

import (
	"fmt"
	"encoding/json"
	//"os"
	"io/ioutil"
)

type DataItem struct {
	A   int `json:"a"`
	B   int `json:"b"`
	Nam string `json:"name"`
}

func main() {
	fmt.Println("start")

	dataList := []DataItem{{1,2, "Ian"}, {3,4, "Kevin"}}
	fmt.Println(dataList)

	fileName := "output/data/data2.json"
	dataListJsonByte, _ := json.Marshal(dataList)
	ioutil.WriteFile(fileName, dataListJsonByte, 0644)

	// Read back
	contentBytes, _ := ioutil.ReadFile(fileName)
	dataList1 := []DataItem{}
	json.Unmarshal(contentBytes, &dataList1)
	fmt.Println("----- Read back -----")
	fmt.Println(dataList1)

}