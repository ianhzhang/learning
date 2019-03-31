from bluetooth import *

nearby_devices = discover_devices()
print(nearby_devices)

for address in nearby_devices:
    print(address)
    print( lookup_name(address))