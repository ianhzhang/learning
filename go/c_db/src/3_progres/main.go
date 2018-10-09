package main
// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
// https://flaviocopes.com/golang-tutorial-rest-api/
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"reflect"
)

const (
	host     = "35.232.138.124"
	port     = 5432
	dbname   = "mydb"
  )


func printTbl(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM fooTbl")
	if err != nil {
		panic(err)
	}
	var id int
	var name string

	for rows.Next() {
		fmt.Println("---------------")
		err = rows.Scan(&id,&name)
		fmt.Println(id, name)
	}
	fmt.Println("---------------")
}

func main() {
	fmt.Println("start")
	user_id := os.Getenv("USER")
	password := os.Getenv("PASS")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				host, port, user_id, password, dbname)
	fmt.Println(psqlInfo)

	fmt.Println("sql open")
	db, err := sql.Open("postgres", psqlInfo)
	fmt.Println(reflect.TypeOf(db))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("sql ping")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Successfully connected!")

	fmt.Println("Drop fooTble")
	_, err = db.Exec("drop table if exists fooTbl;")
	if err!= nil {
		panic(err)
	}

	fmt.Println("Create fooTble")
	_, err = db.Exec("create table fooTbl (id SERIAL primary key, name text);")
	if err!= nil {
		panic(err)
	}

	fmt.Println("Insert fooTble")
	_, err = db.Exec("insert into fooTbl(name) values ('Ian Zhang');")
	if err!= nil {
		panic(err)
	}
	fmt.Println("Insert fooTble")
	_, err = db.Exec("insert into fooTbl(name) values ('Kevin Zhang');")
	if err!= nil {
		panic(err)
	}
	fmt.Println("Insert fooTble")
	_, err = db.Exec("insert into fooTbl(name) values ('Lucas Zhang');")
	if err!= nil {
		panic(err)
	}
	printTbl(db)

	fmt.Println("Update fooTble")
	_, err = db.Exec("update fooTbl set name = 'Hong Sun' WHERE id =3;")
	if err!= nil {
		panic(err)
	}
	printTbl(db)

	fmt.Println("Delete from fooTble")
	_, err = db.Exec("delete from fooTbl WHERE id =3;")
	if err!= nil {
		panic(err)
	}
	printTbl(db)
}