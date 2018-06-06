
# Extract number of times a url has been hit and the number and count the number of times that url returned a non-200 status code

from collections import OrderedDict
import re

log = '''unicomp6.unicomp.net - - [01/Jul/1995:00:00:06 -0400] "GET /shuttle/countdown/ HTTP/1.0" 200 3985
burger.letters.com - - [01/Jul/1995:00:00:11 -0400] "GET /shuttle/countdown/liftoff.html HTTP/1.0" 304 0
burger.letters.com - - [01/Jul/1995:00:00:12 -0400] "GET /images/NASA-logosmall.gif HTTP/1.0" 304 0
burger.letters.com - - [01/Jul/1995:00:00:12 -0400] "GET /shuttle/countdown/video/livevideo.gif HTTP/1.0" 200 0
d104.aa.net - - [01/Jul/1995:00:00:13 -0400] "GET /shuttle/countdown/ HTTP/1.0" 200 3985
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /shuttle/countdown/count.gif HTTP/1.0" 200 40310
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /images/NASA-logosmall.gif HTTP/1.0" 200 786
unicomp6.unicomp.net - - [01/Jul/1995:00:00:14 -0400] "GET /images/KSC-logosmall.gif HTTP/1.0" 200 1204
d104.aa.net - - [01/Jul/1995:00:00:15 -0400] "GET /shuttle/countdown/count.gif HTTP/1.0" 200 40310
d104.aa.net - - [01/Jul/1995:00:00:15 -0400] "GET /images/NASA-logosmall.gif HTTP/1.0" 200 786'''


log_lines = log.split("\n")

urls = OrderedDict()

regx = re.compile(r'.*\"\w+ (\S+) \S+\" (\d+)')

for line in log_lines:
    match = regx.match(line)
    if not match:
        continue

    url, status_code = match.groups()

    if urls.has_key(url):
        urls[url]["count"] += 1
        if status_code != 200:
            urls[url]["errors"] += 1
        continue

    urls[url] = dict(count=1, errors=0)

sorted_urls = sorted(urls.items(), key=lambda k: k[1]["errors"], reverse=True)

for u in sorted_urls:
    print u[0], u[1]["count"], u[1]["errors"]