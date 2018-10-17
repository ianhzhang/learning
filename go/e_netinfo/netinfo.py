import psutil


def get_v4_netinfo():
    iflist = []

    nics = psutil.net_if_addrs()

    for ifname in nics:
        mac_addr = ""
        ipv4_addr = ""
        ipv6_addr = ""

        for snicaddr in nics[ifname]:

            if snicaddr.family == 18:  # IPV4
                mac_addr = snicaddr.address
            if snicaddr.family == 2:
                ipv4_addr = snicaddr.address
            if snicaddr.family == 30:
                ipv6_addr = snicaddr.address

        if mac_addr != "" and ipv4_addr !="":
            iflist.append({"name": ifname, "mac": mac_addr, "ipv4": ipv4_addr})

    return iflist

if_list = get_v4_netinfo()
print if_list