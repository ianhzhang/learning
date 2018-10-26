package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

func main() {
	fmt.Println("Hello, 世界")
	//str := `{ "ln": "Zhang", "fn":"Ian" }`
	str := `{ "ln": "Zhang", "fn":"Ian","son":{"ln":"Zhang","fn":"Kevin"}, "email":["izhang@extreme.com", "izhang@gmail.com"] }`

	var result map[string]interface{}
	
	// from byte array to map
	_ = json.Unmarshal( []byte(str), &result);
	fmt.Println("result:", result)
	fmt.Println(reflect.TypeOf(result))
	fmt.Println("ln:", result["ln"])
	fmt.Println("fn:", result["fn"])
	
	son := result["son"].(map[string]interface{})
	fmt.Println(son)
	fmt.Println(reflect.TypeOf(son))
	_ = son["fn"]
	fmt.Println("son fn:", son["fn"])
	fmt.Println("\n\n")
	son["fn"]="Lucas"



	emails := result["email"]
	fmt.Println("emails:", emails)
	fmt.Println("emails type:", reflect.TypeOf(emails))
	
	

	email0 := result["email"].([]interface{})

	fmt.Println(email0)
	fmt.Println(reflect.TypeOf(email0))
	fmt.Println(email0[0])

	// change the list
	result["email"] = append(email0, "abc@def.com")
	

	result["new"]="new"
	fmt.Println(result)

	fmt.Println("------------- End --------------")
}