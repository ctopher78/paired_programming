
import re
import json

show_int = '''Interface     IP-Address     OK?  Method  Status                  Protocol
Ethernet0     10.108.00.5    YES  NVRAM   up                      up      
Ethernet1     unassigned     YES  unset   administratively down   down    
Loopback0     10.108.200.5   YES  NVRAM   up                      up      
Serial0       10.108.100.5   YES  NVRAM   up                      up      
Serial1       10.108.40.5    YES  NVRAM   up                      up      
Serial2       10.108.100.5   YES  manual  proto                   down     
Serial3       unassigned     YES  unset   administratively down   down'''


show_int = show_int.split("\n")

ifaces = list()

regx = re.compile(r'^([ELS]\S+)\s+(\S+)\s+\S+\s+\S+\s+(\w+(?: \w+)?)\s+(\w+)')

for line in show_int:
    line = line.strip()

    match = regx.match(line)

    if not match:
        continue
    if match:
        iface, ip, status, proto = match.groups()
        ifaces.append(dict(iface=iface, ip=ip, status=status, proto=proto))

for i in ifaces:
    print i["iface"], i["status"]

print "\ndown interfaces:"
for i in ifaces:
    if i["status"] != "up":
        print i["iface"]

print "\nSorted interfaces:"
iface_sorted = sorted(ifaces, key=lambda k: k["status"])

for i in iface_sorted:
    print i["iface"], i["status"], i["proto"]
    
print json.dumps(ifaces, indent=2)
