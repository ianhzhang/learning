from tinydb import TinyDB, Query

people = [
    {
        "name": "Ian Zhang",
        "address": {
            "street": "101 Birdwood Ct",
            "city": "Cary",
            "zip": 27509
        }
    },
    {
        "name": "Kevin Zhang",
        "address": {
            "street": "201 CapeCod Ct",
            "city": "Raleigh",
            "zip": 27611
        }
    }
]

book1 = {
    "name": "Go",
    "desp": "good go book"
}

book2 = {
    "name": "Java",
    "desp": "bad book"
}

db = TinyDB("./myDB.json")
db.purge()      # does purge table

tblPeople = db.table('people')
tblPeople.purge()
tblPeople.insert({
    "people":people
})

tblBook = db.table('book')
tblBook.purge()
tblBook.insert(book1)
tblBook.insert(book2)
 
