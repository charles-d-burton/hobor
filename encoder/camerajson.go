//go:build excludetinygo

package encoder

import (
	"encoding/json"
	"errors"
)

func (c *Camera) MarshalJSON() ([]byte, error) {
	if c.ImageEncoding != "" {
		if c.ImageEncoding != "b64" {
			return nil, errors.New("image encoding not set to b64")
		}
	}

	if c.Topic == "" {
		return nil, errors.New("subscribe topic is empty")
	}
	return json.Marshal(c)
}
