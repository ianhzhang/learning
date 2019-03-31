import socket

# hostMACAddress = '14:4F:8A:06:45:EA'  #nuc
hostMACAddress = '00:1A:7D:DA:71:11'    #yike

# The MAC address of a Bluetooth adapter on the server. The server might have multiple Bluetooth adapters.

port = 3 # 3 is an arbitrary choice. However, it must match the port used by the client.
backlog = 1
size = 1024

s = socket.socket(socket.AF_BLUETOOTH, socket.SOCK_STREAM, socket.BTPROTO_RFCOMM)
s.bind((hostMACAddress,port))
s.listen(backlog)
try:
    client, address = s.accept()
    while 1:
        data = client.recv(size)
        if data:
            print(data)
            client.send(data)
except:	
    print("Closing socket")	
    client.close()
    s.close()
