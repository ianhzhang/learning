package main

// https://stackoverflow.com/questions/49006594/interface-to-byte-conversion-in-golang?rq=1
import (
    "encoding/json"
    "fmt"
)

func main() {
    // ********************* Marshal *********************
    u := map[string]interface{}{}
    u["name"] = "kish"
    u["age"] = 28
    u["work"] = "engine"
    //u["hobbies"] = []string{"art", "football"}
    u["hobbies"] = "art"

    fmt.Println("--------------- map -----------------")
    fmt.Println(u)
    b, err := json.Marshal(u)
    if err != nil {
        panic(err)
    }
    fmt.Println("\n\n-------------- from map to byte --------------")
    fmt.Println(string(b))

    fmt.Println("\n\n--------------from byte to map again ---------------------- ")
    // ********************* Unmarshal *********************
    var a interface{}
    err = json.Unmarshal(b, &a)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(a)

    //============================================================
    fmt.Println("\n\n")

    str := `{ "ln": "Zhang", "fn":"Ian","email":["izhang@extreme.com", "izhang@gmail.com"] }`

	result := map[string]interface{}{}
	
    _ = json.Unmarshal( []byte(str), &result);
    result["new"] = "newname"
    fmt.Println(result)
    b1, err := json.Marshal(result) 
	
	if err != nil {
        fmt.Println(err)
		
	} else {
		fmt.Println(string(b1))
    }

}