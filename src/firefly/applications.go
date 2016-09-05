package firefly

import "strconv"

func (client Client) ShowApplications() (r ApplicationListResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "applications"

	r = ApplicationListResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}

func (client Client) ListEUIsOfApplicationDevices(applicationId int) (r DevicesEuiListResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "applications/" + strconv.Itoa(applicationId) + "/euis"

	r = DevicesEuiListResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}