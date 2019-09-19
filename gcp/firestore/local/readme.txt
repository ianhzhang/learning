console:

gopath
go get console1
go build console1
./console1

======================= Console =====================================================

1. Get client
  client := firestore.NewClient(cts, projectID)

2. Get Documents
  2.1 Get Collection (Table)
  2.2 From collection, get all Documents
  client.Collection("Collection_Name").Documents(ctx) 

3. Get Data from Documents
    data := doc.Data()

----------------------------------------------------------------------
ctx := context.Background()
client, err := firestore.NewClient(ctx, projectID)
defer client.Close()

// GetAll
iter := client.Collection("TagContainer").Documents(ctx)
docIter, err := iter.Next()
data := docIter.Data()


// Get sub collection
addr_docs,_ := docIter.Ref.Collection("address").Documents(ctx).GetAll()
for _,  addr_document := range addr_docs {
    addrInfo := addr_document.Data()
    fmt.Println(addrInfo)
    fmt.Println(addrInfo["street"])
}

=======================================================================================

======================= Cloud =====================================================
1. Get client

ctx := context.Background()
client, err := firestore.NewClient(ctx, projectID)
2. Get Documents
  2.1 Get Collection (Table)
    collection := client.Collection(collectionID)
  2.2 Get document from collection
    docs = collection.Documents(ctx)
3. Get Data 
    data := doc.Data()

======================= Flutter =====================================================

doc = Firestore.instance.collection("table_name").document('doc_id')
doc.get() => then( function(document) {
    print(document["name"])
})

result= await Firestore.instance.collection("table_name").getDocument();
List<DocumentSnapshot> documents = result.documents;
 
===========================================================
    // 1. get all
    QuerySnapshot docs = await myCollectionRef.getDocuments();
    List<DocumentSnapshot> result =  docs.documents;
    print(result);
    for (var d in result) {
      print("xxxxx----");
      print(d.data);
      if (d.data!=null) {
        print(d.data["ln"]);
      }
    }
    print("---------------------------1111");
    // 2. Get by id (async)
    var doc = Firestore.instance.collection("Family").document("NhRLqcnG9TuHC9HvvMT4");
    doc.get().then( (d) {
      print("1111aaaa ------------");
      print(d.data);
    });
    print("-----------------------------------222");

    // 3. get by id (snapshot,)
    DocumentReference myDocRef = myCollectionRef.document('NhRLqcnG9TuHC9HvvMT4');
    var mysnapshot = myDocRef.snapshots();
    print(mysnapshot.runtimeType);    //  _BroadcastStream<DocumentSnapshot>

    // 3.1    
    myDocRef.get().then((d) {
      print("2222aaaa ----------------");
      print(d.data);
    });
    
    // 3.2
    mysnapshot.forEach( (d) {
      print(" 3333 aaaa ------------------");
      var myData = d.data;
      print(myData);
      if (myData == null) {
        print("not exist");
      }
    });
