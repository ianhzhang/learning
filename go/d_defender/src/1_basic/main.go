package main

import (
	"fmt"
	"os"
	"gopkg.in/resty.v1"
	"encoding/json"
	"crypto/tls"
)

func main() {
	fmt.Println("Start")
	userId := os.Getenv("ECA_USER")
	passwd := os.Getenv("ECA_PASS")
	
	var root_url string = "https://10.68.67.222:5825/"

	data := map[string]interface{}{}
    data["grantType"] = "password"
    data["userId"] = userId
    data["password"] = passwd
	data["scope"] = "myScope"

	dataByte,_ := json.Marshal(data)
	var auth_url string = root_url + "management/v1/oauth2/token"
	fmt.Println(auth_url)
	resty.SetDebug(true)
	resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(dataByte).
		Post(auth_url)
	fmt.Println("-----------------")
	fmt.Println(err)
	fmt.Println("==================")
	fmt.Println(resp)
	fmt.Println("*********************")
}