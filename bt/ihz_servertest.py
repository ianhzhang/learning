import bluetooth

print "start"
server_sock=bluetooth.BluetoothSocket( bluetooth.RFCOMM )
print "bind"
server_sock.bind(("",bluetooth.PORT_ANY))

print "listen"
server_sock.listen(1)

print "advertise"
bluetooth.advertise_service(server_sock, "yike-amd",
                     service_classes=[bluetooth.SERIAL_PORT_CLASS],
                     profiles=[bluetooth.SERIAL_PORT_PROFILE])

print "accept"
client_sock, address = server_sock.accept()
print "Accepted connection from ",address

data = client_sock.recv(1024)
print "received [%s]" % data

client_sock.close()