import psycopg2
import os
# Need manually create mydb
user = os.environ['USER']
passwd = os.environ['PASS']


def printTbl():
    print("--------------------")
    print("print table rows")
    mycursor.execute("select * from fooTbl")
    myresult = mycursor.fetchall()

    for x in myresult:
        print(x)

    print("--------------------")

conn = None
try:
    conn = psycopg2.connect(host='35.232.138.124', dbname='mydb', user=user, password=passwd, connect_timeout=2.5)
except:
    print "Cannot connect "

if conn:
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

    print("4. select row")
    mycursor.execute("select * from fooTbl where id=1")
    rslt = mycursor.fetchall()
    print rslt
    print len(rslt)     ## if it is 0, that means find nothing.


    print("5. update row")
    mycursor.execute("update fooTbl set name = 'Hong Sun' WHERE id =3")
    printTbl()

    print("6. delete row")
    mycursor.execute("delete from fooTbl WHERE id =3")
    printTbl()



