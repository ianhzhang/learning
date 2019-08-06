
package myfirestore

import (
	"fmt"
	"os"
	"golang.org/x/net/context"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	//"github.com/mitchellh/mapstructure"
)

var (
	
)
type Person struct {
	Ln string	`json:"ln"`
	Fn string   `json:"fn"`
	Address struct {
		Street string `json:"street"`
		City string `json:"city"`
		State string `json:"state"`
		Zip string `json:"zip"`
	} `json:"address"`
}

func Test_firestore() {
	projectID := os.Getenv("GCLOUD_PROJECT")
	ctx := context.Background()
	fmt.Println("111:", projectID)
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("error 1 ")
		os.Exit(1)
	}
	defer client.Close()

	iter := client.Collection("TagContainer").Documents(ctx)
	defer iter.Stop()

	for {
		docIter, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Print("error 2 ")
		}
		println("-----------------------------");
		data := docIter.Data()
		fmt.Println(data)

		/*
		person := Person{}
		mapstructure.Decode(data, &person)
		fmt.Println(person)
		*/

		addr_docs,_ := docIter.Ref.Collection("address").Documents(ctx).GetAll()
		for _,  addr_document := range addr_docs {
			addrInfo := addr_document.Data()
			fmt.Println(addrInfo)
			fmt.Println(addrInfo["street"])
		}
	}
}
