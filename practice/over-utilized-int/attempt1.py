
import re

ints = ['''GigabitEthernet0/0 is up, line protocol is up 
    Hardware is 82543 (Livengood), address is 00d0.ffb6.4c00 (bia 00d0.ffb6.4c00)
    Internet address is 10.1.1.3/8
    MTU 1500 bytes, BW 1000000 Kbit, DLY 10 usec, 
       reliability 255/255, txload 1/255, rxload 1/255
    Encapsulation ARPA, loopback not set
    Keepalive set (10 sec)
    Full-duplex mode, link type is autonegotiation, media type is SX
    output flow-control is on, input flow-control is on
    ARP type:ARPA, ARP Timeout 04:00:00
    Last input 00:00:04, output 00:00:03, output hang never
    Last clearing of "show interface" counters never
    Queueing strategy:fifo
    Output queue 0/40, 0 drops; input queue 0/75, 0 drops
    5 minute input rate 898000000 bits/sec, 0 packets/sec
    5 minute output rate 0 bits/sec, 0 packets/sec
       2252 packets input, 135120 bytes, 0 no buffer
       Received 2252 broadcasts, 0 runts, 0 giants, 0 throttles
       55 input errors, 0 CRC, 0 frame, 0 overrun, 0 ignored
       0 watchdog, 0 multicast, 0 pause input
       0 input packets with dribble condition detected
       2631 packets output, 268395 bytes, 0 underruns
       0 output errors, 0 collisions, 2 interface resets
       0 babbles, 0 late collision, 0 deferred
       0 lost carrier, 0 no carrier, 0 pause output
       0 output buffer failures, 0 output buffers swapped out''',
	'''GigabitEthernet0/1 is up, line protocol is up 
    Hardware is 82543 (Livengood), address is 00d0.ffb6.4c00 (bia 00d0.ffb6.4c00)
    Internet address is 10.1.1.5/8
    MTU 1500 bytes, BW 1000000 Kbit, DLY 10 usec, 
       reliability 255/255, txload 1/255, rxload 1/255
    Encapsulation ARPA, loopback not set
    Keepalive set (10 sec)
    Full-duplex mode, link type is autonegotiation, media type is SX
    output flow-control is on, input flow-control is on
    ARP type:ARPA, ARP Timeout 04:00:00
    Last input 00:00:04, output 00:00:03, output hang never
    Last clearing of "show interface" counters never
    Queueing strategy:fifo
    Output queue 0/40, 0 drops; input queue 0/75, 0 drops
    5 minute input rate 2000000 bits/sec, 10000000 packets/sec
    5 minute output rate 850050000 bits/sec, 0 packets/sec
       2252 packets input, 135120 bytes, 0 no buffer
       Received 2252 broadcasts, 0 runts, 0 giants, 0 throttles
       0 input errors, 30 CRC, 0 frame, 0 overrun, 0 ignored
       0 watchdog, 0 multicast, 0 pause input
       0 input packets with dribble condition detected
       2631 packets output, 268395 bytes, 0 underruns
       0 output errors, 0 collisions, 2 interface resets
       0 babbles, 0 late collision, 0 deferred
       0 lost carrier, 0 no carrier, 0 pause output
       0 output buffer failures, 0 output buffers swapped out''']


regx = re.compile('.*Ethernet(\d+/\d+).*input rate (\d+).*output rate (\d+).* (\d+) input errors, (\d+) CRC', re.S)
ifaces = list()

for i in ints:
    match = regx.match(i).groups()
    if match:
        name, inRate, outRate, inError, crcError = regx.match(i).groups()

    ifaces.append(dict(name=name, inRate=int(inRate), outRate=int(outRate), inError=int(inError), crcError=int(crcError)))

sort_iface = sorted(ifaces, key=lambda k: k["outRate"], reverse=True)

for i in sort_iface:
    print "int: {iface} in: {in_rate} out: {out_rate}".format(iface=i["name"], in_rate=i["inRate"], out_rate=i["outRate"])

for i in ifaces:
    if i["inError"] > 1:
        print i["name"], " has errors"