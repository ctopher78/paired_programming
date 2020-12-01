
import sys
import re



def main():
    with open("input/show-ip-bgp.log", "r") as _f:
        table = _f.readlines()

    prefix = re.compile(f"^(\d+\S+)\s+.*")
    as_path = re.compile(f"^\s+AS path:\s(.*)\sI,.*")

    prefixes = dict()
    current_prefix = str
    for row in table:
        prefix_row = prefix.match(row)
        aspath_row = as_path.match(row)
        if prefix_row:
            current_prefix = prefix_row.group(1)
            prefixes[current_prefix] = list()
        if aspath_row:
            prefixes[current_prefix].append(aspath_row.group(1))
    
    for prefix, paths in prefixes.items():
        print(f"\nPrefix: {prefix}")
        for path in paths:
            print(path)


if __name__ == "__main__":
    sys.exit(main())