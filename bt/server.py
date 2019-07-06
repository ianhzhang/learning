
import bluetooth
import datetime

print("start")
server_sock=bluetooth.BluetoothSocket( bluetooth.RFCOMM )

print("bind")
server_sock.bind(("",bluetooth.PORT_ANY))

print("listen")
server_sock.listen(1)

print("advertise")
uuid = "00001101-0000-1000-8000-00805F9B34FB"
bluetooth.advertise_service(server_sock, "yike-amd", uuid)

client_sock = None
while 1:
    try:
        print "accept ...."
        client_sock, address = server_sock.accept()
        print "Accepted connection from ",address

        while 1:
            data = client_sock.recv(1024)
            print "received [%s]" % data
            retStr = data + " " + str(datetime.datetime.now().second)
            print "send back  ----- " + retStr
            client_sock.send(retStr)
    except:	
        print("again: ")
        #break
        # client_sock.close()

if client_sock:  
    client_sock.close()
