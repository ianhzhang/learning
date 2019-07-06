import bluetooth

print "start"
server_sock=bluetooth.BluetoothSocket( bluetooth.RFCOMM )
print "bind"
server_sock.bind(("",bluetooth.PORT_ANY))

print "listen"
server_sock.listen(1)

print "advertise"
uuid = "1e0ca4ea-299d-4335-93eb-27fcfe7fa848"
bluetooth.advertise_service(server_sock, "yike-amd", uuid)

print "accept"
client_sock, address = server_sock.accept()
print "Accepted connection from ",address

data = client_sock.recv(1024)
print "received [%s]" % data
print "send back -----"
client_sock.send(data)

client_sock.close()
