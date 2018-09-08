
# https://www.w3schools.com/python/python_mongodb_update.asp

from pymongo import MongoClient

client = MongoClient('mongodb://localhost:27017/')
my_db = client["ihzdb1"]
my_family_collection = my_db["family"]

person = {
    "ln": "Zhang",
    "fn": "Ian"
}

result = my_family_collection.find_one(person)

if result is not None:
    print "find: ", result
    new_value = {"$set": {"fn": "Lucas"}}
    my_family_collection.update_one(result, new_value)
else:
    print "not find"

result = my_family_collection.find()
print "print all found ============"
for each in result:
    print each