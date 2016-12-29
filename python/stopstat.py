#!/usr/bin/env python

import os
import sys
import json
import time
from datetime import datetime,timedelta


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


def pretty_timestamp(t):
    #
    # python2's strptime doesn't handle %z so here's a workaround, from:
    #   http://stackoverflow.com/questions/1101508/how-to-parse-dates-with-0400-timezone-string-in-python/23122493#23122493
    #
    ret = datetime.strptime(t[0:16], '%Y-%m-%dT%H:%M')
    if t[18] == '+':
        ret += timedelta(hours=int(t[19:22]), minutes=int(t[23:]))
    elif t[18] == '-':
        ret -= timedelta(hours=int(t[19:22]), minutes=int(t[23:]))

    return ret.strftime("%I:%M%p")

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
