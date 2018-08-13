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
	str := `[{ "ln": "Zhang", "fn":"Ian" },{ "ln": "Sun", "fn":"Hong" }]`

	var names []map[string]interface{}
	
	_ = json.Unmarshal( []byte(str), &names);
	fmt.Println("ihz1:", names)
	fmt.Println("ihz2:", names[0])
	fmt.Println("ihz3:", names[1]["ln"])

	for i, name := range names {
		fmt.Println(i)
		fmt.Println(name["ln"], ":", name["fn"])
	}
	fmt.Println("end")
}