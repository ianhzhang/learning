from tinydb import TinyDB, Query
db = TinyDB("./myDB.json")


tblBook = db.table('book')

myQuery = Query()
print tblBook.all()
goBook = tblBook.remove(myQuery.name=='Go')
print "-----------"
print tblBook.all()