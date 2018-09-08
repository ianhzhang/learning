from pymongo import MongoClient

client = MongoClient('mongodb://localhost:27017/')
print client.list_database_names()

my_db = client["ihzdb1"]
my_family_collection = my_db["family"]

person = {
    "ln": "Zhang",
    "fn": "Ian"
}
result = my_family_collection.insert_one(person)

print result
print my_db.list_collection_names()
print client.list_database_names()

