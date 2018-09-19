import json
from tinymongo import TinyMongoClient
from pprint import pprint

# Initial db
connection = TinyMongoClient('defender')    # create defender/ directory
defenderDb = connection.defenderDb  # create defender/defenderDb.json
clientCollection_Tbl = defenderDb.client    # create client collection(table)


clients = []
with open('data-client.json') as f:
    clients = json.load(f)
    print len(clients)

for client in clients:
    client_item = {
        'macAddress': client['macAddress'].replace(" ",""),
        'accessPointSerialNumber': client['accessPointSerialNumber'],
        'lastSeen': client['lastSeen']
    }
    print client_item

    #found = clientCollection_Tbl.find_one({"macAddress": client_item["macAddress"], "accessPointSerialNumber": client['accessPointSerialNumber'], "lastSeen": client['lastSeen']})
    found = clientCollection_Tbl.find_one(client_item)
    found_cnt = clientCollection_Tbl.find(client_item).count()
    print "=================="
    print found_cnt
    print "------------------"
    if found_cnt ==0 :
        print "insert_one"
        clientCollection_Tbl.insert_one(client_item)

    # find '00:50:56:AB:31:20'
    # 
    print "*************************************************************************"
    mac_items = clientCollection_Tbl.find({"macAddress": '00:50:56:AB:31:20'})
    for item in mac_items:
        print item