// https://gist.github.com/kousik93/6d95c4c4d37d8c731d7b

package main

import (
	"fmt"
	"encoding/json"
)
/*
{
  "birds": {
    "pigeon":"likes to perch on rocks",
    "eagle":"bird of prey"
  },
  "animals": "none"
}
*/

func main() {
	fmt.Println("Hello, 世界")
	str := `{"ZhangFamily": [{ "ln": "Zhang", "fn":"Ian" },{ "ln": "Sun", "fn":"Hong" }]}`

	var f interface{}
	

	_ = json.Unmarshal( []byte(str), &f);
	

	
	m := f.(map[string]interface{})
	fmt.Println(m["ZhangFamily"])

	for k, v := range m {
		fmt.Println(v)
		fmt.Println("----")
    switch vv := v.(type) {
    case string:
        fmt.Println(k, "is string", vv)
    case int:
        fmt.Println(k, "is int", vv)
    case []interface{}:
				fmt.Println(k, "is an array:")
				fmt.Println(vv)
        for i, u := range vv {
            fmt.Println(i, u)
        }
    default:
        fmt.Println(k, "is of a type I don't know how to handle")
    }
}
	
	fmt.Println("------------------- Loop ------------------")
	
	for k, v := range m {
		fmt.Println(v)
		fmt.Println("----")
    switch vv := v.(type) {
    case string:
        fmt.Println(k, "is string", vv)
    case int:
        fmt.Println(k, "is int", vv)
    case []interface{}:
				fmt.Println(k, "is an array:")
				fmt.Println(vv)
        for i, u := range vv {
            fmt.Println(i, u)
        }
    default:
        fmt.Println(k, "is of a type I don't know how to handle")
    }
}


	

	fmt.Println("end")
}