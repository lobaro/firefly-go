package firefly

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func setupValidClient(t *testing.T) *Client {
	return NewClient("valid-key")
}

func setupTestClient(t *testing.T) *Client {
	client := NewClient("key")
	httpmock.ActivateNonDefault(client.http)
	return client
}
