package firefly

import (
	"time"
	"strconv"
)

func (client Client) ShowAllDevices() (r DeviceListResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "devices"

	r = DeviceListResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}

func (client Client) ShowDeviceByEui(eui EUI) (r DeviceResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "devices/eui/" + string(eui)

	r = DeviceResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}

func (client Client) ShowDeviceByAddress(address DeviceAddress) (r DeviceResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "devices/address/" + string(address)

	r = DeviceResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}

func (client Client) CreateDevice(request DeviceCreateRequest) (r DeviceResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "devices"

	r = DeviceResponse{}
	err = client.postAndDecode(reqUrl, request, &r)
	return
}

func (client Client) UpdateDevice(eui EUI, request DeviceUpdateRequest) (r DeviceResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "devices/eui/" + string(eui)

	r = DeviceResponse{}
	err = client.putAndDecode(reqUrl, request, &r)
	return
}
func (client Client) DeleteDevice(eui EUI) (err error) {
	reqUrl := client.Url()
	reqUrl.Path += "devices/eui/" + string(eui)

	err = client.deleteNoContent(reqUrl)
	return
}

func (client Client) ListDevicePackets(eui EUI, params ListDevicePacketsParams) (r PacketListResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "devices/eui/" + string(eui) + "/packets"

	q := reqUrl.Query()
	if (params.Direction != "") {
		q.Set("direction", params.Direction)
	}
	if (params.LimitToLast != 0) {
		q.Set("limit_to_last", strconv.Itoa(params.LimitToLast))
	}
	if (params.Offset != 0) {
		q.Set("offset", strconv.Itoa(params.Offset))
	}
	if (params.PayloadOnly != false) {
		q.Set("payload_only", strconv.FormatBool(params.PayloadOnly))
	}
	if (params.ReceivedAfter != nil) {
		q.Set("received_after", params.ReceivedAfter.Format(time.RFC3339))
	}

	reqUrl.RawQuery = q.Encode()

	r = PacketListResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}

func (client Client) SendPacketToDevice(eui EUI, request SendPacketRequest) (r SendPacketResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "devices/eui/" + string(eui) + "/packets"

	err = client.postAndDecode(reqUrl, request, &r)
	return
}
