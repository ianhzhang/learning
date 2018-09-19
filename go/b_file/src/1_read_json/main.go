package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"reflect"
	//"strings"
	"regexp"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}


func main() {
	jsonFile, _ := os.Open("users.json")
	defer jsonFile.Close()

	// read file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// print byte array
	fmt.Println(byteValue)
	// print original string
	fmt.Println(string(byteValue))
	/*
		[
			{
				"a":1,
				"b":2
			},
			{
				"a":3,
				"b":4
			}
		]
	*/

	// remove new line and white space
	re := regexp.MustCompile(`\t|\n|\s`)
	s := re.ReplaceAllString(string(byteValue),"")
	fmt.Println(s)


	// *************** struct ****************************
	fmt.Println("Try struct ---------------------")
	var data []Data
	json.Unmarshal(byteValue, &data)

	fmt.Println(data)				// [{1 2} {3 4}]
	fmt.Println(reflect.TypeOf(data))	// [] main.Data
	fmt.Println(data[0].A)				// 1

	byteArray,_ := json.Marshal(data)
	fmt.Println(byteArray)			// [91 123 34 97 34 58 49 44 34 98 34 58 50 125 44 123 34 97 34 58 51 44 34 98 34 58 52 125 93]
	fmt.Println(reflect.TypeOf(byteArray))		// []uint8
	jsonStr := string(byteArray)
	fmt.Println(jsonStr)	// [{"a":1,"b":2},{"a":3,"b":4}]
	fmt.Println(reflect.TypeOf(jsonStr))		// string


	// ******************* interface ****************************
	fmt.Println("Try map ---------------------")
	fmt.Println(byteValue)
	var result []map[string]interface{}
	_ = json.Unmarshal( byteValue, &result);	
	fmt.Println(reflect.TypeOf(result))		// []map[string]interface {}
	fmt.Println("result:", result)			// result: [map[a:1 b:2] map[a:3 b:4]]

	for _, elem := range result {
		fmt.Println(elem)		// map[a:1 b:2], map[a:3 b:4]
		fmt.Println(elem["a"])	// 1, 3

	}

}