package firefly

import "testing"

func TestNewClient(t *testing.T) {
	client := NewClient("api-key")
	if client.secretApiKey != "api-key" {
		t.Error("Api Key is not set")
	}

	if client.Url().String() != "https://api.fireflyiot.com/api/v1/?auth=api-key"  {
		t.Error("Url is wrong: " + client.Url().String())
	}
}