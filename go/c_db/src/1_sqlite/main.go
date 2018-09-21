package main
//https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	fmt.Println("sqlite")
	os.Remove("./my.db")

	db, _ := sql.Open("sqlite3", "./my.db")
	defer db.Close()
	fmt.Println("111")
	sqlStmt := `
	create table fooTbl (id integer not null primary key, name text);
	delete from fooTbl;
	`

	_, err := db.Exec(sqlStmt)
	fmt.Println("222",err)

	tx, err := db.Begin()
	sqlStmt = "insert into fooTbl(name) values ('Ian Zhang')"
	_, err = db.Exec(sqlStmt)
	sqlStmt = "insert into fooTbl(name) values ('Kevin Zhang')"
	_, err = db.Exec(sqlStmt)
	tx.Commit()
	sqlStmt = "insert into fooTbl(name) values ('Lucas Zhang')"
	_, err = db.Exec(sqlStmt)
	sqlStmt = "update fooTbl set name = 'Hong Sun' WHERE id =3"
	_, err = db.Exec(sqlStmt)
	sqlStmt = "delete from fooTbl WHERE id =3"
	_, err = db.Exec(sqlStmt)
	fmt.Println("333",err)

	rows, err :=db.Query("select * from fooTbl")

	defer rows.Close()
	fmt.Println("444",err)
	for rows.Next() {
		var id int
		var name string
		_ = rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}