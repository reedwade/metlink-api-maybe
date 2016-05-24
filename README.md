# metlink-api-maybe

Some code I've got working that uses an undocumented API found inside https://www.metlink.org.nz/

There seems to be one call of interest which, given a train or bus stop code returns a reasonably understandable whack of JSON.

`https://www.metlink.org.nz/api/v1/StopDepartures/JOHN`

Where, `JOHN` in this case is the J'vill train stop.

It's not very tidy but it does seem to work. The output is ugly still but it's a start.

If you happen to have a Go compiler installed then typing:

```
$ go build -o stopstat src/stopstat/main.go 

$ ./stopstat -stop 7093
Taiaroa Street (near 10)
Notices:
    Some cancelled services & CLOSED CBD bus stops for Massey Uni Pde, 1-1.30pm approx. Thurs 26 May. metlink.org.nz
Services:
    2016-05-25T06:43:00+12:00  Strathmore Park - Outbound - Wgtn Station
    2016-05-25T06:59:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T07:13:00+12:00  Strathmore Park - Outbound - Molesworth Street
    2016-05-25T07:19:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T07:28:00+12:00  Strathmore Park - Outbound - Wgtn Station
    2016-05-25T07:43:00+12:00  Strathmore Park - Outbound - Molesworth Street
    2016-05-25T07:53:00+12:00  Strathmore Park - Outbound - Wgtn Station
    2016-05-25T07:55:00+12:00  Strathmore Park - Inbound - School Bus
    2016-05-25T07:59:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T08:08:00+12:00  Strathmore Park - Outbound - Victoria University
    2016-05-25T08:30:00+12:00  Strathmore Park - Outbound - Victoria University
    2016-05-25T08:59:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T09:29:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T09:59:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T10:29:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T10:59:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T11:29:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T11:59:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T12:29:00+12:00  Strathmore Park - Outbound - Khandallah
    2016-05-25T12:59:00+12:00  Strathmore Park - Outbound - Khandallah
```


The output could be a lot better.

