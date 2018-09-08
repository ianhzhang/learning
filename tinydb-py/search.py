from tinydb import TinyDB, Query
db = TinyDB("./myDB.json")


tblBook = db.table('book')

myQuery = Query()

goBook = tblBook.search(myQuery.name=='Go')
print goBook