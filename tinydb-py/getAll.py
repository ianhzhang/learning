from tinydb import TinyDB, Query
db = TinyDB("./myDB.json")

print db.tables()
all_tables = db.tables()
for tblName in all_tables:
    print "-----tbl ----------"
    print tblName
    tbl = db.table(tblName)
    print tbl.all()
    print "==================="



