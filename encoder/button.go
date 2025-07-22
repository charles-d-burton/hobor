package encoder

import (
	"encoding/json"
	"errors"

	"github.com/fxamacker/cbor/v2"
)

// Button defines a Button message
// https://www.home-assistant.io/integrations/button.mqtt/
type Button struct {
	Config
	CommandTemplate string `json:"command_template,omitempty"`
	Platform        string `json:"platform"`
}

func NewButton(config Config) (*Button, error) {
	return &Button{Config: config}, nil
}

func (b *Button) GetTopic() ([]byte, error) {
	err := b.validateComponent("button")
	if err != nil {
		return nil, err
	}
	return []byte(b.ConfigTopic), nil
}

func (b *Button) Marshal(es EncoderSwitch) ([]byte, error) {
	b.Platform = "button"
	if b.CommandTopic == "" {
		return nil, errors.New("error:command topic is empty")
	}
	switch es {
	case JSON:
		return json.Marshal(b)
	case CBOR:
		return cbor.Marshal(b)
	default:
		return nil, errors.New(EncoderTypeError)
	}
}
