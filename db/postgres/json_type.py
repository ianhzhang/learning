import psycopg2
import os
import json

# http://www.postgresqltutorial.com/postgresql-json/


# Need manually create mydb
user = os.environ['USER']
passwd = os.environ['PASS']

conn = psycopg2.connect(host='35.232.138.124', dbname='mydb', user=user, password=passwd)

def printTbl():
    print("--------------------")
    print("print table rows")
    mycursor.execute("select * from jsonTbl")
    myresult = mycursor.fetchall()

    for x in myresult:
        print(x)
        print type(x[1])    # dict  don't need json.loads
        y = x[1]
        if y is dict and "customer" in y.keys():
            print (y["customer"])
        else:
            print y
        print "------------"

    print("--------------------")

mycursor = conn.cursor()

print("1. drop table jsonTbl")
mycursor.execute("drop table if exists jsonTbl;")

print("2. Create jsonTbl table")
mycursor.execute("create table jsonTbl (id serial NOT NULL primary key, info json NOT NULL);")

print("3. Insert tble")
jData1 = {
    "customer": "John Doe1",
    "items": {
        "product": "Beer",
        "qty": 61
    }
}
jData2 = {
    "customer": "John Doe2",
    "items": {
        "product": "Beer",
        "qty": 62
    }
}
jData3 = {
    "customer": "John Doe3",
    "items": {
        "product": "Beer",
        "qty": 63
    }
}
jData4 = ["Java", "Python", "C++"]


mycursor.execute("insert into jsonTbl (info) values(%s)", [json.dumps(jData1)])
mycursor.execute("insert into jsonTbl (info) values(%s)", [json.dumps(jData2)])
printTbl()


print("4. update row")
mycursor.execute("update jsonTbl set info = (%s) WHERE id =1", [json.dumps(jData4)])
printTbl()




