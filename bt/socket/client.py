"""
A simple Python script to send messages to a sever over Bluetooth using
Python sockets (with Python 3.3 or above).
"""

import socket

# serverMACAddress = '14:4F:8A:06:45:EA' # nuc
serverMACAddress = '00:1A:7D:DA:71:11'

port = 22
s = socket.socket(socket.AF_BLUETOOTH, socket.SOCK_STREAM, socket.BTPROTO_RFCOMM)
s.connect((serverMACAddress,port))
while 1:
    text = input()
    print("text:", text)
    if text == "quit":
        break
    s.send(bytes(text, 'UTF-8'))
s.close()
