from uuid import getnode as get_mac

# get_mac is mac address in int form 84440852124015
# hex(get_mac()) is the hex string 0x4ccc6afed16f
# [2:] is to remove the prefix 0x

mac_hex_str = hex( get_mac() )
mac_str=':'.join(format(s, '02x') for s in bytes.fromhex(mac_hex_str[2:]))
print(mac_str)

