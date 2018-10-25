import json

str1 = '{"Seq":1234,"Data":"{\"a\": 298202, \"c\": 258720}"}'

#str1 = '{"Seq":1234,"Data":"abd"}'
str1 = '{"seq":1234, "data":{"a": 298202, "c": 258720}}'
print str1
print type(str1)

item = json.loads(str1)
print item
print item["data"]
