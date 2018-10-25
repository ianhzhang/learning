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
    conn = psycopg2.connect(host='35.232.138.124', dbname='mydb', user=user, password=passwd, connect_timeout=5)
except:
    print "Cannot connect "

if conn:
    mycursor = conn.cursor()
    printTbl()



