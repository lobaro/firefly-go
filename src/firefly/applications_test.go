package firefly

import "testing"

func TestShowApplications(t *testing.T) {
	c := setupValidClient(t)
	r, err := c.ShowApplications()

	if err != nil {
		t.Error(err)
	}

	t.Error(r)
}
