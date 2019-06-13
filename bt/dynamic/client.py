from bluetooth import *

services=find_service(name="helloService",
                            uuid=SERIAL_PORT_CLASS)

for i in range(len(services)):
   print(i)
   match=services[i]
   print(match["name"])
   if(match["name"]=="helloService"):
      print("match")
      port=match["port"]
      name=match["name"]
      host=match["host"]

      print(name, port, host)

      client_socket=BluetoothSocket( RFCOMM )
      print("connect")
      client_socket.connect((host, port))
      print("send")
      client_socket.send("Hello world")
      print("close")
      client_socket.close()
      print("end")