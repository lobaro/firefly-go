package firefly

import "testing"

func SKIP_TestShowApplications(t *testing.T) {
	c := setupValidClient(t)
	r, err := c.ShowApplications()

	if err != nil {
		t.Error(err)
	}
	
	t.Log("Applications:",r.Applications)

}
