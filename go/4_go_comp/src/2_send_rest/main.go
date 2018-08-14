package main
// https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/

// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go

// https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c

// http://polyglot.ninja/golang-making-http-requests/


import (
    "bytes"
    "encoding/json"
    "fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Person struct {
	Fn   string
	Ln   string
}

func getAll(url string) {
	var people []Person
	fmt.Println(url)
	resp, err := http.Get(url)

	if err!= nil {
		fmt.Print("HTTP request failed with error %s\n", err)
	}  else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(reflect.TypeOf(data))	// uint 8

		// convert to string
		data1 := string(data)
		fmt.Println(reflect.TypeOf(data1))
		fmt.Println(data1)

		// convert to object
		_ = json.Unmarshal(data, &people)
		fmt.Println(people)
		fmt.Println(people[0])

	}
}

func getOne(url string) {
	person := Person{}
	fmt.Println(url)
	resp, err := http.Get(url)

	if err!= nil {
		fmt.Print("HTTP request failed with error %s\n", err)
	}  else {

		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(reflect.TypeOf(data))	// uint 8

		// Convert to string
		data1 := string(data)
		fmt.Println(reflect.TypeOf(data1))
		fmt.Println(data1)

		// Convert to Object
		_ = json.Unmarshal(data, &person)
		fmt.Println("--------------------------------")
		fmt.Println(person)
		fmt.Println(person.Fn)
	}
}

func addPerson(url string) {
	var person Person

	fmt.Println(url)
	// Play with map (json) data
	jsonData := map[string]string{"ln": "Zhang", "fn": "Cat"}
	fmt.Println(reflect.TypeOf(jsonData))						// map[string]string

	fmt.Println(jsonData)										// map[ln:Zhang fn:Cat]
	fmt.Println(jsonData["ln"])									// Zhang
	jsonData["ln"]="Sun"
	fmt.Println(jsonData)										// map[ln:Zhang fn:Cat]

	// convert
	jsonValue, _ := json.Marshal(jsonData)
	fmt.Println(reflect.TypeOf(jsonValue))
	fmt.Println(string(jsonValue))								// {"fn":"Cat","ln":"Sun"}

	// POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
		fmt.Printf("\nParse the body-------==========================\n\n")
	
		_ = json.NewDecoder(resp.Body).Decode(&person)
		fmt.Println(person)

		//fmt.Println("12222----------------")
		//data, _ := ioutil.ReadAll(resp.Body)
		
		//fmt.Println(reflect.TypeOf(data))			// []uint8
		
		//fmt.Println(string(data))
		/*
		fmt.Println("00000---------------------")
		var result map[string]interface{}
		json.Unmarshal([]byte(data), &result)
		jsonData := result["data"]
		fmt.Println(reflect.TypeOf(jsonData))			// string
		fmt.Println(jsonData)
		*/
		fmt.Println("111111---------------------")
	}
	defer resp.Body.Close()
}


func main() {
	fmt.Println("------------------- get all -------------------")
	getAll("http://localhost:5000/family/getAll")

	fmt.Println("------------------- get one -------------------")
	//getOne("http://localhost:5000/family/getPerson/1")

	//fmt.Println("------------------- add person -------------------")
	//addPerson("http://localhost:5000/family/addPerson")
}