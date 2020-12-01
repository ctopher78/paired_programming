
"""
Note: Run from root directory in order to find file
"""

import re
import sys
from collections import OrderedDict

def main():
    with open("input/ip-access-list.log", "r") as _f:
        acl = _f.readlines()

    matches = list()
    denied = re.compile(f"denied \w+ (\S+) -> \S+, (\d+) packet (\d+)")
    for line in acl:
        match = denied.search(line)
        if match:
           matches.append(
               {
                   "source": match.group(1),
                   "packets": int(match.group(2)), 
                   "bytes": int(match.group(3))
                }
            )

    sorted_matches = sorted(matches, key=lambda x: x["bytes"])

    for match in sorted_matches:
        print(match["source"], match["packets"], match["bytes"])
    # for src, data in matches.items():
    #     print(src, data)

    



if __name__ == "__main__":
    sys.exit(main())