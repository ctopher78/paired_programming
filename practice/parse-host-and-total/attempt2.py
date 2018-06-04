
from collections import OrderedDict, Counter
import re

'''
count the number of unique hosts and the number of bytes for each host.  Sort by number of occurances then by bytes

host1 10 50
host2 6  200
host6 2  100

host2 6  200
host6 2  100
host1 10 50

...
'''

logs = '''burger.letters.com - - [01/Jul/1995:00:00:11 -0400] "GET /shuttle/countdown/liftoff.html HTTP/1.0" 304 0
unicomp6.unicomp.net - - [01/Jul/1995:00:00:06 -0400] "GET /shuttle/countdown/ HTTP/1.0" 200 3985
burger.letters.com - - [01/Jul/1995:00:00:11 -0400] "GET /shuttle/countdown/liftoff.html HTTP/1.0" 304 0
d104.aa.net - - [01/Jul/1995:00:00:13 -0400] "GET /shuttle/countdown/ HTTP/1.0" 200 3985
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /shuttle/countdown/count.gif HTTP/1.0" 200 40310
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /images/NASA-logosmall.gif HTTP/1.0" 200 786
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /images/KSC-logosmall.gif HTTP/1.0" 200 1204
d104.aa.net - - [01/Jul/1995:00:00:15 -0400] "GET /shuttle/countdown/count.gif HTTP/1.0" 200 40310
d104.aa.net - - [01/Jul/1995:00:00:15 -0400] "GET /images/NASA-logosmall.gif HTTP/1.0" 200 786'''

log_lines = logs.split("\n")

hosts = OrderedDict()

regx = re.compile(r'^(\S+) .* (\d+)$')

for line in log_lines:
    line = line.strip()

    match = regx.match(line)
    if not match:
        continue
    
    name, byte_c = match.groups()
    if hosts.has_key(name):
        hosts[name]["count"] += 1
        hosts[name]["byte_c"] += int(byte_c)
        continue
    
    hosts[name] = {"count": 1, "byte_c": int(byte_c)}
        
print sorted(hosts.items(), key=lambda k: k[1]["count"], reverse=True)
