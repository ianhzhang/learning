package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	fmt.Println("Hello, 世界")
	str := `{ "ln": "Zhang", "fn":"Ian" }`

	var result map[string]interface{}
	
	_ = json.Unmarshal( []byte(str), &result);
	fmt.Println("ihz1:", result)
	fmt.Println("ihz1:", result["ln"])
	fmt.Println("ihz1:", result["fn"])

	fmt.Println("end")
}