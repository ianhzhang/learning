package main

import (
	"fmt"
	"log"
	"reflect"
	"os"
	"github.com/fatih/structs"
	"test1/container_api"
	"test1/firestore_api"
)



type A struct {
	A int
	B int
}

type B struct {
	A int 
	B int 
	C string
}

func main() {
	fmt.Println("Go Generic")

	list := container_api.List{}

	a1 := A{1,2}
	a2 := A{3,4}

	list.Append(a1)
	list = append(list, a2)
	fmt.Println(list)

	b1 := B{5,6,"Hello 1"}
	b2 := B{7,8,"Hello 3"}
	list.Append(b1)
	list = append(list, b2)
	fmt.Println(list)
	fmt.Println(list.GetByIndex(2))

	fmt.Println("-----------------------------------------------")
	firestoreConn := firestore_api.FirestoreConn{}
	fmt.Println(firestoreConn)
	projectID := os.Getenv("GCLOUD_PROJECT")
	fmt.Println("projectID:", projectID)

	client, err := firestoreConn.Init(projectID, "IntentTest1")

	if err != nil {
		log.Println("firestore error: export GOOGLE_APPLICATION_CREDENTIALS=\"firestore_key.json\"")
		os.Exit(1)
	}
	defer client.Close()
	for _, item := range list {
		fmt.Println(reflect.TypeOf(item))
		fmt.Println(item)
		intfItem := structs.Map(item)
		firestoreConn.AddDoc(intfItem)
	}

}