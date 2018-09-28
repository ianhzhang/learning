import psycopg2

# Need manually create mydb
conn = psycopg2.connect("dbname='mydb' user='postgres' host='35.232.138.124' password='abc123'")

print conn



def printTbl():
    print("--------------------")
    print("print table rows")
    mycursor.execute("select * from fooTbl")
    myresult = mycursor.fetchall()

    for x in myresult:
        print(x)

    print("--------------------")

mycursor = conn.cursor()

print("1. drop table fooTbl")
mycursor.execute("drop table if exists fooTbl;")

print("2. Create fooTbl table")
mycursor.execute("create table fooTbl (id SERIAL primary key, name text);")

print("3. Insert tble")
mycursor.execute("insert into fooTbl(name) values ('Ian Zhang')")
mycursor.execute("insert into fooTbl(name) values ('Kevin Zhang')")
mycursor.execute("insert into fooTbl(name) values ('Lucas Zhang')")
printTbl()

print("4. update row")
mycursor.execute("update fooTbl set name = 'Hong Sun' WHERE id =3")
printTbl()

print("5. delete row")
mycursor.execute("delete from fooTbl WHERE id =3")
printTbl()



