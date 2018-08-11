package main
// https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
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

	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"}, "animals":"none"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)
	
	// The object stored in the "birds" key is also stored as 
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	/*
	_ = json.Unmarshal( []byte(str), &result);
	fmt.Println("ihz1:", result)
	fmt.Println("ihz1:", result["ln"])
	*/

	birds := result["birds"].(map[string]interface{})
	
	fmt.Println("------------------- loop -----------")
	for key, value := range birds {
	  // Each value is an interface{} type, that is type asserted as a string
	  fmt.Println(key, ":", value.(string))
	}
	fmt.Println("--------------- Access directly")
	fmt.Println("eagle:", birds["eagle"])

	fmt.Println("end")
}