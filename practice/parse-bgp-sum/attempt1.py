import sys
import re

def main():
    bgpsum = bgp_sum.split("\n")

    peers = list()

    regx = re.compile(r'^(\d+[.:]\S+)\s+\d\s+(\d+).* (\S+)$')

    for line in bgpsum:
        match = regx.match(line)
        if not match:
            continue

        ip, asn, status = match.groups()
        peers.append(dict(ip=ip, asn=asn, status=status))

    for peer in peers:
        if not peer["status"].isdigit():
            print peer["ip"], peer["asn"], peer["status"]



bgp_sum = '''R1>show bgp ipv4 unicast summary
BGP router identifier 10.0.64.1, local AS number 11111
BGP table version is 456732550, main routing table version 456732550
424983 network entries using 51847926 bytes of memory
525139 path entries using 27307228 bytes of memory
78207/72451 BGP path/bestpath attribute entries using 5943732 bytes of memory
6 BGP rrinfo entries using 144 bytes of memory
64596 BGP AS-PATH entries using 2210108 bytes of memory
4342 BGP community entries using 330426 bytes of memory
13 BGP extended community entries using 328 bytes of memory
0 BGP route-map cache entries using 0 bytes of memory
0 BGP filter-list cache entries using 0 bytes of memory
BGP using 87639892 total bytes of memory
Dampening enabled. 209 history paths, 111 dampened paths
28 received paths for inbound soft reconfiguration
BGP activity 8379913/7954718 prefixes, 488102372/487576978 paths, scan interval 60 secs

Neighbor        V    AS MsgRcvd MsgSent   TblVer  InQ OutQ Up/Down  State/PfxRcd
10.224.0.6      4 65013   37281   36372 456732556    0    0 2d03h           2
10.0.30.53     4 22222 2825586   19277 456732541    0    0 1w5d       159648
10.0.224.129   4 11111 2870058 2061843 456732556    0    0 1w5d        52904
10.0.224.135   4 11111 2451224 2061733 456732556    0    0 1w5d       262535
10.0.224.137   4 11111   19280 3808184 456732556    0    0 1w5d           30
10.0.224.138   4 11111   19323 3808184 456732556    0    0 1w5d          110
10.0.224.141   4 11111 1792201 3808355 456732556    0    0 1w5d        49596
Neighbor        V    AS MsgRcvd MsgSent   TblVer  InQ OutQ Up/Down  State/PfxRcd
10.0.227.75    4 65003  102654  110352 456732541    0    0 1w5d           11
10.0.228.17    4 65010   37280   36372 456732566    0    0 2d03h           8
10.0.229.38    4 65007       0       0        0    0    0 11w5d    Idle
10.0.229.223   4 65002   37281   36372 456732566    0    0 2d03h           1
10.0.234.133   4 65005  105138  110376 456732541    0    0 1w5d            2
10.0.234.135   4 65006  105134  110368 456732541    0    0 1w5d            2
10.0.241.145   4 65011   37282   36373 456732566    0    0 2d03h           2
10.100.1.1      4   200      26      22      199    0    0 00:14:23        23
10.200.1.1      4   300      21      51      199    0    0 00:13:40        0
2001:1890:ff:ffff:12:122:83:238 6 300      21      51      199    0    0 00:13:40        connect'''

if __name__ == "__main__":
    sys.exit(main())