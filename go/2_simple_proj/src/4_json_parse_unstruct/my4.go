package main

import (
	"fmt"
	"encoding/json"
)


func main() {
	fmt.Println("Hello, 世界")
	received_JSON_str := `{ "ln": "Zhang", "fn":"Ian" }`

	var arbitrary_json map[string]interface{}
	
	_ = json.Unmarshal( []byte(received_JSON_str), &arbitrary_json);
    fmt.Println("ihz1: ", arbitrary_json)
    

    fmt.Println("--------------------------------------------")
    received_JSON_str1 := `[{ "ln": "Zhang", "fn":"Ian1" },{ "ln": "Sun", "fn":"Hong" }]`
    var arbitrary_json1 []map[string]interface{}
    _ = json.Unmarshal( []byte(received_JSON_str1), &arbitrary_json1);
    fmt.Println("ihz2: ", arbitrary_json1)



	fmt.Println("end")
}