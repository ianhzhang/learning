"""
A simple Python script to send messages to a sever over Bluetooth
using PyBluez (with Python 2).

Ian: When the port and MAC address match.  It works.
When the client quit.   The server quits also.
We don't need to have a connection setup before sending data.
I see at beginning, there is a connection refused, I tried to use 
computer to setup connection.  Then the error disappear even I close
the connection.
Start server first.
I used python 2 to do it.
"""

import bluetooth

# ncu 14:4F:8A:06:45:EA

serverMACAddress = '00:1A:7D:DA:71:11'
port = 20

print("open socket")
s = bluetooth.BluetoothSocket(bluetooth.RFCOMM)

print("connect ....")
s.connect((serverMACAddress, port))

while 1:
    print("please input:")
    text = raw_input() # Note change to the old (Python 2) raw_input
    print("input:", text)
    if text == "quit":
    	break
    s.send(text)
sock.close()
