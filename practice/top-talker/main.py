import sys
import operator


def main():
    with open("input/hosts_access_log_00.txt") as _f:
        logs = _f.readlines()

    hosts = []

    for line in logs:
        if "HTTP" not in line:
            continue
        
        parts = line.split(" ")
        name = parts[0]
        byts = int(parts[-1])
        hosts = ()

    print hosts
        

if __name__ == "__main__":
    sys.exit(main())