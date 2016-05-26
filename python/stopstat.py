#!/usr/bin/env python

import os
import sys
import json
import time


if len(sys.argv) != 2:
    print("usage: stopstat.py stop-code")
    sys.exit(1)

stop_code = sys.argv[1].upper()

MetLinkAPIv1StopDeparturesUrl = "https://www.metlink.org.nz/api/v1/StopDepartures/" + stop_code


def service_summary(service):
    realtime = "(sched)"
    if service['IsRealtime']:
        realtime = "(real)"

    return "    {} {} {} to {}".format(
        pretty_timestamp(service['DisplayDeparture']),
        realtime,
        service['OriginStopName'],
        service['DestinationStopName'],
    )


def pretty_timestamp(rfc339_timestamp):
    t = time.strptime(rfc339_timestamp, "%Y-%m-%dT%H:%M:%S+12:00")
    return time.strftime("%I:%M%p", t)

try:
    # python2
    import urllib2 as ur
except ImportError:
    # python3
    import urllib.request as ur

try:
    response = ur.urlopen(MetLinkAPIv1StopDeparturesUrl)
    raw_json = response.read()
except:
    print "failed, maybe a bad stop code?"
    sys.exit(1)

if type(raw_json) != str:
    raw_json = raw_json.decode()

try:
    stop_departures = json.loads(raw_json)
except:
    print "failed parsing json, maybe a bad stop code?"
    sys.exit(1)


print(stop_departures.get('Stop',{}).get('Name'))
if stop_departures.get('Notices'):
    print ("Notices:")

    for notice in stop_departures['Notices']:
        print("    " + notice.get('LineNote'))

print("Services:")
for service in stop_departures['Services']:
    print(service_summary(service))
