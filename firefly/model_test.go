package firefly

import (
	"testing"
	"encoding/json"
	"time"
	"net/url"
)

func TestDeviceUpdateCanBeEmpty(t *testing.T) {
	bytes, err := json.Marshal(&DeviceUpdate{})

	if err != nil {
		t.Error(err)
	}
	str := string(bytes)

	if str != "{}" {
		t.Error("Expected empty json object but got: " + str)
	}
}

type TimeTestObj struct {
	Time LocalTimeWithoutZone `json:"time"`
}

func TestMarshalDataWithoutTimezone(t *testing.T) {
	jsonObj := &TimeTestObj{
		Time: LocalTimeWithoutZone{time.Unix(1473152776, 0)},
	}

	//t.Log("jsonObj: ", jsonObj)

	b, err := json.Marshal(&jsonObj)

	if err != nil {
		t.Error(err)
	}

	expected := `{"time":"2016-09-06T11:06:16"}`
	if string(b) != expected {
		t.Log(string(b))
		t.Error(string(b) + " != " + expected)
	}
}

func TestParseDataWithoutTimezone(t *testing.T) {
	parsed := TimeTestObj{}

	data := `{"time":"2016-09-06T11:06:16"}`
	err := json.Unmarshal([]byte(data), &parsed)

	if err != nil {
		t.Error(err)
	}

	// NOTE: This works only with local time in germany... 
	// TODO: What time is digimondo using? Should we fix it?
	if parsed.Time.Unix() != 1473152776 {
		t.Error(parsed.Time.Unix(), "!=", 1473152776)
	}
}

func TestListDevicePacketsParamsFromQuery(t *testing.T) {
	q := url.Values{}

	params := ListDevicePacketsParamsFromQuery(q)

	if params.LimitToLast != 0 {
		t.Error("LimitToLast not 0")
	}
	if params.Direction != "" {
		t.Error("Direction not empty")
	}
	if params.Offset != 0 {
		t.Error("Offset not 0")
	}
	if params.PayloadOnly != false {
		t.Error("PayloadOnly not false")
	}
	if params.ReceivedAfter != nil {
		t.Error("ReceivedAfter not nil")
	}

	q.Set("limit_to_last", "42")
	q.Set("direction", "asc")
	q.Set("offset", "23")
	q.Set("payload_only", "true")
	q.Set("received_after", "2016-01-02T13:55:01")

	params = ListDevicePacketsParamsFromQuery(q)

	if params.LimitToLast != 42 {
		t.Error("LimitToLast not 42")
	}
	if params.Direction != "asc" {
		t.Error("Direction not asc")
	}
	if params.Offset != 23 {
		t.Error("Offset not 23")
	}
	if params.PayloadOnly != true {
		t.Error("PayloadOnly not true")
	}
	expectedTime, _ := time.ParseInLocation(localTimeWithoutZoneFormat, "2016-1-2T13:55:01", time.UTC);
	if params.ReceivedAfter.Equal(expectedTime) {
		t.Error("ReceivedAfter not " + expectedTime.Format(localTimeWithoutZoneFormat))
	}
}

func TestListAllPacketsParamsFromQuery(t *testing.T) {
	q := url.Values{}

	params := ListAllPacketsParamsFromQuery(q)

	if params.LimitToLast != 0 {
		t.Error("LimitToLast not 0")
	}
	if params.Direction != "" {
		t.Error("Direction not empty")
	}
	if params.Offset != 0 {
		t.Error("Offset not 0")
	}
	if params.PayloadOnly != false {
		t.Error("PayloadOnly not false")
	}
	if params.ReceivedAfter != nil {
		t.Error("ReceivedAfter not nil")
	}
	if params.SkipSuborgs != false {
		t.Error("SkipSuborgs not false")
	}

	q.Set("limit_to_last", "42")
	q.Set("direction", "asc")
	q.Set("offset", "23")
	q.Set("payload_only", "true")
	q.Set("received_after", "2016-01-02T13:55:01")
	q.Set("skip_suborgs", "true")

	params = ListAllPacketsParamsFromQuery(q)

	if params.LimitToLast != 42 {
		t.Error("LimitToLast not 42")
	}
	if params.Direction != "asc" {
		t.Error("Direction not asc")
	}
	if params.Offset != 23 {
		t.Error("Offset not 23")
	}
	if params.PayloadOnly != true {
		t.Error("PayloadOnly not true")
	}
	expectedTime, _ := time.ParseInLocation(localTimeWithoutZoneFormat, "2016-1-2T13:55:01", time.UTC);
	if params.ReceivedAfter.Equal(expectedTime) {
		t.Error("ReceivedAfter not " + expectedTime.Format(localTimeWithoutZoneFormat))
	}
	if params.SkipSuborgs != true {
		t.Error("SkipSuborgs not true")
	}
}
