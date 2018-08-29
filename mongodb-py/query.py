


from pymongo import MongoClient

client = MongoClient('mongodb://localhost:27017/')
my_db = client["ihzdb1"]
my_family_collection = my_db["family"]

person = {
    "ln": "Zhang",
    "fn": "Ian1"
}

result = my_family_collection.find_one(person)

if result is not None:
    print "find: ", result
else:
    print "not find"

result = my_family_collection.find(person)
print "print all found ============"
for each in result:
    print each