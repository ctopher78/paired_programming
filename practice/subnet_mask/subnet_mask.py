"""
Given a CIDR (ex. /24), print out a subnet mask or wildcard mask

conver the CIDR to a string binary mask (network bits 1 and host bits 0)
"""
import sys

MASK_LEN = 32  # 255.255.255.255 

def ip_mask(cidr):

    # Convert to binary mask
    host_bits = "0"*(MASK_LEN-cidr)
    net_bits = "1"*cidr
    binary = net_bits + host_bits
    print("BINARY: ", binary)

    # Split biary into 4 octets
    octets = list()
    octet = 8
    for i in range(0, MASK_LEN, octet):
        print("octet: ", i)
        octets.append(str(int(binary[i:octet+i], base=2)))
    print("Octets: ", octets)

    mask = ".".join(octets)
    return(mask)
    

def main():
    subnet = "1.1.1.1/30"
    cidr = int(subnet.split("/")[1])

    print("Mask: ", ip_mask(cidr))

if __name__ == "__main__":
    sys.exit(main())