package encoder

import (
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

func (sensor *Sensor) MarshalCBOR() ([]byte, error) {
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
	return cbor.Marshal(sensor)
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
		return nil, errors.New("topic not set")
	}
	return []byte(t.Topic), nil
}

func (t *Temperature) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(t)
}

type Humidity struct {
	Topic           string  `json:"-"`
	HumidityReading float32 `json:"humidity"`
}

func (h *Humidity) GetTopic() ([]byte, error) {
	if h.Topic == "" {
		return nil, errors.New("topic not set")
	}
	return []byte(h.Topic), nil
}

func (h *Humidity) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(h)
}

type TempAndHumidity struct {
	Topic           string  `json:"-"`
	TempReading     float32 `json:"temperature"`
	HumidityReading float32 `json:"humidity"`
}

func (th *TempAndHumidity) GetTopic() ([]byte, error) {
	if th.Topic == "" {
		return nil, errors.New("topic not set")
	}
	return []byte(th.Topic), nil
}

func (th *TempAndHumidity) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(th)
}
