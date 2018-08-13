package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

type Baz struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	RawData []int  `json:"rawdata"`
	Epoch   string `json:"epoch"`
}

type Bar struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	RawData []*Qux `json:"rawdata"`
	Epoch   string `json:"epoch"`
}

//nested inside Bar
type Qux struct {
	Key    string `json:"key"`
	Values []int  `json:"values`
}

func main() {
	str := `{ "foo" : { "id":"234", "type":"bar", "rawdata":[{"key":"dog", "values":[123,234]}, {"key":"cat", "values":[23, 45]}], "epoch":"1433120656705"}}`
	str2 := `{ "foo": { "id":"124", "type":"baz", "rawdata":[123, 345,345],"epoch": "1433120656704"}}`

	var objmap map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(str), &objmap); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ihz:", objmap["foo"])
	typ, err := getTypeByUnmarshalToInterface(objmap["foo"])
	if err != nil {
		log.Fatal(err)
	}



	fmt.Println("type:", typ)

	if err := json.Unmarshal([]byte(str2), &objmap); err != nil {
		log.Fatal(err)
	}

	typ2, err := getTypeByRegexp(objmap["foo"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("type:", typ)

	switch typ2 {
	case "bar":
		var b Bar
		if err := json.Unmarshal(*objmap["foo"], &b); err != nil {
			log.Fatal(err)
		}
		fmt.Println(b)
	case "baz":
		var b Baz
		if err := json.Unmarshal(*objmap["foo"], &b); err != nil {
			log.Fatal(err)
		}
		fmt.Println(b)
	default:
		log.Fatal("unknown type")
	}
}

func getTypeByUnmarshalToInterface(foo *json.RawMessage) (string, error) {
	var tmp map[string]interface{}
	if err := json.Unmarshal(*foo, &tmp); err != nil {
		return "", err
	}
	fmt.Println("ihz1:", tmp["type"])
	if typ, ok := tmp["type"].(string); ok {
		return typ, nil
	} else {
		return "", fmt.Errorf("type should be a string")
	}
}

func getTypeByRegexp(foo *json.RawMessage) (string, error) {
	re := regexp.MustCompile(`"type"\s*:\s*"([^"]*)"`)
	matches := re.FindAllSubmatch(*foo, -1)
	if len(matches) > 0 {
		return string(matches[0][1]), nil
	} else {
		return "", fmt.Errorf("cannot find type")
	}
}

