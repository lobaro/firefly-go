package firefly

import (
	"time"
	"strconv"
)

func (client Client) AllPackets(params ListAllPacketsParams) (r PacketListResponse, err error) {
	reqUrl := client.Url()
	reqUrl.Path += "/packets"

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

	if (params.SkipSuborgs != false) {
		q.Set("received_after", strconv.FormatBool(params.SkipSuborgs))
	}

	reqUrl.RawQuery = q.Encode()

	r = PacketListResponse{}
	err = client.getAndDecode(reqUrl, &r)
	return
}
