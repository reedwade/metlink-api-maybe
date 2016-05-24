package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var StopCode string

func init() {
	flag.StringVar(&StopCode, "stop", "CROF", "train or bus stop code, examples: WELL, JOHN, CROF")
}

func main() {

	flag.Parse()

	report, err := GetStopReport(StopCode)
	if err != nil {
		fmt.Println("PROBLEM:", err)
		return
	}

	fmt.Print(report)
}

const (
	MetLinkAPIv1StopDeparturesUrl = "https://www.metlink.org.nz/api/v1/StopDepartures/%s"
)

type MetLinkAPIv1StopDeparturesResponse struct {
	LastModified string             `json:"LastModified"`
	Notices      NoticesStructList  `json:"Notices"`
	Stop         StopStruct         `json:"Stop"`
	Services     ServicesStructList `json:"Services"`
}

// {'OriginStopID': 'JOHN', 'Direction': 'Inbound', 'OperatorRef': 'RAIL', 'VehicleRef': None, 'DestinationStopID': 'WELL', 'Service': {'Code': 'JVL', 'Link': 'timetables\\/train\\/JVL', 'Name': 'Johnsonville Line (Johnsonville - Wellington)', 'TrimmedCode': 'JVL', 'Mode': 'Train'}, 'DisplayDeparture': '2016-05-24T19:00:00+12:00', 'IsRealtime': False, 'DepartureStatus': None, 'ServiceID': 'JVL', 'VehicleFeature': None, 'OriginStopName': 'Johnsonville Stn', 'DestinationStopName': 'WELL - All stops', 'AimedArrival': None, 'AimedDeparture': '2016-05-24T19:00:00+12:00', 'DisplayDepartureSeconds': 705, 'ExpectedDeparture': None}
type ServicesStruct struct {
	OriginStopID        string `json:"OriginStopID"`
	OriginStopName      string `json:"OriginStopName"`
	Direction           string `json:"Direction"`
	DestinationStopID   string `json:"DestinationStopID"`
	DestinationStopName string `json:"DestinationStopName"`
	DisplayDeparture    string `json:"DisplayDeparture"`
	IsRealtime          bool   `json:"IsRealtime"`
}

type ServicesStructList []*ServicesStruct

func (n *ServicesStruct) String() string {
	var realtime string
	if n.IsRealtime {
		realtime = "(realTime)"
	}
	return fmt.Sprintf("    %s %s %s - %s - %s\n", n.DisplayDeparture, realtime, n.OriginStopName, n.Direction, n.DestinationStopName)
}

func (n ServicesStructList) String() string {
	if len(n) == 0 {
		return ""
	}
	out := "Services:\n"
	for _, v := range n {
		out += v.String()
	}
	return out
}

// {"RecordedAtTime":"2016-05-24T18:48:09+12:00","MonitoringRef":"JOHN","LineRef":"","DirectionRef":"","LineNote":"Services on the JVL are experiencing delays of up to 10 mins"}
type NoticesStruct struct {
	RecordedAtTime string `json:"RecordedAtTime"`
	LineNote       string `json:"LineNote"`
}

type NoticesStructList []*NoticesStruct

// Stop:  {'Name': 'Johnsonville Station', 'LastModified': '2015-09-03T11:14:30+12:00', 'Sms': 'JOHN', 'Long': '174.8047433', 'Farezone': '3', 'Lat': '-41.223345', 'Icon': '\\/assets\\/StopImages\\/JOHN.jpg'}
type StopStruct struct {
	Name string `json:"Name"`
}

func (m *MetLinkAPIv1StopDeparturesResponse) String() string {
	return fmt.Sprintf("%s\n%s%s\n", m.Stop.Name, m.Notices, m.Services)
}

func (n *NoticesStruct) String() string {
	return fmt.Sprintf("    %s\n", n.LineNote)
}

func (n NoticesStructList) String() string {
	if len(n) == 0 {
		return ""
	}
	notices := "Notices:\n"
	for _, v := range n {
		notices += v.String()
	}
	return notices
}

func GetStopReport(stopCode string) (*MetLinkAPIv1StopDeparturesResponse, error) {

	stopCode = strings.ToUpper(stopCode)

	url := fmt.Sprintf(MetLinkAPIv1StopDeparturesUrl, StopCode)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed, maybe a bad stop code?")
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var response MetLinkAPIv1StopDeparturesResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
