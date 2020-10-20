"""
Given a CIDR (ex. /24), print out a subnet mask or wildcard mask
"""
import sys


def ip_mask(cidr):
    # Convert to binary mask
    mask_len = 32 # 255.255.255.255 

    host_bits = "0"*(mask_len-cidr)
    net_bits = "1"*cidr
    binary = net_bits + host_bits

    # Split biary into 4 octets
    octets = list()
    octet = 8
    for i in range(0, mask_len, octet):
        print(i)
        octets.append(str(int(binary[i:octet+i], base=2)))
    
    print(".".join(octets))
    return(cidr)
    

def main():
    subnet = "1.1.1.1/16"
    cidr = int(subnet.split("/")[1])

    print(ip_mask(cidr))

if __name__ == "__main__":
    sys.exit(main())