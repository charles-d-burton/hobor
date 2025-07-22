package encoder

import (
	"encoding/json"
	"errors"

	"github.com/fxamacker/cbor/v2"
)

// https://www.home-assistant.io/integrations/camera.mqtt/
type Camera struct {
	Config
	ImageEncoding string `json:"image_encoding"`
	Topic         string `json:"topic"`
}

func NewCamera(config Config) (*Camera, error) {
	return &Camera{Config: config}, nil
}

func (c *Camera) GetTopic() ([]byte, error) {
	err := c.validateComponent("camera")
	if err != nil {
		return nil, err
	}
	return []byte(c.ConfigTopic), nil
}

func (c *Camera) Marshal(es EncoderSwitch) ([]byte, error) {
	if c.ImageEncoding != "" {
		if c.ImageEncoding != "b64" {
			return nil, errors.New("error:image encoding not set to b64")
		}
	}

	if c.Topic == "" {
		return nil, errors.New("error:subscribe topic is empty")
	}
	switch es {
	case JSON:
		return json.Marshal(c)
	case CBOR:

		return cbor.Marshal(c)
	default:
		return nil, errors.New(EncoderTypeError)
	}
}
