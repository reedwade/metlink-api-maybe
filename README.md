# metlink-api-maybe

Some code I've got working that uses an undocumented API found inside https://www.metlink.org.nz/

There seems to be one call of interest which, given a train or bus stop code returns a reasonably understandable whack of JSON.

`https://www.metlink.org.nz/api/v1/StopDepartures/JOHN`

Where, `JOHN` in this case is the J'ville train stop.

----

What I've written so far is not very tidy but it does seem to work. Needs tests.

If you happen to have a Go compiler installed then try this:

```
$ go build -o stopstat src/stopstat/main.go 

$ ./stopstat -stop 7093
Taiaroa Street (near 10)
Notices:
    Some cancelled services & CLOSED CBD bus stops for Massey Uni Pde, 1-1.30pm approx. Thurs 26 May. metlink.org.nz
Services:
    6:43AM  Strathmore Park - Outbound - Wgtn Station
    6:59AM  Strathmore Park - Outbound - Khandallah
    7:13AM  Strathmore Park - Outbound - Molesworth Street
    7:19AM  Strathmore Park - Outbound - Khandallah
    7:28AM  Strathmore Park - Outbound - Wgtn Station
    7:43AM  Strathmore Park - Outbound - Molesworth Street
    7:53AM  Strathmore Park - Outbound - Wgtn Station
    7:55AM  Strathmore Park - Inbound - School Bus
    7:59AM  Strathmore Park - Outbound - Khandallah
    8:08AM  Strathmore Park - Outbound - Victoria University
    8:30AM  Strathmore Park - Outbound - Victoria University
    8:59AM  Strathmore Park - Outbound - Khandallah
    9:29AM  Strathmore Park - Outbound - Khandallah
    9:59AM  Strathmore Park - Outbound - Khandallah
    10:29AM  Strathmore Park - Outbound - Khandallah
    10:59AM  Strathmore Park - Outbound - Khandallah
    11:29AM  Strathmore Park - Outbound - Khandallah
    11:59AM  Strathmore Park - Outbound - Khandallah
    12:29PM  Strathmore Park - Outbound - Khandallah
    12:59PM  Strathmore Park - Outbound - Khandallah
```

----

I'm writing a Slack bot which will use this.
