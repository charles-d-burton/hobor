package encoder

import (
	"encoding/json"
	"errors"

	"github.com/fxamacker/cbor/v2"
)

type SensorOptions []string

// https://www.home-assistant.io/integrations/sensor.mqtt/
type Sensor struct {
	Config
	Platform                  string         `json:"platform"`
	Encoding                  *string        `json:"encoding,omitempty"`
	EntityCategory            *string        `json:"entity_category,omitempty"`
	EntityPicture             *string        `json:"entity_picture,omitempty"`
	ExpireAfter               int            `json:"expire_after"`
	ForceUpdate               bool           `json:"force_update"`
	Icon                      *string        `json:"icon,omitempty"`
	JSONAttributesTemplate    *string        `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic       *string        `json:"json_attributes_topic,omitempty"`
	LastResetValueTemplate    *string        `json:"last_reset_value_template,omitempty"`
	Options                   *SensorOptions `json:"options,omitempty"`
	PayloadAvailable          *string        `json:"payload_available,omitempty"`
	PayloadNotAvailable       *string        `json:"payload_not_available,omitempty"`
	SuggestedDisplayPrecision *int           `json:"suggested_display_precision,omitempty"`
	QOS                       int            `json:"qos,omitempty"`
	StateClass                *string        `json:"state_class,omitempty"`
	UnitOfMeasurement         *string        `json:"unit_of_measurement,omitempty"`
	ValueTemplate             *string        `json:"value_template,omitempty"`
}

func NewSensor(config Config) (*Sensor, error) {
	return &Sensor{Config: config}, nil
}

func (sensor *Sensor) GetTopic() ([]byte, error) {
	err := sensor.validateComponent("sensor")
	if err != nil {
		return nil, err
	}
	sensor.Platform = "sensor"
	return []byte(sensor.ConfigTopic), nil
}

func (sensor *Sensor) Marshal(es EncoderSwitch) ([]byte, error) {
	err := sensor.validateAvailability()
	if err != nil {
		return nil, err
	}
	if sensor.UniqueID == "" {
		return nil, errors.New("unique_id-undefined")
	}
	if sensor.Name == "" {
		return nil, errors.New("name-undefined")
	}
	switch es {
	case JSON:
		return json.Marshal(sensor)
	case CBOR:
		return cbor.Marshal(sensor)
	default:
		return nil, errors.New(EncoderTypeError)
	}
}

/* This particular device has a lot of different possible permutations
 * Due to the constrained nature of tinygo they'll need to be defined here
 * Alternatively the writer can implement the Mmessage interface themselves
 * Adding commonly used ones here
 */

const (
	TemperatureTemplate = `{{ value_json.temperature }}`
	HumidityTemplate    = `{{ value_json.humidity }}`
)

type Temperature struct {
	Topic       string  `json:"-"`
	TempReading float32 `json:"temperature"`
}

func (t *Temperature) GetTopic() ([]byte, error) {
	if t.Topic == "" {
		return nil, errors.New("error:topic not set")
	}
	return []byte(t.Topic), nil
}

func (t *Temperature) Marshal(es EncoderSwitch) ([]byte, error) {
	return json.Marshal(t)
}

type Humidity struct {
	Topic           string  `json:"-"`
	HumidityReading float32 `json:"humidity"`
}

func (h *Humidity) GetTopic() ([]byte, error) {
	if h.Topic == "" {
		return nil, errors.New("error:topic not set")
	}
	return []byte(h.Topic), nil
}

func (h *Humidity) Marshal(es EncoderSwitch) ([]byte, error) {
	return json.Marshal(h)
}

type TempAndHumidity struct {
	Topic           string  `json:"-"`
	TempReading     float32 `json:"temperature"`
	HumidityReading float32 `json:"humidity"`
}

func (th *TempAndHumidity) GetTopic() ([]byte, error) {
	if th.Topic == "" {
		return nil, errors.New("error:topic not set")
	}
	return []byte(th.Topic), nil
}

func (th *TempAndHumidity) Marshal(es EncoderSwitch) ([]byte, error) {
	return json.Marshal(th)
}
