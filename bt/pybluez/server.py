"""
A simple Python script to receive messages from a client over
Bluetooth using PyBluez (with Python 2).
"""

import bluetooth

# The MAC address of a Bluetooth adapter on the server. The server might have multiple Bluetooth adapters.
# hostMACAddress = '14:4F:8A:06:45:EA' 

hostMACAddress = '00:1A:7D:DA:71:11' 

port = 20
backlog = 1
size = 1024

print("open socket")
s = bluetooth.BluetoothSocket(bluetooth.RFCOMM)

print("bind address")
s.bind((hostMACAddress, port))

print("listen")
s.listen(backlog)

try:
    print("accepting ....")
    client, clientInfo = s.accept()
    print("accepted ....")
    while 1:
        data = client.recv(size)
        if data:
            print(data)
            client.send(data) # Echo back to client
except:	
    print("Closing socket")
    client.close()
    s.close()
