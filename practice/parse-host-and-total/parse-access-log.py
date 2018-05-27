import sys
import re

logFile = "hosts_access_log_00.txt"

def main():
    
    with open(logFile, "r") as _in:
        logs = _in.readlines()

    host_re = re.compile(r'^(\S+) ')
    host_counts = dict()

    for line in logs:
        if 'HTTP' not in line:
            continue
        
        host = host_re.match(line)
        if not host.group():
            continue
        host = host.group(1)

        if not host_counts.has_key(host):
            host_counts[host] = 1
            continue
        
        host_counts[host] += 1
    
    with open("py_records_"+logFile, "w+") as _out:
        for host, count in host_counts.items():
            _out.write("{} {}\n".format(host, count))
            

if __name__ == "__main__":
    sys.exit(main())