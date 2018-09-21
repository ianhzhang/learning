package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	DRIVER_NAME = "mysql"
	DATA_SOURCE_NAME = "ihz_user:abc123@tcp(35.232.174.244:3306)/mydb"
)

func main() {
	fmt.Println("start")
	db, err := sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	fmt.Println("111:", err)
	defer db.Close()

	sqlStmt := "drop table if exists fooTbl;"
	_, err = db.Exec(sqlStmt)
	fmt.Println("222:", err)

	sqlStmt = "create table fooTbl (id integer not null primary key AUTO_INCREMENT, name text);"
	_, err = db.Exec(sqlStmt)
	fmt.Println("333:", err)

	// Inserts
	sqlStmt = "insert into fooTbl(name) values ('Ian Zhang')"
	_, err = db.Exec(sqlStmt)
	fmt.Println("444:", err)
	sqlStmt = "insert into fooTbl(name) values ('Kevin Zhang')"
	_, err = db.Exec(sqlStmt)
	sqlStmt = "insert into fooTbl(name) values ('Lucas Zhang')"
	_, err = db.Exec(sqlStmt)
	
	// Update
	sqlStmt = "update fooTbl set name = 'Hong Sun' WHERE id =3"
	_, err = db.Exec(sqlStmt)
	
	// Delete
	sqlStmt = "delete from fooTbl WHERE id =3"
	_, err = db.Exec(sqlStmt)
	fmt.Println("444",err)

	// Select
	rows, err :=db.Query("select * from fooTbl")
	defer rows.Close()
	fmt.Println("555",err)
	for rows.Next() {
		var id int
		var name string
		_ = rows.Scan(&id, &name)
		fmt.Println(id, name)
	}



}