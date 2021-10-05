from IPy import IP

ips = IP("192.168.1.0/24")
for ip in ips:
    print(ip)