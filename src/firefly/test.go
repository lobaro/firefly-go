package firefly

import "testing"

func setupValidClient(t *testing.T) *Client {
	return NewClient("valid-key")
}
