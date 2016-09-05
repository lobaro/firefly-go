package firefly

import (
	"testing"
	"encoding/json"
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
