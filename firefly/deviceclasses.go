package firefly

import "strconv"

func (client Client) ShowDeviceClasses() (r DeviceClassesListResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "device_classes"

	r = DeviceClassesListResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}

func (client Client) ShowSingleDeviceClass(deviceClassId int) (r DeviceClassResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "device_classes/" + strconv.Itoa(deviceClassId)

	r = DeviceClassResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}


func (client Client) ListEUIsOfDeviceClassDevices(deviceClassId int) (r DevicesEuiListResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "device_classes/" + strconv.Itoa(deviceClassId) + "/euis"

	r = DevicesEuiListResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}
