package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

func main() {
	fmt.Println("Hello")
	type Data struct {
		A int
		B int
	}
	type RedisItem struct {
		Seq   int
		Data  Data
	}

	str := `{"a": 298202, "c": 258720}`
	
	var result map[string]interface{}

	_ = json.Unmarshal([]byte(str), &result)
	//_ = json.Unmarshal( []byte(str), &result);
	//fmt.Println(result)
	
	mydata := Data{1,2}

	item := &RedisItem{
		Seq:   1234,
		Data:  mydata,
	}

	fmt.Println(reflect.TypeOf(item))
	fmt.Println(item)

	itemByte, _ := json.Marshal(item)
	fmt.Println("\n")
	fmt.Println(string(itemByte))

	// my own
	seq := 1234
	dataStr := fmt.Sprintf(`{"seq":%d, "data":%s}`, seq, str)
	fmt.Println(dataStr) 
	
}