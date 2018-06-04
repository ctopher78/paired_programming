
from collections import OrderedDict, Counter
import re

'''
count the number of unique hosts and sort hosts by number of occurances

host1 10
host2 6
host6 2
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

regx = re.compile(r'^(\S+) ')

for i in log_lines:
    match = regx.match(i)
    if match and not hosts.has_key(match.group(1)):
        hosts[match.group(1)] = 1
        continue
    if match:
        hosts[match.group(1)] += 1

sorted_hosts = sorted(hosts.items(), key=lambda t: t[1], reverse=True)

print sorted_hosts