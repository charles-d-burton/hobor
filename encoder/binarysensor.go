package encoder

import (
	"errors"

	"github.com/fxamacker/cbor/v2"
)

// BinarySensor represents a Binary Sensor MQTT struct
// homeassistant.com/integrations/binary_sensor.mqtt/
type BinarySensor struct {
	Config
	Encoding               string `json:"encoding,omitempty,omitzero"`
	EntityCategory         string `json:"entity_category,omitempty,omitzero"`
	EntityPicture          string `json:"entity_picture,omitempty,omitzero"`
	ExpireAfter            int    `json:"expire_after,omitempty,omitzero"`
	ForceUpdate            bool   `json:"force_update,omitempty,omitzero"`
	Icon                   string `json:"icon,omitempty,omitzero"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty,omitzero"`
	JSONAttributesTopic    string `json:"json_attributes_topic,omitempty,omitzero"`
	Name                   string `json:"name,omitempty,omitzero"`
	ObjectID               string `json:"object_id,omitempty,omitzero"`
	OffDelay               int    `json:"off_delay,omitempty,omitzero"`
	PayloadAvailable       string `json:"payload_available,omitempty,omitzero"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty,omitzero"`
	PayloadOff             string `json:"payload_off,omitempty,omitzero"`
	PayloadOn              string `json:"payload_on,omitempty,omitzero"`
	Platform               string `json:"platform,omitzero"`
	Qos                    int    `json:"qos,omitempty,omitzero"`
	StateTopic             string `json:"state_topic,omitzero"`
	UniqueID               string `json:"unique_id,omitempty,omitzero"`
	ValueTemplate          string `json:"value_template,omitempty,omitzero"`
}

func NewBinarySensor(config Config) (*BinarySensor, error) {
	return &BinarySensor{Config: config}, nil
}

func (b *BinarySensor) GetTopic() ([]byte, error) {
	err := b.validateComponent("binary_sensor")
	if err != nil {
		return nil, err
	}
	return []byte(b.ConfigTopic), nil
}

func (b *BinarySensor) Marshal() ([]byte, error) {
	return nil, nil
}

type BinarySensorState struct {
	StateTopic string `json:"-"`
	State      string `cbor:"1,keyasint" json:"state"`
}

func (bss *BinarySensorState) GetTopic() ([]byte, error) {
	if bss.StateTopic == "" {
		return nil, errors.New("state topic is empty")
	}
	return []byte(bss.StateTopic), nil
}

func (bss *BinarySensorState) MarshalCBOR() ([]byte, error) {
	switch bss.State {
	case "ON", "OFF":
		return cbor.Marshal(bss)
	default:
		return nil, errors.New("state is not ON or OFF")
	}
}
