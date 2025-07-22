package encoder

import "math/rand/v2"

type ComponentID uint8

// Components defined from https://www.home-assistant.io/integrations/mqtt/#configuration-via-mqtt-discovery
const (
	AlarmControlPanelID ComponentID = iota // 0
	BinarySensorID                         // 1
	ButtonID                               // 2
	CameraID                               // 3
	CoverID                                // 4
	DeviceTrackerID                        // 5
	DeviceTriggerID                        // 6
	EventID                                // 7
	FanID                                  // 8
	HumidifierID                           // 9
	ImageID                                // 10
	ClimateHVACID                          // 11
	LawnmowerID                            // 12
	LightID                                // 13
	LockID                                 // 14
	NotifyID                               // 15
	NumberID                               // 16
	SceneID                                // 17
	SelectID                               // 18
	SensorID                               // 19
	SirenID                                // 20
	SwitchID                               // 21
	UpdateID                               // 22
	TagScannerID                           // 23
	TextID                                 // 24
	VacuumID                               // 25
	ValveID                                // 26
	WaterHeaterID                          // 27
)

func (comp ComponentID) Topic() {
	switch comp {
	case AlarmControlPanelID:
	case BinarySensorID:
	case ButtonID:
	}
}

// type Message struct {
// 	Component Component `cbor:"1,keyasint"`
// 	Topic     string    `cbor:"2,keyasint"`
// }

const (
	CELSIUS    = `°C`
	FAHRENHEIT = `°F`
	HUMIDITY   = `%`
)

/* Message interface for interacting with the MQTT connection
 * Allows anyon to implemnt custom Message types if the provided
 * Messages are insufficient
 */
type Message interface {
	GetTopic() ([]byte, error)
}

type MessageJSON interface {
	MarshalJSON() ([]byte, error)
}

type MessageCBOR interface {
	MarshalCBOR() ([]byte, error)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GenerateID creates a random alpha-numeric string of a given length n
func GenerateID(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

func randInt16() uint16 {
	a := rand.Uint32()
	a %= (65535 - 1)
	a += 1
	return uint16(a)
}

const (
	EncoderTypeError = "invalid encoder type"
)
