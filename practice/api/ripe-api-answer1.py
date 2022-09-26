# This is an exercise that tests basic API calling and response handling using Python.
#
# The exercise uses a public API provided by RIPE, which is the "Regional Internet Registry (RIR) that governs IPv4 address trade in Europe".

# Tasks:
# 1. Call ripe API endpoint for the prefix passed in by sys.stdin
# 2. Save the JSON output from requests 'Response' object
# 3. Parse the JSON output and collect the locations for each RRC (Remote Route Collector) in the output
# 4. Sort the list of locations alphabetically and print the locations

### example output ###
#
# Amsterdam, Netherlands
# Amsterdam, Netherlands
# Amsterdam, Netherlands
# Barcelona, Spain
# Bucharest, Romania
# Dubai, UAE
# Frankfurt, Germany
# Geneva, Switzerland
# Johannesburg, South Africa
# London, United Kingdom
# Miami, Florida, US
# Milan, Italy
# Montevideo, Uruguay
# Moscow, Russian Federation
# New York City, New York, US
# ...

# Focus on R.R.F 
# Make it..Run
# Make it..Right (idiomatic (i.e. Pythonic), readable, error checking, etc)
# Make it..Fast

### Code Starts Here ###

import sys
import json

import requests


prefix = sys.argv[1]
URL = f'https://stat.ripe.net/data/looking-glass/data.json?resource={prefix}&look_back_limit=30'


def main():
    try:
        resp = requests.get(URL)
    except requests.ConnectTimeout as e:
        print(e)
    except requests.ConnectionError as e:
        print(e)
    except requests.RequestException as e:
        print("Generic requests error: ", e)

    if resp.status_code != 200:
        print("Connection error: ", resp.json)
        sys.exit(resp.status_code)
        

    locs =  [rrc["location"] for rrc in json.loads(resp.text)["data"]["rrcs"]]

    for loc in sorted(locs):
        print(loc)


    resp.close()


if __name__ == "__main__":
    main()