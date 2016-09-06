package firefly

import (
	"time"
	"fmt"
)

type EUI string
type DeviceAddress string

type DigimondoReponse struct {
	Error string `json:"error"`
}

type LocalTimeWithoutZone struct {
	time.Time
}

func (t *LocalTimeWithoutZone) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(`"%s"`, t.Time.Format("2006-01-02T15:04:05"))
	return []byte(stamp), nil
}

const localTimeWithoutZoneFormat = "2006-01-02T15:04:05"

func (t *LocalTimeWithoutZone) UnmarshalJSONXX(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t.Time, err = time.Parse(`"` + time.RFC3339 + `"`, string(data))
	return err
}

func (t *LocalTimeWithoutZone) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.ParseInLocation(`"` + localTimeWithoutZoneFormat + `"`, string(data), time.Local)

	if err != nil {
		return err
	}
	return nil
}

type DeviceCreateRequest struct {
	Organization int         `json:"organization"`
	Application  int         `json:"application"`
	Device       DeviceCreate      `json:"device"`
}

type DeviceCreate struct {
	Address               string `json:"address"`
	ApplicationKey        string `json:"application_key"`
	ApplicationSessionKey string `json:"application_session_key"`
	Description           string `json:"description"`
	Eui                   string `json:"eui"`
	NetworkSessionKey     string `json:"network_session_key"`
	Otaa                  bool `json:"otaa"`
	Tags                  []string `json:"tags"` // TODO: valid? comma separated list in example.
}

type DeviceUpdate struct {
	Address               string `json:"address,omitempty"`
	ApplicationKey        string `json:"application_key,omitempty"`
	ApplicationSessionKey string `json:"application_session_key,omitempty"`
	Description           string `json:"description,omitempty"`
	Eui                   string `json:"eui,omitempty"`
	NetworkSessionKey     string `json:"network_session_key,omitempty"`
	Otaa                  bool `json:"otaa,omitempty"`
	Tags                  []string `json:"tags,omitempty"` // TODO: valid?
}

type DeviceUpdateRequest struct {
	Device DeviceUpdate `json:"device"`
}

type Device struct {
	Address               string `json:"address"`
	ApplicationKey        string `json:"application_key"`
	ApplicationSessionKey string `json:"application_session_key"`
	Description           string `json:"description"`
	Eui                   string `json:"eui"`
	NetworkSessionKey     string `json:"network_session_key"`
	Otaa                  bool `json:"otaa"`
	Tags                  []string `json:"tags"`
	CreatedAt             string `json:"created_at"`
	DeviceClassId         int `json:"device_class_id"`
	UpdatedAt             string `json:"updated_at"`
}

type DeviceListResponse struct {
	DigimondoReponse
	Devices []Device `json:"devices,omitempty"`
}

type DeviceResponse struct {
	DigimondoReponse
	Device Device `json:"device,omitempty"`
}

type ListPacketsParams struct {
	Direction     string
	LimitToLast   int
	Offset        int
	PayloadOnly   bool
	ReceivedAfter *time.Time
}

type ListAllPacketsParams struct {
	Direction     string
	LimitToLast   int
	Offset        int
	PayloadOnly   bool
	ReceivedAfter *time.Time
	SkipSuborgs   bool
}

type PacketListResponse struct {
	DigimondoReponse
	Packets []struct {
		Ack              bool `json:"ack"`
		Bandwidth        int `json:"bandwidth"`
		Codr             string `json:"codr"`
		Datr             string `json:"datr"` // TODO: type? Not specified in docu
		DeviceEui        string `json:"device_eui"`
		Fopts            string `json:"fopts"`
		Fcnt             int `json:"fcnt"`
		Freq             float64 `json:"freq"`
		Gwrx             []struct {
			Gweui string `json:"gweui"`
			Lsnr  float64 `json:"lsnr"`
			Rssi  int `json:"rssi"`
			Time  time.Time `json:"time"`
			Tmst  int64 `json:"tmst"`
		} `json:"gwrx"`
		Modu             string `json:"modu"`
		Mtype            string `json:"mtype"`
		Parsed           string `json:parsed`
		Payload          string `json:"payload"`
		PayloadEncrypted bool `json:"payload_encrypted"`
		Port             int `json:"port"`
		ReceivedAt       string `json:"received_at"`
		Size             int `json:"size"`
		SpreadingFactor  int `json:"spreading_factor"`
	} `json:"packets,omitempty"`
}

type CreateDeviceRequest struct {
	Organization int `json:"organization"`
	Application  int `json:"application"`
	Device       struct {
			     Otaa        bool `json:"otaa"`
			     Eui         string `json:"eui"`
			     Description string `json:"description"`
			     Address     string `json:"address"`
			     Tags        string `json:"tags"`
		     } `json:"device"`
}

type UpdateDeviceRequest struct {
	Device struct {
		       Description string `json:"description"`
	       } `json:"device"`
}

type SendPacketRequest struct {
	Payload  string `json:"payload"`
	Encoding string `json:"encoding"`
	Port     int `json:"port"`
}

type SendPacketResponse struct {
	DigimondoReponse
	SentPacket struct {
			   Fcnt    int `json:"fcnt"`
			   Id      int `json:"id"`
			   Payload string `json:"payload"`
			   Port    int `json:"port"`
		   } `json:"sent_packet,omitempty"`
}

type ApplicationListResponse struct {
	DigimondoReponse
	Applications []struct {
		CreatedAt   LocalTimeWithoutZone `json:"created_at"`  // TODO: used in doc but not in example
		InsertedAt  LocalTimeWithoutZone `json:"inserted_at"` // TODO: used in example but not in doc
		Description string `json:"description"`
		Eui         string `json:"eui"`
		Id          int `json:"id"`
		Name        string `json:"name"`
		Sink        interface{} `json:"sink"`
		UpdatedAt   LocalTimeWithoutZone `json:"updated_at"`
	} `json:"applications,omitempty"`
}

type DevicesEuiListResponse struct {
	DigimondoReponse
	Devices []struct {
		Address string `json:"address"`
		Eui     string `json:"eui"`
		Id      string `json:"eui,omitempty"` // Used by "List EUIs of Devices" in device classes but not application
	} `json:"devices,omitempty"`
}

type DeviceVariables struct {
	Gps          DeviceVariable `json:"gps"`
	BatteryLevel DeviceVariable `json:"batteryLevel"`
	Battery      DeviceVariable `json:"battery"`
}

type DeviceVariable struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type DeviceParseElement struct {
	Type   string `json:"type"`
	Target string `json:"target"`
	Bits   string `json:"bits"`
	Signed bool `json:"signed,omitempty"`
}

type DeviceCalculateElement struct {
	Target  string `json:"target"`
	Formula string `json:"formula"`
}

type DeviceClass     struct {
	Description string `json:"description"`
	Id          int `json:"id"`
	InsertedAt  string `json:"inserted_at"`
	Name        string `json:"name"`
	Script      struct {
			    Variables         DeviceVariables `json:"variables"`
			    ParseElements     []DeviceParseElement `json:"parseElements"`
			    CalculateElements []DeviceCalculateElement `json:"calculateElements"`
		    } `json:"script"`
	UpdatedAt   string `json:"updated_at"`
}

type DeviceClassesListResponse struct {
	DigimondoReponse
	DeviceClasses []DeviceClass `json:"device_classes,omitempty"`
}

type DeviceClassResponse struct {
	DigimondoReponse
	DeviceClasses DeviceClass `json:"device_classes,omitempty"` // TODO: typo in doc of really 'device_classes'?
}

