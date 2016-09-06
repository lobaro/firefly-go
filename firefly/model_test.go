package firefly

import (
	"testing"
	"encoding/json"
	"time"
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
