
# https://www.w3schools.com/python/python_mongodb_delete.asp

from pymongo import MongoClient

client = MongoClient('mongodb://localhost:27017/')
my_db = client["ihzdb1"]
my_family_collection = my_db["family"]

person = {
    "ln": "Zhang",
    "fn": "Kevin"
}

result = my_family_collection.find_one(person)

if result is not None:
    print "find: ", result
    my_family_collection.delete_one(person)
else:
    print "not find"

result = my_family_collection.find()
print "print all found ============"
for each in result:
    print each