package firestore_api
// https://firebase.google.com/docs/firestore/manage-data/add-data
import (
	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type JsonMapStruct map[string]interface{}

type FirestoreConn struct {
	my_client     *firestore.Client
	my_ctx        context.Context
	my_collection *firestore.CollectionRef
}

func (conn *FirestoreConn) Init(projectID string, collectionID string) (*firestore.Client, error) {
	ctx := context.Background()
	conn.my_ctx = ctx

	client, err := firestore.NewClient(ctx, projectID)
	if err == nil {
		conn.my_client = client

		collection := client.Collection(collectionID)
		conn.my_collection = collection
	}

	return client, err
}

func (conn *FirestoreConn) AddDoc(doc map[string]interface{}) {
	conn.my_collection.Add(conn.my_ctx, doc)
}

func (conn *FirestoreConn) ClearAll() {
	iter := conn.my_collection.Documents(conn.my_ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		doc.Ref.Delete(conn.my_ctx)
	}
}

func (conn *FirestoreConn) IsExist(myId string) bool {
	found := false
	iter := conn.my_collection.Documents(conn.my_ctx)
	for {
		doc, err := iter.Next() // doc: *firestore.DocumentSnapshot
		if err == iterator.Done {
			break
		}
		docData := doc.Data() // map[string]interface {}

		if docData["UUID"] == myId {
			found = true
			break
		}
	}
	return found
}

func (conn *FirestoreConn) GetById(myId string) JsonMapStruct {
	iter := conn.my_collection.Documents(conn.my_ctx)
	for {
		doc, err := iter.Next() // doc: *firestore.DocumentSnapshot
		if err == iterator.Done {
			break
		}
		docData := doc.Data() // map[string]interface {}

		if docData["UUID"] == myId {
			return docData
		}
	}
	return nil
}

func (conn *FirestoreConn) GetAll() []JsonMapStruct {
	returnList := make([]JsonMapStruct, 0)

	iter := conn.my_collection.Documents(conn.my_ctx)
	for {
		doc, err := iter.Next() // doc: *firestore.DocumentSnapshot
		if err == iterator.Done {
			break
		}
		docData := doc.Data() // map[string]interface {}
		returnList = append(returnList, docData)

	}
	return returnList
}

func (conn *FirestoreConn) UpdateDoc(newDoc map[string]interface{}) bool {
	myId := newDoc["UUID"]

	found := false
	iter := conn.my_collection.Documents(conn.my_ctx)
	for {
		doc, err := iter.Next() // doc: *firestore.DocumentSnapshot
		if err == iterator.Done {
			break
		}
		docData := doc.Data() // map[string]interface {}

		if docData["UUID"] == myId {
			found = true
			doc.Ref.Set(conn.my_ctx, newDoc)
			break
		}
	}
	return found
}

func (conn *FirestoreConn) DeleteDoc(myId string) {
	iter := conn.my_collection.Documents(conn.my_ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		docData := doc.Data() // map[string]interface {}
		if docData["UUID"] == myId {
			doc.Ref.Delete(conn.my_ctx)
		}

	}
}