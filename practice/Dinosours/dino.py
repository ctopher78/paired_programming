"""
The first file contains statistics about various dinosaurs. The second file contains additional data.
Given the following formula, speed = ((STRIDE_LENGTH / LEG_LENGTH) - 1) * SQRT(LEG_LENGTH * g)
Where g = 9.8 m/s^2 (gravitational constant)

Write a program to read in the data files from disk, it must then print the names of only the bipedal dinosaurs from fastest to slowest.
Do not print any other information.
"""



# def merge_info(files) -> dict:
#     dino_list = list()

#     with open(file, "r") as f:


import csv
import sys

from math import sqrt
from collections import OrderedDict

g = 9.8**2

def main():
    dinos = dict()

    # dininfo = merge_info(files)
    with open("dino1.csv") as f1:
        dino = csv.DictReader(f1)
        for d in dino:
            dinos[d['NAME']] = d

    with open("dino2.csv") as f2:
        dino2 = csv.DictReader(f2)
        for d in dino2:
            try:
                dinos[d["NAME"]].update(d)
            except KeyError:
                pass

    for v in dinos.values():
        if not v.get("STRIDE_LENGTH") and v.get("LEG_LENGTH"):
            dinos.pop(v["NAME"])
            continue

        dinos[v["NAME"]]["speed"] = ((float(v.get("STRIDE_LENGTH")) / float(v.get("LEG_LENGTH"))) - 1) * sqrt(float(v.get("LEG_LENGTH")) * g)

    sorted_dinos = OrderedDict(sorted(dinos.items(), key=lambda x: x[1]["speed"]))           

    # import json; print(json.dumps(sorted_dinos.items(), indent=2))

    for d in sorted_dinos.items():
        print(d)
        if d.get("STANCE") == "bipedal":
            print(d["NAME"])


if __name__ == "__main__":
    sys.exit(main())