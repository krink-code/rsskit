#!/usr/bin/env python3

import sys
import json

try:
    values = json.load(sys.stdin)
except ValueError as e:
    print(json.dumps({"Error":"invalid json data"}))
    sys.exit(1)

import urllib.parse
import urllib.request

if sys.argv[1:]:
    url = sys.argv[1]
else:
    print(json.dumps({"Error":"no url"}))
    sys.exit(1)

params = json.dumps(values).encode('utf8')

req = urllib.request.Request(url, data=params, headers={'content-type': 'application/json'})
with urllib.request.urlopen(req) as response:
    the_response = response.read()

print(str(the_response))


