import sqlite3

print(sqlite3.version)
conn = sqlite3.connect("./my.db")

cursor = conn.cursor()

# create table
sqlstmt = "drop table if exists fooTbl;"
cursor.execute(sqlstmt)

sqlstmt = "create table fooTbl (id integer not null primary key, name text);"
cursor.execute(sqlstmt)

# insert item
sqlstmt = "insert into fooTbl(name) values ('Ian Zhang')"
cursor.execute(sqlstmt)
sqlstmt = "insert into fooTbl(name) values ('Kevin Zhang')"
cursor.execute(sqlstmt)
sqlstmt = "insert into fooTbl(name) values ('Lucas Zhang')"
cursor.execute(sqlstmt)

# update item
sqlstmt = "update fooTbl set name = 'Hong Sun' WHERE id =3"
cursor.execute(sqlstmt)

# delete item
sqlstmt = "delete from fooTbl WHERE id =3"
cursor.execute(sqlstmt)

# select
sqlstmt = "select * from fooTbl"
cursor.execute(sqlstmt)

rows = cursor.fetchall()
for row in rows:
    print(row)